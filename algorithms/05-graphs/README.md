# Algorithms 05 · Graphs

**Builds on:** language 02 (maps), 07 (seen-sets), 08 (stacks and queues), 11 (DFS/BFS on trees), algorithms 04 (recursion).

```sh
go test ./algorithms/05-graphs/
```

A tree is a graph with no cycles and one root. Graphs drop both rules, so the one new idea is a **visited set**, so you never process a node twice.

## Complexity

V = vertices (nodes), E = edges.

| Kata | Time | Space |
|------|------|-------|
| HasPath (DFS) | O(V + E) | O(V) |
| ShortestPathGrid (BFS) | O(rows × cols) | O(rows × cols) |
| NumIslands | O(rows × cols) | O(rows × cols) |
| TopologicalSort (Kahn's) | O(V + E) | O(V) |
| Dijkstra (simple scan) | O(V²) | O(V) |
| Dijkstra (with a heap, after module 06) | O((V + E) log V) | O(V) |

## Representations

- **Adjacency map:** `map[string][]string` (here, `Graph`). `g["a"]` is the list of nodes `a` points to. A missing key reads as an empty slice, which is handy.
- **Grid as an implicit graph:** each cell is a node, and its up/down/left/right neighbours are its edges. You never build a map; you compute neighbours on the fly.

## Go you'll lean on

- `visited := map[string]bool{}`, or `map[Cell]bool{}`, because structs of comparable fields work as map keys.
- Direction vectors: `dirs := []Cell{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}`.
- Bounds checks before indexing: `r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r])`.
- `grid[r][c]` on a `string` gives a **byte**, so compare with `'#'` (a rune literal, which works for ASCII).
- Queues from language 08: `queue = queue[1:]`.
- Recursive closures from algorithms 04.

---

## The katas

### HasPath — depth-first search

**The concept:** go as deep as possible along one path before backing up to try the next. Mark each node visited the moment you reach it, or a cycle will loop you forever.

```js
// JS: the same walk over a folder structure that might contain symlink loops
const seen = new Set();
const walk = (dir) => { if (seen.has(dir.realPath)) return; seen.add(dir.realPath); dir.children.forEach(walk); };
```

**Go building blocks:** a recursive closure *or* an explicit stack (`[]string`), a visited map, and an early `return true` when you reach `to`.

### ShortestPathGrid — breadth-first search

**The concept:**
- BFS explores in rings: all cells 1 step away, then all 2 steps away, and so on.
- So the **first time** BFS reaches a cell, it got there by a shortest path.
- DFS doesn't have this property, which is why "shortest path, unweighted" always means BFS.

```js
// JS: the same ring-by-ring idea, "who's within 2 hops in my network?"
let ring = [me], dist = 0; const seen = new Set([me]);
while (ring.length && dist < 2) { ring = ring.flatMap((p) => p.friends.filter((f) => !seen.has(f) && seen.add(f))); dist++; }
```

**Go building blocks:**
- a queue of `Cell`s
- a `dist map[Cell]int` that doubles as the visited set
- the four direction vectors

### NumIslands — counting connected components

**The concept:** scan every cell. When you find land you haven't visited, that's a **new island**: count it, then **flood fill** (DFS or BFS) to mark every cell connected to it as visited, so it's never counted again.

```js
// JS: the "paint bucket" fill tool is the same flood fill
const fill = (x, y) => { if (!inBounds(x, y) || px(x, y) !== target) return; setPx(x, y, colour); fill(x+1, y); fill(x-1, y); fill(x, y+1); fill(x, y-1); };
```

**Go building blocks:** nested loops over rows and columns, a `visited map[Cell]bool`, and a recursive flood fill. (Strings are immutable, so you can't "sink" land by overwriting it. The visited set does that job.)

### TopologicalSort — Kahn's algorithm

**The concept:** order tasks so every dependency comes before what needs it (build systems, `npm install`, course prerequisites). **Kahn's algorithm:**
1. Count each node's unmet dependencies.
2. Start a queue with every node that has none.
3. Take one from the queue, emit it, and decrement the count of everything that depends on it. Anything that hits 0 joins the queue.

If you finish having emitted fewer nodes than exist, the rest are stuck in a **cycle**.

```
dependsOn: app→[http db]  http→[log]  db→[log config]  log→[config]
ready: [config] → emit config → log, db lose a dependency → log ready → emit log → ...
```

**Go building blocks:**
- collect every node (keys *and* the nodes named in the values)
- `remaining := map[string]int{}`
- a reversed map, `neededBy := map[string][]string{}`
- a queue, and `ErrCycle` when `len(order) < len(nodes)`

### Dijkstra — cheapest paths with weights

**The concept:** BFS assumes every edge costs 1. With weights, always expand the **cheapest unfinished node next**. Once a node is the cheapest unfinished one, no cheaper route to it can exist (weights are non-negative), so its distance is final. Then **relax** its edges: if going through it makes a neighbour cheaper, update the neighbour.

```
from A:  A=0 → relax: C=1, B=4  → finish C (1) → relax: B=min(4, 1+2)=3, D=8 → finish B (3) → D=min(8, 3+1)=4
```

**Go building blocks:**
- `dist := map[string]int{source: 0}` and `done := map[string]bool{}`
- each round, scan for the cheapest not-done node (that's O(V²), which is fine for now)
- after algorithms 06, swap the scan for a priority queue
