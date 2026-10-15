# Algorithms 04 · Recursion & Backtracking

**Builds on:** language 11 (recursive helpers), algorithms 01 (MergeSort's divide and conquer).

```sh
go test -timeout 10s ./algorithms/04-recursion/
```

## Complexity

| Kata | Time | Why |
|------|------|-----|
| Fibonacci (naive) | O(2ⁿ) | every call makes two more |
| Fibonacci (memo) | **O(n)** | each `n` is computed once, then looked up |
| Permutations | O(n · n!) | there are n! answers, each of length n |
| Subsets | O(n · 2ⁿ) | each element is in or out |
| CombinationSum | exponential | bounded by pruning |

When the output itself is exponential, no algorithm can beat that. The skill is not doing *extra* work.

## The backtracking template

Every kata after Fibonacci has the same skeleton:

```
explore(path, choicesLeft):
    if path is a complete answer: record a COPY of path
    for each choice:
        add choice to path      ← choose
        explore(path, ...)      ← recurse
        remove choice from path ← un-choose (backtrack)
```

What changes between katas is only **what counts as complete** and **which choices are allowed next**.

## Go you'll lean on

- **Recursive closures.** Declare first, then assign, so the function can call itself:
  ```go
  var explore func(path []int)
  explore = func(path []int) { ... explore(...) ... }
  ```
- **Memo as a map.** `memo := map[int]int{}`, checked with comma-ok before computing.
- **Choose/un-choose on a slice.** `path = append(path, x)` … `path = path[:len(path)-1]`.
- **Copy before recording.** `results = append(results, slices.Clone(path))`. Without the clone, every recorded answer aliases the same backing array (language 01's lesson, at full strength).
- **Used-markers.** A `[]bool` the same length as the input, for "is this element already in my permutation?"

---

## The katas

### Fibonacci — recursion, then memoisation

**The concept:** the recursive definition (`fib(n) = fib(n-1) + fib(n-2)`) recomputes the same values exponentially often. A **memo** caches each answer the first time, so the tree of calls collapses into a line. This is top-down dynamic programming, and module 07 picks it up.

```js
// JS: memoising an expensive pure function in general
const memo = (fn) => { const cache = new Map(); return (x) => cache.has(x) ? cache.get(x) : (cache.set(x, fn(x)), cache.get(x)); };
```

**Go building blocks:** a map memo plus a recursive closure, or a helper that takes the map as a parameter.

### Permutations — choose from what's unused

**The concept:** at each position, try every element not already used. Recurse to fill the next position, then un-choose so the next iteration starts clean. A path of length `n` is complete.

```
[]  → [1] → [1 2] → [1 2 3] ✓   back to [1] → [1 3] → [1 3 2] ✓   back to [] → [2] → ...
```

**Go building blocks:** a `used []bool`, choose/un-choose on both `path` and `used`, and `slices.Clone` when recording.

### Subsets — in or out

**The concept:** every element is a binary decision: include it or not. Equivalently: record the current path at *every* node of the recursion (not only at leaves), and only choose elements *after* the last one chosen.

```
start=0, path [] ✓ → take 1: [1] ✓ → take 2: [1 2] ✓ → take 3: [1 2 3] ✓
                                   → take 3: [1 3] ✓
                   → take 2: [2] ✓ → ...
```

**Go building blocks:** a `start int` parameter so you never look backwards, and record `slices.Clone(path)` at the top of each call.

### CombinationSum — reuse allowed, prune the overshoot

**The concept:** like Subsets, but you may pick the same candidate again (so recurse with the *same* start index, not `start+1`). Track the remaining target. When it hits 0, record; below 0, stop (prune). Sorting candidates first lets you `break` as soon as one is too big.

```js
// JS: the neighbouring idea, all the ways to make change, counted rather than listed
const ways = (coins, amt, i = 0) => amt === 0 ? 1 : amt < 0 || i === coins.length ? 0 : ways(coins, amt - coins[i], i) + ways(coins, amt, i + 1);
```

**Go building blocks:** `slices.Sort(candidates)` (on a clone), plus a `remaining int` parameter with base cases for `== 0` and `< 0`.
