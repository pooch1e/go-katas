# Algorithms 06 · Heaps & Priority Queues

**Builds on:** language 09 (generic types, interfaces), 10–11 (building your own containers), algorithms 01 (sift = partial sorting), 05 (Dijkstra wants one of these).

```sh
go test ./algorithms/06-heaps/
```

A **heap** answers one question fast: *what's the smallest thing right now?*, even while things are being added and removed. That's a **priority queue**: job schedulers, "top k" problems, Dijkstra, merging sorted streams.

## Complexity

| Operation | Binary heap | Sorted slice | Unsorted slice |
|-----------|-------------|--------------|----------------|
| Push | **O(log n)** | O(n) | O(1) |
| Pop min | **O(log n)** | O(1) | O(n) |
| Peek min | O(1) | O(1) | O(n) |
| Build from n items | O(n) with `heap.Init` | O(n log n) | — |

## How a binary heap sits in a slice

```
         1               index:  0  1  2  3  4  5
       /   \             value: [1  3  2  5  9  8]
      3     2
     / \   /             parent(i) = (i-1)/2
    5   9 8              children(i) = 2i+1, 2i+2
```

The rule is only that **every parent ≤ its children**. Siblings can be in any order, so the slice is *not* sorted.

- **Push:** append at the end, then **sift up**: swap with the parent while smaller.
- **Pop:** take index 0, move the last element to the root, then **sift down**: swap with the *smaller* child while bigger.

## Go you'll lean on

- **`container/heap`.** The standard library does the sifting for any type that implements `heap.Interface`: `Len`, `Less`, `Swap` (from `sort.Interface`), plus `Push(any)` and `Pop() any`.
- ⚠️ You call **`heap.Push(h, x)` / `heap.Pop(h)`**, never `h.Push(x)` directly. Your methods only add or remove at the end of the slice; the `heap` package functions call them *and* sift.
- **`Pop() any` returns `any`.** Get your type back with a type assertion: `heap.Pop(h).(int)`.
- **Pointer receivers for Push/Pop** (they change the slice length), value receivers for the other three.
- For "max-heap", flip `Less`. For structs (value + frequency, value + which list), make the heap a slice of a small struct.

---

## The katas

### MinHeap[T] — sift up, sift down

**The concept:** implement the two sifts by hand on a generic slice. Once you've done this, `container/heap` stops being magic.

```js
// JS: the parent/child index arithmetic, used here to draw a tree level by level
const level = (i) => Math.floor(Math.log2(i + 1));
const parent = (i) => Math.floor((i - 1) / 2);
```

**Go building blocks:**
- private `up(i)` and `down(i)` helpers that loop while the heap rule is broken
- multi-assign swaps
- in `down`, pick the smaller of the (up to) two children

### IntHeap — plugging into container/heap

**The concept:** **adapting your type to a standard interface** (language 09's lesson). Implement five small methods and the library brings the algorithm.

**Go building blocks:**
- `Less` compares `h[i] < h[j]`
- `Push` appends `x.(int)`
- `Pop` removes and returns the **last** element (the library has already swapped the minimum there)

### KthLargest — a size-k min-heap

**The concept:** to find the k-th largest, you don't need everything sorted. Keep a min-heap of the **k largest seen so far**. When it grows past k, pop the smallest. At the end, the heap's top is the answer.

```js
// JS: the neighbouring idea, a top-3 leaderboard that only keeps 3 entries
const top3 = [];
const add = (score) => { top3.push(score); top3.sort((a, b) => b - a); top3.length = Math.min(3, top3.length); };
```

**Go building blocks:** `&IntHeap{}`, `heap.Push`/`heap.Pop`, `(*h)[0]` to peek, and guards for `k`.

### TopKFrequent — count, then heap by frequency

**The concept:** two steps. First count frequencies (language 02). Then select the top k by count, using a heap of `(value, count)` pairs or sorting the distinct values. A tie-break makes the answer deterministic.

**Go building blocks:**
- `counts := map[int]int{}`
- a small `type pair struct{ val, count int }`
- `slices.SortFunc` with `cmp.Or` (language 06), or a heap of pairs

### MergeKSorted — a heap of "next candidates"

**The concept:** you have k sorted lists. The next output value is the smallest *front* among them, so keep a min-heap holding one entry per list: its current front value, plus which list and index it came from. Pop the smallest, emit it, and push that list's next element.

```
heap: (1 from L0) (2 from L1) (3 from L2) → pop 1, push 4 from L0 → pop 2, push 5 from L1 → ...
```

**Go building blocks:** a heap of `struct{ val, list, idx int }`, skipping empty lists when seeding.
