# Algorithms 02 · Binary Search

**Builds on:** language 01 (indexes, loops), algorithms 01 (the "halving gives log n" idea).

```sh
go test ./algorithms/02-binary-search/
```

## Complexity

| | Time | Space |
|-|------|-------|
| Every kata here | **O(log n)** | O(1) |

log₂(1,000,000) ≈ 20: a million-element slice needs about 20 steps.

## The template

The whole module is variations on one loop. The hard part is never the idea; it's the **boundaries**:
- `<` or `<=` in the loop condition?
- `mid` or `mid+1` when moving `lo`?
- what to return when the loop ends?

Pick one convention and stick to it:

| Convention | Loop | Meaning |
|------------|------|---------|
| closed `[lo, hi]` | `for lo <= hi` | `hi` starts at `len-1`; the answer, if any, is inside |
| half-open `[lo, hi)` | `for lo < hi` | `hi` starts at `len`; when the loop ends, `lo == hi` is the boundary |

Use the closed form for "find it or -1", and half-open for "find the first position where something becomes true" (`SearchInsert`, `FirstBadVersion`).

## Go you'll lean on

- `mid := lo + (hi-lo)/2`. On 64-bit Go, `(lo+hi)/2` won't realistically overflow, but this form is the habit to build.
- The standard library versions, *after* you've written your own: `slices.BinarySearch(xs, x)` returns `(index, found)`, and `sort.Search(n, pred)` finds the first `i` where `pred(i)` is true.

---

## The katas

### BinarySearch — find or -1

**The concept:** in a sorted slice, look at the middle. If it's too small, the answer can only be to the right, so throw away the left half. And vice versa. Each step halves what's left.

```js
// JS: the same idea in real life, git bisect, where every check halves the commits left
// good=1, bad=100 → test 50 → bad → test 25 → good → test 37 → ...
```

**Go building blocks:** the closed-interval template, comparing `xs[mid]` with `target`.

### SearchInsert — the lower bound

**The concept:** find the first index whose value is `>= target`. If the target exists, that's its index; if not, it's where it would go. This "first position where a condition flips to true" is the most reusable form of binary search.

```
xs = [1 3 5 6], target 2:   1<2 ✗   3>=2 ✓ ...   → answer 1
```

**Go building blocks:** the half-open template. Decide which comparison means "the answer is definitely right of `mid`" and which means "`mid` might be the answer, so keep it". When the loop ends, `lo` is the boundary.

### FirstBadVersion — binary search on an answer, not a slice

**The concept:** there's no slice at all, just a yes/no function over `1..n` that flips from false to true exactly once. Binary search works on **any monotonic predicate**. Spotting this is a big DSA unlock ("minimum speed to finish in time", "smallest capacity that ships in D days"…).

```js
// JS: the same shape, finding the smallest font size that overflows a box
// fits(12) ✓  fits(24) ✗  fits(18) ✓  fits(21) ✗ ...
```

**Go building blocks:** the same lower-bound shape as `SearchInsert`, but the range is the version numbers `1..n` and the test is a function call.

### SearchRotated — which half is sorted?

**The concept:** a sorted slice rotated at some point: `[4 5 6 7 0 1 2]`. Split it at `mid` and **one side is always properly sorted**. Check whether the target lies within that sorted side's range. If so, search there; if not, search the other side.

```
[4 5 6 7 0 1 2], target 1, mid=7: left [4..7] is sorted, 1 isn't in 4..7 → go right
```

**Go building blocks:** the closed template, plus `xs[lo] <= xs[mid]` to detect that the left half is sorted.
