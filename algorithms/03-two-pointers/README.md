# Algorithms 03 · Two Pointers & Sliding Window

**Builds on:** language 01 (`ReverseInPlace`), 07 (`HasPairWithSum`, the seen-set), algorithms 02 (sorted input as leverage).

```sh
go test ./algorithms/03-two-pointers/
```

## Complexity

| Kata | Brute force | With pointers |
|------|-------------|---------------|
| TwoSumSorted | O(n²) pairs | **O(n)** time, O(1) space |
| RemoveDuplicatesSorted | O(n) with a new slice | **O(n)**, O(1) in place |
| MaxSumSubarrayK | O(n·k) re-summing | **O(n)** |
| LongestUniqueSubstring | O(n²)–O(n³) | **O(n)** time, O(alphabet) space |

The trick behind every O(n) here: **each pointer only moves forward**, so together they take at most 2n steps.

## The two shapes

- **Opposite ends** (`i` from the left, `j` from the right, moving inwards). Use this on *sorted* data, where comparing tells you which pointer to move.
- **Same direction / window** (`left` and `right` both move right). `right` grows the window, and `left` shrinks it when it breaks a rule. Use it for "best contiguous run" problems.

## Go you'll lean on

- `for i, j := 0, len(xs)-1; i < j; { ... }` with the moves inside the body
- the built-in `max(a, b)`
- `[]rune(s)` before windowing over a string, so indexes are characters
- `map[rune]int` for "last index I saw this character at"

---

## The katas

### TwoSumSorted — opposite-end pointers

**The concept:** with the slice sorted, check `nums[i] + nums[j]`. If it's too small, the only way to grow it is to move `i` right; if it's too big, move `j` left. Each move rules out a whole row of pairs at once.

```js
// JS: the neighbouring idea, matching big and small items to fit a weight limit
// sort boxes; heaviest + lightest fit? ship both : ship heaviest alone
```

**Go building blocks:** a `for` loop that runs while the pointers haven't crossed, and a tagless `switch` with three cases (found / too small / too big).

### RemoveDuplicatesSorted — read pointer + write pointer

**The concept:** one pointer **reads** every element, and another marks where the next keeper should be **written**. In sorted input, duplicates are adjacent, so "is this different from the last thing I kept?" is all you need.

```js
// JS: the same read/write shape with a different rule, dropping blank lines in place
let w = 0;
for (const line of lines) if (line.trim() !== "") lines[w++] = line;
lines.length = w;
```

**Go building blocks:**
- a write index that only advances when you keep something
- a `range` or index loop as the read pointer
- your keep-rule compares against the last *kept* value, not the previous *read* one
- return the write index; guard the empty slice

### MaxSumSubarrayK — fixed-size sliding window

**The concept:** don't re-add `k` numbers for every window. Slide it: **add the number entering on the right and subtract the one leaving on the left**. Each step is O(1).

```
[2 1 5 1 3 2], k=3:  (2+1+5)=8 → +1 −2 = 7 → +3 −1 = 9 → +2 −5 = 6   best 9
```

**Go building blocks:**
- sum the first `k` to seed both `window` and `best`
- loop from `k` to `len(nums)`
- `best = max(best, window)`

### LongestUniqueSubstring — variable-size window

**The concept:** grow the window to the right one character at a time. When the new character is already inside the window, **jump `left` past its previous position**. Track the widest window you've seen.

```
"abcabcbb":  [abc] → a repeats, left jumps past the first a → [bca] → ... best 3
```

**Go building blocks:**
- a `map[rune]int` remembering where each character was last seen
- `range` over `[]rune(s)` so the index is the window's right edge
- comma-ok on the map, and be careful: a "last seen" position that's already *behind* `left` doesn't count
- `max` to track the best width
