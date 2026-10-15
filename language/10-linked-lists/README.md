# 10 · Linked Lists

**Builds on:** 03 (pointers, pointer receivers), 06 (function values), 09 (generic types).

```sh
go test ./language/10-linked-lists/
```

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `{ value, next: null }` | `&Node[T]{Value: v}` | `Next` is a `*Node[T]`, and its zero value is `nil` |
| `node.next = other` | `node.Next = other` | identical. Objects in JS are always pointers; in Go you choose |
| `while (node) { ...; node = node.next }` | `for n := l.head; n != nil; n = n.Next { }` | the canonical list walk |
| `*[Symbol.iterator]() { yield ... }` | `func (l *L) All() iter.Seq[T]` | Go 1.23+ **range-over-func** iterators |
| `for (const v of list)` | `for v := range l.All()` | |
| `[...list]` | `slices.Collect(l.All())` | |

## New syntax you'll meet

- **Recursive types need pointers.** `Next *Node[T]` works. `Next Node[T]` would make the struct infinitely big.
- **Nil pointer dereference.** Reading `n.Next` when `n` is `nil` panics. Every loop condition and edge case (empty list, one node) is about not doing that.
- **Iterators.** An `iter.Seq[T]` is just `func(yield func(T) bool)`. Your function calls `yield(v)` for each value. When `yield` returns `false`, the caller hit `break`, and you **must stop**. It's the same contract as a JS generator's `return()`.

  ```go
  func Count(n int) iter.Seq[int] {
      return func(yield func(int) bool) {
          for i := range n {
              if !yield(i) {
                  return
              }
          }
      }
  }
  ```

---

## The katas

### LinkedList[T] — nodes, head and tail

**The concept:** a list as a chain of nodes, each knowing only the next one.
- Inserting at the front is O(1); walking to index *i* is O(*i*).
- A `tail` pointer makes appending O(1) too, but every operation must keep `head`, `tail` and `size` consistent. Most linked-list bugs are one of those three going stale.

```js
// JS: a browser history as nodes, where going back follows a pointer
let current = { url: "/home", prev: null };
const visit = (url) => (current = { url, prev: current });
const back = () => (current = current.prev ?? current);
```

**Go building blocks:**
- `&Node[T]{Value: v}`
- the `nil` checks for empty-list cases
- `FromSlice` can just loop over `PushBack`

### Reverse — pointer juggling

**The concept:** walk the list once, flipping each `Next` to point backwards. You need three variables (previous, current, next) so you don't lose the rest of the chain mid-flip. Then swap `head` and `tail`.

```js
// JS: the same three-variable dance, rotating a shift schedule
let prev = null, cur = firstShift;
while (cur) { const following = cur.after; cur.after = prev; prev = cur; cur = following; }
```

**Go building blocks:** `var prev *Node[T]` declares a nil node pointer. Go's multi-assignment (`a, b = b, a`) evaluates every right-hand side before assigning, which can shorten the flip, but write it longhand first.

### All — range-over-func iterator

**The concept:** an **iterator** hands out values one at a time without building a slice first. The caller can stop early, and the iterator must respect that.

```js
// JS: a generator that stops when the consumer breaks
function* countdown(n) { while (n > 0) yield n--; }
for (const n of countdown(10)) if (n < 7) break;
```

**Go building blocks:** the `Count` example above, with the canonical list walk in place of the `range`.

### HasCycle — Floyd's tortoise and hare

**The concept:** two pointers moving at different speeds. If there's a loop, the fast one eventually laps the slow one and they meet. If there isn't, the fast one hits `nil`. O(1) memory, unlike a seen-set.

```js
// JS: the same race on a circular running track
let slow = 0, fast = 0;
do { slow = (slow + 1) % lap; fast = (fast + 2) % lap; } while (slow !== fast);
```

**Go building blocks:** two pointer variables, a loop condition that guards *every* `.Next` you're about to follow, and `==` on pointers (identity, not value).
