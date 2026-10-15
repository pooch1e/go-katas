# 03 · Structs & Methods

**Builds on:** 01 (errors as return values), 02 (reference vs value types).

```sh
go test ./language/03-structs/
```

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `class Point { constructor(x, y) {...} }` | `type Point struct { X, Y int }` | just data. No constructor, no `this`, no inheritance |
| `new Point(1, 2)` | `Point{X: 1, Y: 2}` | a **struct literal**. Missing fields get zero values |
| `p.add(q)` (method) | `func (p Point) Add(q Point) Point` | methods are declared *outside* the struct, with a **receiver** |
| `this.balance += x` | `func (a *BankAccount) Deposit(...)` | a **pointer receiver** (`*T`) is needed to mutate |
| `#balance` (private field) | `balance int` (lowercase) | lowercase means private to the package |
| `static create(...)` | `func NewBankAccount(...) *BankAccount` | constructors are plain functions named `NewX` by convention |
| `const b = a` (shares the object) | `b := a` (**copies** the struct) | structs are values. Use `&a` / `*T` to share |
| `{...a, x: 5}` | `b := a; b.X = 5` | the copy is free and automatic |

## New syntax you'll meet

- **Value vs pointer receivers.** `func (p Point) ...` gets a *copy* of `p`, so changes inside vanish. `func (a *BankAccount) ...` gets a pointer, so changes stick. Rule of thumb: if any method needs a pointer receiver, make them all pointer receivers for that type.
- **`&` and `*`.** `&x` is "address of x" and gives you a `*T`. You rarely write `*p` to dereference, because Go does it automatically for field access and method calls (`acc.balance` works on a `*BankAccount`).
- **Struct equality.** `==` compares field by field when every field is comparable (numbers, strings, other such structs). Slices and maps aren't comparable, so a struct containing them can't use `==`.
- **Encapsulation is per package, not per type.** Code anywhere in `structkata` can read `a.balance`. The tests live in `structkata_test`, so they can't, and must go through `Balance()`.
- **No getters prefix.** The convention is `acc.Owner()`, not `acc.GetOwner()`.

---

## The katas

### Point — a value object

**The concept:** an immutable value. Operations return a *new* value instead of changing the existing one, so passing it around can never surprise anyone.

```js
// JS: immutable money as a frozen object, where operations return new ones
const add = (a, b) => Object.freeze({ amount: a.amount + b.amount, currency: a.currency });
```

**Go building blocks:**
- value receivers `(p Point)`
- returning a struct literal `Point{X: ..., Y: ...}`
- `math.Sqrt` and `float64(intValue)` conversions. Go never converts number types implicitly, so `int * float64` is a compile error.

### BankAccount — state behind methods

**The concept:** **invariants**. The balance can't go negative and deposits must be positive. The struct's fields are private so the only way to change them is through methods that enforce the rules. `TransferTo` adds **atomicity**: either both halves happen or neither does.

```js
// JS: a counter that refuses to go below zero
class Stock {
  #units = 0;
  take(n) {
    if (n > this.#units) throw new Error("not enough stock");
    this.#units -= n;
  }
}
```

**Go building blocks:**
- `return &BankAccount{owner: owner}` from the constructor
- `(a *BankAccount)` receivers
- `if err := x.Withdraw(n); err != nil { return err }`, which is the most common line in all of Go
- In `TransferTo`, think about *ordering*: which call can fail, and therefore should go first?
