# Go Katas

Learn Go test-first, one data structure at a time, then put it to work on classic algorithms. It's the same loop as vitest katas: red, green, refactor. It uses Go's built-in test runner and [testify](https://github.com/stretchr/testify) for vitest-style assertions.

```
language/     ← 12 modules: Go itself, taught through data structures
algorithms/   ←  7 modules: sorting, searching, graphs, heaps, DP…
```

Each module folder holds everything for that topic, side by side:

```
language/01-slices/
├── README.md              ← read first: JS → Go mapping + a section per kata
├── slices.go              ← your code (every function starts as a stub returning a zero value)
├── sum_all_test.go        ← one test file per kata
└── find_max_test.go
```

## Setup

You need Go, plus a couple of tools. With [mise](https://mise.jdx.dev):

```sh
mise install       # reads mise.toml: Go, gotestsum, gopls, dlv
go mod download    # fetches testify
```

**New to Go?** Read [HOW-GO-WORKS.md](HOW-GO-WORKS.md) first. It covers compiling, packages, `go.mod`, where `console.log` went, and debugging in Zed.

## The loop

```sh
gotestsum --watch -- -timeout 10s                   # like `vitest`: reruns the package you just saved
go test ./language/01-slices/                       # one module, once
go test ./language/01-slices/ -run TestFindMax      # one kata          (it.only)
go test ./language/01-slices/ -run 'TestFindMax/negative'   # one case
go test -v ./language/01-slices/                    # list every test, including SKIPs
go test ./...                                       # everything
```

`-run` takes a regex matched against `TestName/subtest_name` (spaces in subtest names become `_`). It works with `gotestsum --watch -- -run TestFindMax` too. The `-timeout 10s` matters in the algorithms track: a correct-but-exponential solution would otherwise block the watcher for Go's default 10 minutes.

1. **Red.** Run the kata's tests and watch them fail. The stubs return zero values, so you get readable assertion failures, not crashes. (The odd case passes by accident, e.g. summing an empty slice to 0.)
2. **Green.** Replace the stub body with a real implementation. Make it pass however you can.
3. **Refactor.** Tidy up while green. Then try a second approach: a loop vs the `slices` package, recursion vs iteration.

### Two kinds of test

- **Written tests.** Read them to learn testify and Go test syntax, then make them pass.
- **Todo tests.** Subtests containing only `t.Skip("todo")`, the equivalent of `it.todo`. Delete the `t.Skip` line, write the assertions (copy the shape of a written test in the same file), watch it go red, then implement. Many have a hint comment above them pointing at the gotcha.

The share of todo tests grows as you go. The algorithms track is roughly half todo.

> **Compile errors are per package.** A typo in `language/04-strings` stops only that module's tests, not the whole repo. But inside a module, one compile error stops every test in it.

## Modules

Do the language track in order. The algorithms track assumes you've done it (at least through 11). **Before starting a module, read its `README.md`.** Each kata gets a section explaining the *concept* with a JS example of the same idea on a different problem, plus the Go building blocks you'll need. It never gives the solution.

### Language track

| # | Module | You'll learn | Builds on |
|---|--------|--------------|-----------|
| 01 | [Slices & Arrays](language/01-slices/README.md) | `[]int`, `len`, `range`, `append`, multiple returns, sentinel errors, **slices share memory** | — |
| 02 | [Maps](language/02-maps/README.md) | `map[K]V`, comma-ok, zero values, random iteration order, `rune` | 01 |
| 03 | [Structs & Methods](language/03-structs/README.md) | structs, value vs pointer receivers, `NewX` constructors, unexported fields | 01, 02 |
| 04 | [Strings & Runes](language/04-strings/README.md) | bytes vs runes, UTF-8, `strings.Builder`, `strconv`, `unicode` | 01, 02 |
| 05 | [Errors](language/05-errors/README.md) | wrapping with `%w`, `errors.Is`/`As`, custom error types, `errors.Join`, `defer`/`recover` | 01–03 |
| 06 | [Functions & `slices`/`maps`](language/06-functions/README.md) | func values, closures, generic Map/Filter/Reduce, `slices.SortFunc`, `cmp.Or` | 01–03 |
| 07 | [Sets](language/07-sets/README.md) | `map[T]struct{}`, generic `Set[T comparable]`, seen-set pattern | 02, 06 |
| 08 | [Stacks & Queues](language/08-stacks-queues/README.md) | slices as stacks/queues, `switch`, `strconv.Atoi` | 01, 04, 05 |
| 09 | [Interfaces & Generics](language/09-interfaces-generics/README.md) | implicit interfaces, type switches, `fmt.Stringer`, constraints, `Stack[T]` | 03, 06–08 |
| 10 | [Linked Lists](language/10-linked-lists/README.md) | pointer-linked nodes, `nil` safety, **range-over-func iterators**, Floyd's cycle check | 03, 09 |
| 11 | [Trees](language/11-trees/README.md) | recursion, generic BST, recursive iterators, BFS by level | 08–10 |
| 12 | [Concurrency](language/12-concurrency/README.md) | goroutines, channels, `WaitGroup`, `select`, `context`, the race detector | 05, 06 |

### Algorithms track

| # | Module | Katas | Key idea |
|---|--------|-------|----------|
| 01 | [Sorting](algorithms/01-sorting/README.md) | Bubble, Insertion, Merge, Quick | O(n²) vs O(n log n), stability, benchmarks |
| 02 | [Binary Search](algorithms/02-binary-search/README.md) | BinarySearch, SearchInsert, FirstBadVersion, SearchRotated | halving; boundaries; searching an answer space |
| 03 | [Two Pointers & Sliding Window](algorithms/03-two-pointers/README.md) | TwoSumSorted, RemoveDuplicatesSorted, MaxSumSubarrayK, LongestUniqueSubstring | O(n²) → O(n) |
| 04 | [Recursion & Backtracking](algorithms/04-recursion/README.md) | Fibonacci, Permutations, Subsets, CombinationSum | memoisation; choose / recurse / un-choose |
| 05 | [Graphs](algorithms/05-graphs/README.md) | HasPath, ShortestPathGrid, NumIslands, TopologicalSort, Dijkstra | DFS, BFS, visited sets, Kahn's algorithm |
| 06 | [Heaps](algorithms/06-heaps/README.md) | MinHeap[T], IntHeap, KthLargest, TopKFrequent, MergeKSorted | priority queues, `container/heap` |
| 07 | [Dynamic Programming](algorithms/07-dynamic-programming/README.md) | ClimbStairs, CoinChange, LCS, Knapsack | state, recurrence, base cases, fill order |

## JS → Go cheat sheet

| JS | Go |
|----|----|
| `const x = 5` / `let x = 5` | `x := 5` (inferred) or `var x int = 5`. `const` is compile-time only |
| `[1, 2, 3]` | `[]int{1, 2, 3}` |
| `arr.length` | `len(arr)` |
| `arr.push(x)` | `arr = append(arr, x)` |
| `for (const x of arr)` | `for _, x := range arr` |
| `{ a: 1 }` as a map | `map[string]int{"a": 1}` |
| `{ x: 1, y: 2 }` as a shape | `type Point struct{ X, Y int }` → `Point{X: 1, Y: 2}` |
| `obj.key ?? fallback` | `v, ok := m[key]; if !ok { v = fallback }` |
| `(x) => x * 2` | `func(x int) int { return x * 2 }` |
| `` `Hi ${name}` `` | `fmt.Sprintf("Hi %s", name)` |
| `===` | `==` (strings, numbers and simple structs compare by value) |
| `undefined` / `null` | zero values (`0`, `""`, `false`), and `nil` for pointers, slices, maps, interfaces |
| `throw new Error("x")` | `return errors.New("x")`, so errors are returned values |
| `try { } catch { }` | `if err != nil { }` |
| `export function f` | `func F`, where a capital letter means exported |
| `console.log(x)` | `fmt.Println(x)` / `fmt.Printf("%+v\n", x)` / `t.Log(x)` in tests |
| `Promise.all` | goroutines + `sync.WaitGroup` |
| `expect(a).toBe(5)` | `assert.Equal(t, 5, a)` (expected first) |
| `expect(a).toEqual([1, 2])` | `assert.Equal(t, []int{1, 2}, a)`, which is deep equality |
| `expect(arr).toEqual(expect.arrayContaining(...))` | `assert.ElementsMatch(t, want, got)` (order-insensitive) |
| `expect(a).toBeCloseTo(0.3)` | `assert.InDelta(t, 0.3, a, 1e-9)` |
| `expect(fn).toThrow(NotFound)` | `assert.ErrorIs(t, err, ErrNotFound)` |
| `expect(err).toBeInstanceOf(VErr)` | `var ve *VErr; assert.ErrorAs(t, err, &ve)` |
| `it.each([...])` | a table of cases + `for` + `t.Run(tc.name, ...)` |
| `it.todo("...")` | `t.Run("...", func(t *testing.T) { t.Skip("todo") })` |

**`assert` vs `require`:** `assert.X` records a failure and keeps going. `require.X` stops the test immediately. Use `require` when later lines would be meaningless (or would crash) after a failure, e.g. `require.NoError(t, err)` before using the result.

## Asserting errors

`assert.Error(t, err, "msg")` only checks that `err` is non-nil. The last argument is a label printed on failure, **not** something matched against the error. Pick the assertion that matches what you actually want to check:

| You want to check | Use |
|---|---|
| Any error occurred | `assert.Error(t, err)` |
| It is (or wraps) a specific sentinel like `ErrOutOfRange` | `assert.ErrorIs(t, err, ErrOutOfRange)` |
| It is (or wraps) a custom error type | `assert.ErrorAs(t, err, &target)` |
| The message contains some text | `assert.ErrorContains(t, err, "out of range")` |
| The message matches exactly | `assert.EqualError(t, err, "index out of range")` |
| No error occurred | `assert.NoError(t, err)` |

Prefer `ErrorIs` for sentinels. It uses `errors.Is`, so it still passes if the function later wraps the error with `fmt.Errorf("...: %w", ErrOutOfRange)`, whereas `EqualError` and `ErrorContains` break when the wording changes.

```go
t.Run("returns ErrOutOfRange for a negative index", func(t *testing.T) {
    _, err := RemoveAt([]int{1, 2, 3}, -1)
    assert.ErrorIs(t, err, ErrOutOfRange)
})
```

Use `ErrorAs` when you define your own error struct and want to inspect its fields:

```go
var target *IndexError
require.ErrorAs(t, err, &target)
assert.Equal(t, -1, target.Index)
```

Use `require.*` instead of `assert.*` when later lines depend on the error being present.
