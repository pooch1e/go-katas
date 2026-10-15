# 08 · Stacks & Queues

**Builds on:** 01 (slices), 04 (strings, `strconv`), 05 (sentinel errors).

```sh
go test ./language/08-stacks-queues/
```

## Coming from JS

| JS | Go (slice as stack/queue) | Why it matters |
|----|----|----------------|
| `stack.push(x)` | `stack = append(stack, x)` | |
| `stack.pop()` | `top := stack[len(stack)-1]; stack = stack[:len(stack)-1]` | two steps: read, then shrink |
| `stack.at(-1)` (peek) | `stack[len(stack)-1]` | no negative indexing. ⚠️ panics if empty |
| `queue.shift()` | `front := q[0]; q = q[1:]` | |
| `queue.push(x)` | `q = append(q, x)` | |
| `switch (op) { case "+": ...; break; }` | `switch op { case "+": ... }` | no `break` needed; Go cases don't fall through |
| `Number("42")` / `parseInt` | `strconv.Atoi("42")` | returns `(int, error)`. `"4x"` is an error, not `4` |

## New syntax you'll meet

- **`switch` without fallthrough.** Each `case` ends by itself. You can list several values (`case "+", "-":`) and use `default:`. A tagless `switch { case x > 0: ... }` replaces `if/else if` chains.
- **Always guard before popping.** `s[len(s)-1]` on an empty slice is an index-out-of-range **panic**. Check `len(s) == 0` first and return an error.
- **Queues via `q[1:]`.** Re-slicing the front is O(1), but the memory behind dropped items isn't freed until the slice is reallocated. That's fine for katas. For heavy use you'd reach for a ring buffer or `container/list`.
- **Bracket pairs as a map.** `map[rune]rune{')': '(', ']': '[', '}': '{'}` maps each closer to its opener.

---

## The katas

### IsBalanced — stack of "what must close next"

**The concept:** a **stack** is last-in, first-out, exactly the shape of nesting. Push every opener. On a closer, the top of the stack must be its partner. At the end, the stack must be empty.

```js
// JS: tracking nested HTML tags with a stack
const open = [];
for (const tag of tags) {
  if (!tag.startsWith("/")) open.push(tag);
  else if (open.pop() !== tag.slice(1)) return false;
}
return open.length === 0;
```

**Go building blocks:**
- `var stack []rune`
- `range` over the string
- the closer→opener map
- comma-ok to ask "is this a closer?"

### EvaluateRPN — stack of operands

**The concept:** in postfix notation, operands arrive first and operators consume the two most recent. A stack holds the numbers waiting to be used. Pop **right then left**, because order matters for `-` and `/`.

```js
// JS: an undo stack, where the most recent action is undone first
const history = [];
const doAction = (a) => { a.apply(); history.push(a); };
const undo = () => history.pop()?.revert();
```

**Go building blocks:**
- `strconv.Atoi` to tell numbers from operators
- `switch tok { case "+": ... }`
- a `pop` helper (a closure is fine)
- careful error returns at every step that can fail

### HotPotato — queue rotation

**The concept:** a **queue** is first-in, first-out. "Passing" means taking from the front and putting at the back. After `passes` rotations, drop whoever is at the front.

```js
// JS: round-robin task scheduling
while (tasks.length) {
  const t = tasks.shift();
  t.runSlice();
  if (!t.done) tasks.push(t);
}
```

**Go building blocks:**
- `slices.Clone(names)` first, so you don't mutate the caller's data
- "take from the front" and "add to the back" from the table above, combined in whatever order the rules need
