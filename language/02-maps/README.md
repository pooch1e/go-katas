# 02 · Maps

**Builds on:** 01 (slices, `range`, the comma-ok idea of returning two values).

```sh
go test ./language/02-maps/
```

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `const m = new Map()` / `{}` as a map | `m := map[string]int{}` or `make(map[string]int)` | key and value types are fixed |
| `m.get(k)` / `obj[k]` | `m[k]` | a missing key gives the **zero value**, never `undefined` |
| `m.has(k)` | `v, ok := m[k]` | the **comma-ok** form: `ok` is `false` when the key is missing |
| `m.set(k, v)` | `m[k] = v` | |
| `m.delete(k)` | `delete(m, k)` | a built-in, safe even if the key isn't there |
| `m.size` | `len(m)` | |
| `for (const [k, v] of m)` | `for k, v := range m` | order is **deliberately random** every run |
| `obj.key ?? fallback` | `if v, ok := m[k]; ok { ... }` | `if` can start with a short statement |

## New syntax you'll meet

- **Writing to a nil map panics.** `var m map[string]int` is `nil`. Reading from it is fine (you get zero values), but `m["x"] = 1` crashes. Always `make` or use a literal `{}` before writing.
- **Zero values make counting easy.** `counts[k]++` works on a missing key because `counts[k]` starts at `0`. No `(counts[k] ?? 0) + 1` dance.
- **`rune`** is Go's type for a single Unicode character (it's an `int32`). Ranging over a string gives you runes: `for _, r := range "héllo"`. Module 04 covers this properly; for now `unicode.IsLetter(r)` and `unicode.ToLower(r)` are all you need.
- **Maps are reference types.** Passing a map to a function passes a handle to the same data, like a JS object. Changes inside the function are visible to the caller.
- **Random iteration order.** Go randomises `range` order over a map on purpose, so you can't come to depend on it. If you need order, collect the keys into a slice and sort them (`slices.Sorted(maps.Keys(m))`).

---

## The katas

### CountLetters — frequency counting

**The concept:** a frequency map. Walk the input once and bump a counter per key. It's one of the most-used patterns in all of DSA (anagrams, duplicates, top-k…).

```js
// JS: counting word frequency
const freq = {};
for (const w of words) freq[w] = (freq[w] ?? 0) + 1;
```

**Go building blocks:**
- `make(map[rune]int)`
- `for _, r := range s`
- `unicode.IsLetter`, `unicode.ToLower`
- `m[k]++` works with no initialisation, thanks to the zero value

### LookupAge — "missing" vs "zero"

**The concept:** a map lookup can't tell "stored 0" from "not there" by value alone. You need a second signal.

```js
// JS: the same trap with ||
const port = config.port || 8080; // breaks when port is 0
const safe = "port" in config ? config.port : 8080;
```

**Go building blocks:** the comma-ok form of a map read, and `if` with a short statement (`if x, ok := ...; ok { }`). Reading from a `nil` map is safe.

### GroupByLength — map of slices

**The concept:** grouping. The key is something you compute from each item; the value is a list that grows.

```js
// JS: grouping people by city
const byCity = {};
for (const p of people) (byCity[p.city] ??= []).push(p.name);
```

**Go building blocks:**
- `m[k] = append(m[k], item)`. This works even when `m[k]` is missing, because the zero value of a slice is `nil` and `append` on `nil` is fine.
- `len(word)` gives bytes, which is fine for ASCII words.

### Invert — building a new map, detecting collisions

**The concept:** swap keys and values, but refuse when two keys share a value. Silently losing data is a bug.

```js
// JS: building a lookup from email → user, refusing duplicates
const byEmail = new Map();
for (const u of users) {
  if (byEmail.has(u.email)) throw new Error("duplicate email");
  byEmail.set(u.email, u);
}
```

**Go building blocks:**
- `make(map[string]string, len(m))`, which pre-sizes the map as a hint
- the comma-ok check before writing
- `return nil, ErrDuplicateValue`

### MergeCounts — combine without mutating

**The concept:** build a fresh result from two inputs, so neither input changes.

```js
// JS: merging settings, with later sources winning
const merged = { ...defaults, ...userPrefs };
```

**Go building blocks:**
- Go has no map spread. Use a `range` over each input, writing into a new `make`d map.
- `maps.Copy(dst, src)` exists, but it overwrites rather than adds, so it's not quite what you want here.
