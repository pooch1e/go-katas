# 06 · Functions, Closures & the `slices`/`maps` packages

**Builds on:** 01 (slices + aliasing), 02 (maps), 03 (structs).

```sh
go test ./language/06-functions/
```

Go has no `.map().filter().reduce()` chains on slices. Instead you get plain `for` loops (the idiomatic default), the `slices`/`maps`/`cmp` packages for the common jobs, and generics so you can write `Map` yourself.

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `(x) => x * 2` | `func(x int) int { return x * 2 }` | no arrow shorthand; parameter and return types are spelled out |
| `arr.map(f)` | your `Map(xs, f)`, or a `for` loop | Go doesn't ship one. You'll write it |
| `arr.sort((a, b) => a - b)` | `slices.SortFunc(xs, func(a, b T) int { ... })` | sorts **in place**, like JS |
| `a - b` comparator | `cmp.Compare(a, b)` | returns -1/0/1 for any ordered type |
| `cmpA \|\| cmpB` (tie-break) | `cmp.Or(cmpA, cmpB)` | returns the first non-zero result |
| `[...arr]` | `slices.Clone(xs)` | |
| `arr.includes(x)` | `slices.Contains(xs, x)` | |
| `arr.findIndex(f)` | `slices.IndexFunc(xs, f)` | -1 when missing, same as JS |
| `_.maxBy(arr, f)` | `slices.MaxFunc(xs, cmpFn)` | ⚠️ panics on an empty slice |
| `Object.keys(m)` | `slices.Collect(maps.Keys(m))` | `maps.Keys` returns an *iterator* (see module 10) |
| closures | closures | identical. Inner functions capture outer variables by reference |

## New syntax you'll meet

- **Function types.** `func(int) bool` is a type like any other. You can store it in a variable, pass it, or return it.
- **Reading generic signatures.** `func Map[T, U any](xs []T, f func(T) U) []U` means "for any types `T` and `U`, take a slice of `T` and a function from `T` to `U`, and return a slice of `U`". The callers never write `[int, string]`, because Go infers it. Module 09 covers *writing* constraints.
- **Closures capture variables, not values.** A closure that does `count++` updates the real `count` in the enclosing function, exactly like JS.
- **`slices.SortFunc` comparator.** It returns negative/zero/positive, like JS. `cmp.Compare(a.Age, b.Age)` saves you from writing `a.Age - b.Age` (which can overflow).

---

## The katas

### Map / Filter / Reduce — writing the classics

**The concept:** higher-order functions. A function that takes *another* function to decide what to do with each element. Once you've written these three you'll never wonder what `.map` does under the hood again.

```js
// JS: a hand-written forEach, the simplest higher-order function
function forEach(arr, fn) {
  for (let i = 0; i < arr.length; i++) fn(arr[i], i);
}
```

**Go building blocks:**
- `make([]U, len(xs))` for Map (you know the length up front)
- `make([]T, 0)` and `append` for Filter (you don't)
- an accumulator variable for Reduce

### Chunk — slicing into windows

**The concept:** walk the input in steps of `size`, taking `xs[i : min(i+size, len(xs))]` each time. The subtle part is ownership: sub-slices share memory with the input.

```js
// JS: paginating results
const page = (items, n, size) => items.slice(n * size, (n + 1) * size);
```

**Go building blocks:**
- a `for` loop that steps by `size` instead of 1
- the built-in `min`, so the last chunk doesn't run off the end
- full slice expressions `xs[lo:hi:max]`, or `slices.Clone` each chunk

### MakeCounter — closures hold state

**The concept:** a function that returns a function, where the inner function keeps access to a variable from the outer one. Each call to the outer function creates a fresh, private variable.

```js
// JS: a once() wrapper. The closure remembers whether fn already ran.
const once = (fn) => { let done = false; return (...a) => done ? undefined : (done = true, fn(...a)); };
```

**Go building blocks:** declare a variable, then `return func() int { ... }`, which uses and updates it.

### SortedByAgeThenName — multi-key sorting

**The concept:** sort by a primary key, and when two items tie, fall back to a secondary key. Also: don't sort the caller's data out from under them.

```js
// JS: sort products by price, then by name
[...products].sort((a, b) => a.price - b.price || a.name.localeCompare(b.name));
```

**Go building blocks:**
- `slices.Clone`
- `slices.SortFunc`
- `cmp.Or(cmp.Compare(...), cmp.Compare(...))`

### Oldest — max-by with a "maybe" answer

**The concept:** find the item with the biggest key, and tell the caller honestly when there's no answer.

```js
// JS: newest post, or null
const newest = posts.length ? posts.reduce((a, b) => (b.date > a.date ? b : a)) : null;
```

**Go building blocks:**
- `len(xs) == 0` guard
- `slices.MaxFunc(xs, func(a, b Person) int { ... })`, or a loop
- return `(Person{}, false)` for the empty case

### CountByCity — group + count

**The concept:** module 02's frequency map, keyed by a struct field instead of a letter.

```js
// JS: count orders per status
const byStatus = Object.groupBy(orders, (o) => o.status); // then take lengths
```

**Go building blocks:** `make(map[string]int)`, `m[p.City]++`.

### AverageAge — numeric conversion

**The concept:** sum, then divide, converting to floating point *before* dividing so you don't lose the fraction.

```js
// JS: every number is a float, so this "just works"
const avgScore = scores.reduce((a, b) => a + b, 0) / scores.length;
```

**Go building blocks:**
- `float64(total) / float64(len(people))`
- your own `Reduce` from earlier
- return `(0, false)` for no people
