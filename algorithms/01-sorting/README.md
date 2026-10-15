# Algorithms 01 · Sorting

**Builds on:** language 01 (in-place vs new slice, swaps), 06 (generics), 09 (`cmp.Ordered`).

```sh
go test ./algorithms/01-sorting/
go test -bench=. ./algorithms/01-sorting/   # compare the four, once they pass
```

In real code you call `slices.Sort` (a hybrid "pattern-defeating quicksort"). You write these by hand to learn **how to reason about cost**: why two nested loops are O(n²), why halving gives you log n, and what "in place" and "stable" mean.

For the algorithms track, the explanations use **traces** (the state after each step) rather than code, because the JS version of a sort *is* the answer. JS sketches show a neighbouring idea only.

## Complexity at a glance

| Sort | Best | Average | Worst | Extra memory | Stable? |
|------|------|---------|-------|--------------|---------|
| Bubble | O(n) with early exit | O(n²) | O(n²) | O(1) | yes |
| Insertion | O(n) | O(n²) | O(n²) | O(1) | yes |
| Merge | O(n log n) | O(n log n) | O(n log n) | O(n) | yes |
| Quick | O(n log n) | O(n log n) | O(n²) | O(log n) stack | no |

**Stable** means equal elements keep their original relative order, which matters when sorting records by one field after another.

## Go you'll lean on

- Generic signatures: `func BubbleSort[T cmp.Ordered](xs []T)`. Use `<` and `>` directly on `T`.
- Swap: `xs[i], xs[j] = xs[j], xs[i]`.
- Sub-slices for recursion: `xs[:mid]` and `xs[mid:]` are free (no copy). That's great for QuickSort's in-place partitions, and a trap for MergeSort if you write results back into a slice you're still reading.
- Benchmarks: `func BenchmarkX(b *testing.B) { for b.Loop() { ... } }`, run with `-bench`.

---

## The katas

### BubbleSort — repeated neighbour swaps

**The concept:** sweep left to right, swapping any adjacent pair that's out of order. After one sweep the largest value has "bubbled" to the end. Repeat on the shrinking unsorted part. **Early exit:** if a whole sweep makes no swaps, you're done.

```
[5 1 4 2]  →  sweep 1: [1 4 2 | 5]  →  sweep 2: [1 2 | 4 5]  →  sweep 3: no swaps, stop
```

**Go building blocks:**
- an outer loop over passes
- an inner loop up to `len(xs)-1-pass`
- a `swapped bool` flag

### InsertionSort — grow a sorted prefix

**The concept:** the left part of the slice is always sorted. Take the next element and shift bigger elements right until its slot opens up, then drop it in. It's how most people sort a hand of cards. It's fast on nearly-sorted data, which is why real hybrid sorts use it for small chunks.

```
[2 5 | 3 1]  take 3: shift 5 right → [2 _ 5 1] → place → [2 3 5 | 1]
```

```js
// JS: the neighbouring idea, inserting ONE item into an already-sorted list
const insertSorted = (sorted, x) => { const i = sorted.findIndex((y) => y > x); return i < 0 ? [...sorted, x] : sorted.toSpliced(i, 0, x); };
```

**Go building blocks:**
- hold the current value in a variable before you start shifting (it'll get overwritten otherwise)
- a `for` with two conditions joined by `&&`: still in bounds, and still bigger than the held value

### MergeSort — divide, sort, merge

**The concept:** **divide and conquer**.
- Split in half, sort each half recursively, then **merge** the two sorted halves by repeatedly taking the smaller front element.
- A slice of 0 or 1 elements is already sorted (the base case).
- There are log n levels of splitting, with O(n) merge work per level.

```
[5 2 4 1] → [5 2] [4 1] → [5][2] [4][1] → [2 5] [1 4] → [1 2 4 5]
```

**Go building blocks:**
- `mid := len(xs) / 2`
- a `merge(a, b []T) []T` helper with two indexes and `make([]T, 0, len(a)+len(b))`
- don't forget to append whichever side has leftovers

### QuickSort — partition around a pivot

**The concept:**
- Pick a **pivot** and rearrange the slice so smaller values are on its left and larger on its right. The pivot is now in its final spot.
- Recurse on the two sides.
- If the pivot is always the smallest or largest, the splits are lopsided and you hit O(n²). Choosing the middle element or a random one avoids that for sorted input.

```
[3 7 1 5 2], pivot 3 → [1 2 | 3 | 7 5] → recurse on [1 2] and [7 5]
```

**Go building blocks:**
- a `partition(xs []T) int` helper (Lomuto: one "boundary" index that walks forward, swapping smaller values behind it)
- recurse on `xs[:p]` and `xs[p+1:]`, which are sub-slices, so it's all in place
