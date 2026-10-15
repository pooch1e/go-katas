# 01 · Slices & Arrays

**Builds on:** nothing. Start here.

```sh
go test ./language/01-slices/
gotestsum --watch -- ./language/01-slices/
```

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `const nums = [1, 2, 3]` | `nums := []int{1, 2, 3}` | a **slice**: growable, and every element is an `int` forever |
| `nums.push(4)` | `nums = append(nums, 4)` | `append` *returns* the slice; you must reassign it |
| `nums.length` | `len(nums)` | a built-in function, not a property |
| `new Array(5)` | `make([]int, 5)` | starts as `[0 0 0 0 0]`. Every type has a **zero value**, never `undefined` |
| `for (const n of nums)` | `for _, n := range nums` | `range` gives `index, value`. `_` throws away the one you don't need |
| `for (let i = 0; i < n; i++)` | `for i := 0; i < n; i++` | Go has only `for`. No `while`, no parentheses |
| `nums.slice(1, 3)` | `nums[1:3]` | **shares memory** with `nums` (see below) |
| `[[1, 2], [3]]` | `[][]int{{1, 2}, {3}}` | a slice of slices |
| `return undefined` on failure | `return 0, ErrEmpty` | functions can return several values; errors are just values |

## New syntax you'll meet

- **Arrays vs slices.** `[3]int` is an *array*: fixed size, the size is part of the type, and it's copied when you pass it around. You'll rarely use them. `[]int` (no number) is a *slice*: a small header pointing into an array somewhere. That's what you'll use almost always.
- **`:=` vs `var`.** `total := 0` declares and infers the type. `var total int` declares with the zero value. Both are fixed-type forever.
- **Exported names.** `SumAll` starts with a capital, so other packages (including the tests) can call it. `sumAll` would be private to this package. That's Go's whole `export` system.
- **Multiple return values.** `func FindMax(nums []int) (int, error)` returns two things. The caller writes `max, err := FindMax(nums)` and checks `if err != nil`.
- **Sentinel errors.** `var ErrEmpty = errors.New("empty slice")` is a package-level error value you return and callers compare against with `errors.Is`. Module 05 goes much deeper.
- **nil slices.** A `var s []int` is `nil`, but `len(s)` is 0 and `range` over it just does nothing. You rarely need to special-case nil.

## The big gotcha: slices share memory

```go
a := []int{1, 2, 3, 4}
b := a[0:2]   // b is [1 2], but it points INTO a
b[0] = 99     // a is now [99 2 3 4]
```

In JS, `arr.slice()` always copies. In Go, slicing **never** copies. It makes a new view onto the same backing array. To get a real copy, make a new slice and `copy` into it, or use `slices.Clone(a)`.

---

## The katas

### SumAll — accumulate over a slice

**The concept:** walk every element once, carrying a running total. It's the simplest *reduce*.

```js
// JS: the same idea on a different problem, counting characters across words
let totalChars = 0;
for (const word of ["go", "is", "fun"]) totalChars += word.length;
```

**Go building blocks:** `for _, v := range xs { ... }`, `+=`, and zero values (an `int` you declare starts at `0`).

### DoubleAll — build a new slice without touching the input

**The concept:** *map*. Produce a new collection of the same length, transforming each element. The input must survive unchanged.

```js
// JS: map makes the new array for you
const lengths = ["go", "is", "fun"].map((w) => w.length);
```

**Go building blocks:** `make([]T, length)` allocates a slice of a known size up front. Then index into it: `out[i] = ...`. (Or start with `make([]T, 0, capacity)` and `append`.) The tests check that empty input gives `[]int{}`, not `nil`.

### FindMax — "best so far", and what to do with no answer

**The concept:** track the best value seen so far. The trick is the **starting value**: it must be something that can't beat a real element by accident.

```js
// JS: the same pattern, finding the longest word
let longest = words[0];
for (const w of words) if (w.length > longest.length) longest = w;
```

**Go building blocks:**
- `if len(nums) == 0 { return 0, ErrEmpty }`. Return the zero value alongside the error, which is the convention.
- `nums[0]` for the first element.
- `return best, nil` on success.

### ReverseInPlace — two pointers that meet in the middle

**The concept:** mutate the caller's data instead of building a new copy. One index starts at the front and one at the back. Swap, step both inwards, and stop when they meet. The function returns nothing; the caller sees the change because a slice points at shared memory.

```js
// JS: two pointers checking symmetry (not swapping)
let i = 0, j = s.length - 1;
while (i < j && s[i] === s[j]) { i++; j--; }
```

**Go building blocks:**
- Go's `for` can declare and update two variables at once: `for a, b := 0, 10; a < b; a, b = a+1, b-1 { }`.
- Swapping needs no temp variable: `x, y = y, x` works on slice elements too.

### CountOccurrences — filter + count

**The concept:** loop, test a condition, increment. Strings compare by value with `==`, exactly like JS `===`.

```js
// JS: counting evens
let evens = 0;
for (const n of nums) if (n % 2 === 0) evens++;
```

**Go building blocks:** `range`, `==` on strings, and `count++` (which is a statement in Go, not an expression, so you can't write `x = count++`).

### Flatten — nested loops into one output

**The concept:** an outer loop over the groups and an inner loop over each group's items, all pushed into a single result.

```js
// JS: collecting every tag from a list of posts
const tags = [];
for (const post of posts) for (const tag of post.tags) tags.push(tag);
```

**Go building blocks:**
- `append(out, x)` adds one element.
- `append(out, other...)` adds a whole slice. The `...` is like JS spread.
- Bonus: work out the total length first and `make` the slice with that capacity, so `append` never has to reallocate.

### RemoveAt — the `append` aliasing trap

**The concept:** return a copy that's missing one element, and **prove** the original is untouched. This kata exists to make the "slices share memory" gotcha bite you once, on purpose, in a test.

```js
// JS: toSpliced copies. splice would mutate.
const withoutSecond = arr.toSpliced(1, 1);
```

**Go building blocks:**
- Bounds checks first, returning `nil, ErrOutOfRange`.
- Then build a fresh slice (`make`) and copy the two halves in, using `copy(dst, src)` or `append`.
- Try the one-liner in the test's hint comment *first*, watch the "does not mutate" test go red, then fix it. That's the lesson.

The last three tests are **todo**. Delete `t.Skip("todo")` and write the assertions using the written tests above as a pattern.
