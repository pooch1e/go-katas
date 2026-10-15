# 09 · Interfaces & Generics

**Builds on:** 03 (methods, value vs pointer receivers), 05 (`error` is an interface), 06 (reading generic signatures), 07 (`Set[T comparable]`), 08 (stacks).

```sh
go test ./language/09-interfaces-generics/
```

## Coming from JS

JS has no interfaces at runtime. The closest comparison is **TypeScript's structural typing**, and Go's interfaces work almost exactly like it.

| JS / TS | Go | Why it matters |
|----|----|----------------|
| `interface Shape { area(): number }` | `type Shape interface { Area() float64 }` | a set of method signatures |
| `class Circle implements Shape` | *(nothing)* | satisfaction is **implicit**: have the methods, and you're a Shape |
| duck typing: "if it has `.area()`, call it" | the same, but checked at compile time | |
| `if (s instanceof Circle)` | `c, ok := s.(Circle)` | a **type assertion**, comma-ok again |
| `switch` on `s.kind` | `switch v := s.(type) { case Circle: ... }` | a **type switch** |
| `toString()` | `String() string` | implement `fmt.Stringer` and `fmt.Print` uses it |
| `function f<T>(x: T)` | `func F[T any](x T)` | |
| `<T extends Comparable>` | `[T cmp.Ordered]` | a **constraint**, saying which types are allowed |
| `class Stack<T>` | `type Stack[T any] struct { ... }` | |

## New syntax you'll meet

- **Small interfaces.** Go interfaces are usually one or two methods (`io.Reader`, `fmt.Stringer`, `error`). Define them where they're *used*, not where they're implemented.
- **`var _ Shape = Circle{}`** at package level is a compile-time check that `Circle` satisfies `Shape`. It's common in real code.
- **`any`** is an alias for `interface{}`, the interface with no methods, so every type satisfies it. Use it sparingly.
- **Constraints.** `any` means "anything". `comparable` means "works with `==`". `cmp.Ordered` means "works with `<`, `>`" (ints, floats, strings, and named types built on them).
- **Named types.** `type Celsius float64` creates a brand-new type that *behaves* like `float64` but can have its own methods. `Celsius(21.5)` converts into it.
- **Zero-value-ready types.** `var s Stack[int]` with a `nil` slice inside should just work, because `append` on `nil` is fine. Aim for types that don't need a constructor.

---

## The katas

### Shapes — polymorphism without inheritance

**The concept:** **polymorphism**. Write `TotalArea` once against the *interface*, and it works for every shape, including ones written later by someone else. When you genuinely need the concrete type (to print a description), use a type switch, with a `default` case for types you've never heard of.

```js
// JS: anything with a .render() can go in the list
const widgets = [new Button(), new Chart(), { render: () => "<hr>" }];
const html = widgets.map((w) => w.render()).join("");
```

**Go building blocks:**
- `math.Pi`
- value receivers
- calling a method on an interface value runs the concrete type's version
- `switch v := s.(type)`
- `fmt.Sprintf("%g", x)` prints a float without trailing zeros

### Stack[T] — your own generic container

**The concept:** a **generic data structure**. Write the logic once and get a type-safe stack of ints, strings or structs. The `(T, bool)` return is how Go says "maybe" without `null`.

```js
// TS: a typed cache wrapper, the same idea
class Cache<V> {
  #data = new Map<string, V>();
  get(k: string): V | undefined { return this.#data.get(k); }
}
```

**Go building blocks:**
- pointer receivers `(s *Stack[T])`
- `var zero T` to return "nothing" for any type
- the pop-two-step from module 08

### Largest — constraints

**The concept:** a generic function can only use operations its constraint guarantees. `any` doesn't promise `>`, so `Largest[T any]` won't compile. `cmp.Ordered` does promise it.

```ts
// TS: the constraint says what you're allowed to do with T
function longest<T extends { length: number }>(a: T, b: T): T {
  return a.length >= b.length ? a : b;
}
```

**Go building blocks:** `[T cmp.Ordered]`, the "best so far" loop from 01's `FindMax`, and the built-in `max(a, b)`.

### Celsius — fmt.Stringer

**The concept:** implementing a standard-library interface lets your type plug into code that's never heard of it. `fmt` checks "does this value have `String()`?" and uses it.

```js
// JS: overriding toString changes how template strings print it
class Money { constructor(p) { this.p = p; } toString() { return `£${(this.p / 100).toFixed(2)}`; } }
`${new Money(1999)}`; // "£19.99"
```

**Go building blocks:** `fmt.Sprintf` with a precision verb like `%.2f`. One trap: formatting `c` itself with `%v` or `%s` inside `String()` calls `String()` again, forever. Convert to the underlying type first.
