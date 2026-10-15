# 11 · Trees

**Builds on:** 08 (queues), 09 (generics, `cmp.Ordered`), 10 (pointer-linked nodes, iterators).

```sh
go test ./language/11-trees/
```

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `{ value, left: null, right: null }` | `&TreeNode[T]{Value: v}` | `Left`/`Right` start as `nil` |
| `function insert(node, v) { ... }` | `func insert[T cmp.Ordered](n *TreeNode[T], v T) *TreeNode[T]` | recursive helpers are plain (unexported) functions |
| `function* inOrder(n) { yield* inOrder(n.left); yield n.value; ... }` | a recursive helper that takes `yield` and returns `bool` | there's no `yield*`, so you pass "keep going?" back up |
| `Math.max(a, b)` | `max(a, b)` | built in since Go 1.21 |

## New syntax you'll meet

- **Recursion on pointers.** The base case is almost always `if n == nil { ... }`. A common shape is a public method that calls a private recursive helper with the root: `func (t *BST[T]) Height() int { return height(t.root) }`.
- **Returning the new subtree.** For insert, a neat Go idiom is `n.Left = insert(n.Left, v)`. The helper returns the (possibly new) node, so you don't need pointer-to-pointer gymnastics.
- **Recursive iterators.** Write `func walk(n *TreeNode[T], yield func(T) bool) bool` that returns `false` as soon as `yield` does, and have every recursive call check it: `if !walk(n.Left, yield) { return false }`.
- **Generic helper functions** need the type parameter too: `func height[T cmp.Ordered](n *TreeNode[T]) int`.

---

## The katas

### Insert / Contains / Min / Max — following the ordering rule

**The concept:** a **binary search tree** halves the search space at every step (when it's balanced). Compare with the current node, go left if smaller, right if bigger. Min is "go left until you can't"; Max is the mirror.

```js
// JS: the same left/right decision, as a guessing game
let lo = 1, hi = 100;
while (lo <= hi) {
  const guess = (lo + hi) >> 1;
  if (guess === secret) break;
  guess < secret ? (lo = guess + 1) : (hi = guess - 1);
}
```

**Go building blocks:**
- `cmp.Compare(v, n.Value)` gives you -1/0/1 to `switch` on
- loops or recursion both work, so try one of each
- `Insert` returns `false` on a duplicate, without changing `size`

### Height — recursive depth

**The concept:** a tree's height is 1 + the taller of its two subtrees. An empty tree has height 0. This "combine the answers from the children" shape is the template for most tree problems.

```js
// JS: the same recursion, measuring how deeply nested a JSON value is
const depth = (v) =>
  typeof v !== "object" || v === null ? 0 : 1 + Math.max(0, ...Object.values(v).map(depth));
```

**Go building blocks:** a generic `height` helper with a `nil` base case, and the built-in `max`.

### InOrder — depth-first traversal as an iterator

**The concept:** an **in-order** traversal (left, self, right) of a BST visits values in sorted order. Make it lazy and early-exit-safe like module 10's `All`.

```js
// JS: depth-first walk of a folder tree with a generator
function* files(dir) {
  for (const child of dir.children) child.isDir ? yield* files(child) : yield child.name;
}
```

**Go building blocks:** the recursive `walk(n, yield) bool` shape described above.

### LevelOrder — breadth-first with a queue

**The concept:** **breadth-first** visits the tree level by level. A queue holds the nodes waiting to be visited. To group by level, process exactly `len(queue)` nodes per round; those are the current level.

```js
// JS: degrees of separation in a social graph, one ring at a time
let ring = [me], seen = new Set([me]);
while (ring.length) {
  const next = [];
  for (const p of ring) for (const f of p.friends) if (!seen.has(f)) { seen.add(f); next.push(f); }
  ring = next;
}
```

**Go building blocks:**
- `queue := []*TreeNode[T]{t.root}`
- an outer `for len(queue) > 0`, with an inner loop over the current level's size
- append children to the queue for the next round

### IsValidBST — passing bounds down

**The concept:** checking each node against only its children isn't enough (see the grandchild test). Instead, pass down the **range of values allowed** at each position. Going left tightens the upper bound; going right tightens the lower bound.

```js
// JS: validating nested time ranges, where each child must sit inside its parent's window
const valid = (ev, lo = -Infinity, hi = Infinity) =>
  ev.start >= lo && ev.end <= hi && ev.children.every((c) => valid(c, ev.start, ev.end));
```

**Go building blocks:**
- a helper taking `lo, hi *int` (with `nil` meaning "no bound"), or `math.MinInt` / `math.MaxInt`
- strict `<` and `>`, so duplicates fail
