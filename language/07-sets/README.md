# 07 · Sets

**Builds on:** 02 (maps, comma-ok), 04 (runes, `strings.Fields`), 06 (reading generic signatures).

```sh
go test ./language/07-sets/
```

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `new Set()` | `map[T]struct{}{}` | Go has no set type. A map with empty values *is* the set |
| `set.add(x)` | `s[x] = struct{}{}` | `struct{}{}` is the empty struct value, and takes zero bytes |
| `set.has(x)` | `_, ok := s[x]` | comma-ok again |
| `set.delete(x)` | `delete(s, x)` | |
| `set.size` | `len(s)` | |
| `new Set(arr)` | a loop that adds each item | or wrap it up in your own `NewSet` |
| `a.intersection(b)` | write it yourself | |

## New syntax you'll meet

- **`struct{}`** is a type with no fields. Its only value is `struct{}{}`. Using it as the map value says "I only care about the keys" and costs no memory. You'll also see `map[T]bool`. It works too, but it's less precise (what does `false` mean?).
- **A generic type.** `type Set[T comparable] struct { items map[T]struct{} }` declares a type with a type parameter. Methods use it as `func (s *Set[T]) Add(item T) bool`.
- **`comparable`** is a built-in constraint: "any type that works with `==`", which map keys must be. Ints, strings, runes and structs of those qualify; slices and maps don't.
- **Explicit type arguments.** `NewSet[int]()` with no arguments can't infer `T`, so you name it.

---

## The katas

### Set[T] — wrapping a map in a type

**The concept:** an **abstract data type**. Hide the representation (a map) behind a small set of operations. Callers say `s.Has(x)`, never `_, ok := s.items[x]`, so you could swap the internals later without breaking anyone.

```js
// JS: a Bag type that hides a Map of counts behind a clear API
class Bag {
  #counts = new Map();
  add(x) { this.#counts.set(x, (this.#counts.get(x) ?? 0) + 1); }
  count(x) { return this.#counts.get(x) ?? 0; }
}
```

**Go building blocks:**
- `make(map[T]struct{})` in the constructor
- variadic `items ...T`
- the comma-ok check inside `Add` so it can return whether the item was new
- iterating the smaller set in `Intersect` is cheaper

### HasDuplicates — the seen-set pattern

**The concept:** the most important trick in this module. Keep a set of everything you've seen. For each new item, check the set *before* adding it. That's O(n) instead of comparing every pair (O(n²)).

```js
// JS: de-duplicating notifications by id while keeping order
const seen = new Set();
const unique = notes.filter((n) => !seen.has(n.id) && seen.add(n.id));
```

**Go building blocks:** `seen := make(map[T]struct{})`, comma-ok, `return true` early.

### CountUniqueWords — normalise, then set

**The concept:** decide what counts as "the same word" (here: case-insensitive, letters only), normalise every word to that form, then let the set do the counting.

```js
// JS: distinct email domains, normalised
new Set(emails.map((e) => e.split("@")[1].toLowerCase())).size;
```

**Go building blocks:**
- `strings.FieldsFunc(text, func(r rune) bool { return !unicode.IsLetter(r) })` splits on anything that isn't a letter
- `strings.ToLower`
- `len(set)`

### FirstRepeatedChar — seen-set with early exit

**The concept:** the same pattern as `HasDuplicates`, but return *which* item repeated first. "First" means the earliest *second* occurrence.

```js
// JS: first user to log in twice
const seen = new Set();
for (const id of logins) { if (seen.has(id)) return id; seen.add(id); }
```

**Go building blocks:** `range` over runes, return `(r, true)`, fall through to `(0, false)`.

### CommonElements — set lookup + ordered output

**The concept:** build a set from one side for O(1) lookups, then walk the other side *in order*. You need a second set to avoid emitting the same value twice.

```js
// JS: which of my friends are also at the event, in my friend-list order
const attending = new Set(eventGuests);
const both = myFriends.filter((f) => attending.has(f));
```

**Go building blocks:**
- two sets (`inB` and `added`)
- `append` into a `make([]int, 0)` result so "nothing shared" gives `[]int{}`

### HasPairWithSum — complement lookup

**The concept:** for each number `n`, the partner you need is `target - n`. Have you *already* seen it? Check *before* adding `n` itself, so you never pair a number with itself.

```js
// JS: does any word have its reverse earlier in the list?
const seen = new Set();
for (const w of words) {
  if (seen.has([...w].reverse().join(""))) return true;
  seen.add(w);
}
```

**Go building blocks:** a `map[int]struct{}` seen-set and one loop. (The algorithms track has a two-pointer version for *sorted* input.)
