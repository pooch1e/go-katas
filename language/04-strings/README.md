# 04 · Strings & Runes

**Builds on:** 01 (slices, two pointers), 02 (maps, `rune`).

```sh
go test ./language/04-strings/
```

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `"héllo".length` → 5 | `len("héllo")` → **6** | `len` counts **bytes**. Go strings are UTF-8 bytes |
| `[..."héllo"].length` | `utf8.RuneCountInString(s)` | counts characters (runes) |
| `s[0]` → `"h"` | `s[0]` → `104` (a `byte`) | indexing gives a byte, not a character |
| `for (const ch of s)` | `for i, r := range s` | `range` over a string decodes runes; `i` is the **byte** offset |
| `[...s]` | `[]rune(s)` | convert to a slice of characters, then index safely |
| `s.toUpperCase()` | `strings.ToUpper(s)` | functions in the `strings` package, not methods |
| `s.split(" ")` | `strings.Split(s, " ")` / `strings.Fields(s)` | `Fields` splits on any run of whitespace |
| `arr.join(", ")` | `strings.Join(parts, ", ")` | |
| `` `${a}-${b}` `` | `fmt.Sprintf("%s-%d", a, b)` | or plain `a + "-" + b` |
| `String(42)` | `strconv.Itoa(42)` | ⚠️ `string(42)` gives `"*"`, the character with code 42 |
| `"a".charCodeAt(0)` | `'a'` | single quotes make a `rune` literal, which is a number |
| `str += x` in a loop | `strings.Builder` | strings are immutable, so `+=` copies every time |

## New syntax you'll meet

- **byte vs rune vs string.** A `string` is an immutable sequence of bytes. A `byte` is one of those bytes (`uint8`). A `rune` is one Unicode code point (`int32`). ASCII characters are one byte; `é` is two; `🚀` is four.
- **Rune arithmetic.** Runes are numbers, so `'c' - 'a'` is `2` and `'a' + 2` is `'c'`. Convert back to a string with `string(r)`.
- **`strings.Builder`.** It works like collecting parts into an array and `join("")`-ing at the end: `var b strings.Builder; b.WriteRune(r); b.WriteString("x"); b.String()`.
- **The `unicode` package.** `unicode.IsLetter`, `IsDigit`, `IsSpace`, `ToUpper` and `ToLower` all take a rune.
- **Modulo of negatives.** `-1 % 26 == -1` in Go (and in JS). To wrap into `0..25`, use `((n % 26) + 26) % 26`.

---

## The katas

### Reverse — characters, not bytes

**The concept:** a string isn't always a list of characters. Turn it into one first, then reverse it. Look at the emoji test case: the bug you're avoiding is *invisible* on ASCII input.

```js
// JS has the same trap: split("") breaks emoji into surrogate halves
"go🚀".split("").length; // 4
[..."go🚀"].length;      // 3
```

**Go building blocks:**
- `[]rune(s)` converts a string into a mutable slice of characters, and `string(...)` converts it back
- module 01 already taught you how to reverse a slice

### IsPalindrome — normalise, then compare

**The concept:** clean the input into a canonical form (lowercase letters and digits only), then compare it with itself backwards. Alternatively, two pointers that skip junk as they go, which avoids building a new string at all.

```js
// JS: normalising before comparing usernames
const norm = (s) => s.toLowerCase().replace(/[^a-z0-9]/g, "");
norm("Ada_Lovelace") === norm("adalovelace");
```

**Go building blocks:**
- `unicode.IsLetter` / `unicode.IsDigit` / `unicode.ToLower`
- building a `[]rune`, then reusing your `Reverse`-style loop

### TitleCase — split, transform, join

**The concept:** the classic pipeline. Break the input into words, transform each, and glue them back together.

```js
// JS: slugifying a title
const slug = title.trim().toLowerCase().split(/\s+/).join("-");
```

**Go building blocks:**
- `strings.Fields` splits on any whitespace and drops the empty pieces for you
- `strings.ToUpper` / `strings.ToLower` work on whole strings, and `w[:1]` / `w[1:]` slice a string (by bytes, which is fine for ASCII)
- `strings.Join`

### CompressRuns — scanning for groups

**The concept:** walk the string tracking the *current run* (which character, how many). When the character changes, emit the finished run and start a new one. Don't forget the last run after the loop ends.

```js
// JS: collapsing repeated log lines into "line (x3)"
const out = [];
let prev = null, n = 0;
for (const line of lines) {
  if (line === prev) { n++; continue; }
  if (prev !== null) out.push(n > 1 ? `${prev} (x${n})` : prev);
  prev = line; n = 1;
}
```

**Go building blocks:**
- `strings.Builder`
- `strconv.Itoa(n)` (never `string(n)`)
- `range` over runes, remembering the previous one

### CaesarCipher — rune arithmetic with wraparound

**The concept:** treat letters as numbers. Find each letter's position in the alphabet (0–25), shift it, wrap it with modulo, and turn it back into a letter.

```js
// JS: rotating a day-of-week index
const days = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"];
const shiftDay = (i, n) => days[(((i + n) % 7) + 7) % 7];
```

**Go building blocks:**
- `r - 'a'` gives the position, `'a' + rune(pos)` gives the letter back
- the safe-modulo formula above
- separate branches for `'a'..'z'` and `'A'..'Z'`, and leave everything else untouched

### ToMorse — lookup table + error on unknown input

**The concept:** encoding through a lookup table, failing loudly (not silently skipping) when the input has something the table doesn't cover.

```js
// JS: translating status codes, refusing unknowns
const labels = { 200: "OK", 404: "Not Found" };
const label = (code) => {
  if (!(code in labels)) throw new Error(`unknown code ${code}`);
  return labels[code];
};
```

**Go building blocks:**
- the comma-ok lookup on the provided `morse` map
- `unicode.ToUpper`
- `strings.Fields` for words
- collecting parts into a `[]string` and `strings.Join`-ing them
