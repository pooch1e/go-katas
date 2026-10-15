# 05 · Errors

**Builds on:** 01–03, which already used `(T, error)` and sentinel errors. This module is the deep end.

```sh
go test ./language/05-errors/
```

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `throw new Error("x")` | `return errors.New("x")` | errors are **values you return**, not control flow |
| `try { f() } catch (e) { ... }` | `if err := f(); err != nil { ... }` | you handle each error where it happens |
| `new Error("load config", { cause: err })` | `fmt.Errorf("load config: %w", err)` | **wrapping** keeps the original error inside |
| `e.cause` | `errors.Unwrap(err)` | you'll rarely call this directly |
| `e === NotFoundError` / `e.code === "ENOENT"` | `errors.Is(err, ErrNotFound)` | checks the whole wrap chain |
| `e instanceof ValidationError` | `errors.As(err, &ve)` | finds a *type* in the chain and fills `ve` |
| `new AggregateError([a, b])` | `errors.Join(a, b)` | |
| `class ValidationError extends Error` | any type with an `Error() string` method | that's the whole `error` interface |
| uncaught exception | `panic(...)` | for *programmer bugs* only, never for expected failures |
| `try/finally` | `defer` | runs when the function returns, however it returns |

## New syntax you'll meet

- **The `error` interface** is just `interface { Error() string }`. Any type with that method *is* an error. No `extends` needed. (Module 09 explains interfaces properly.)
- **`%w` vs `%v`.** `fmt.Errorf("...: %w", err)` wraps, so `errors.Is` can see through it. `%v` only copies the text and breaks the chain.
- **Error message style.** Lowercase, no trailing punctuation, colon-separated context: `"email for user 7: find user 7: not found"`. Each layer prepends what *it* was doing.
- **Pointer receivers on error types.** `func (e *ValidationError) Error()` means `*ValidationError` is the error, so you return `&ValidationError{...}` and `errors.As` into a `*ValidationError`.
- **`defer`** schedules a call for when the surrounding function returns. **`recover()`** only does anything inside a deferred function, where it stops a panic and returns the panic value.
- **Named results.** `func f() (err error)` names the return value, so a deferred function can *change* what gets returned. That's how recover turns into a returned error.

---

## The katas

### ParseAge — wrapping a lower-level error

**The concept:** when you call something that can fail, don't throw its error away and don't pass it up bare. **Add context and wrap it**, so the caller gets a readable message *and* can still check the root cause.

```js
// JS: adding context while keeping the cause
async function loadConfig(path) {
  try { return JSON.parse(await readFile(path, "utf8")); }
  catch (err) { throw new Error(`load config ${path}`, { cause: err }); }
}
```

**Go building blocks:**
- `strconv.Atoi(s)` returns `(int, error)`
- `fmt.Errorf` with the `%w` verb to wrap, and `%q` to print a string with quotes around it
- you can wrap a *sentinel* too, adding detail: `fmt.Errorf("%w: ...", ErrSomething, ...)`

### ValidateUser — custom error types and errors.Join

**The concept:** sometimes the caller needs *data* from the error, such as which field failed, not just a message. Make an error type that carries that data. When several things are wrong, report all of them at once instead of making the user fix them one at a time.

```js
// JS: collecting form errors instead of stopping at the first
const problems = [];
if (!form.title) problems.push({ field: "title", reason: "required" });
if (form.tags.length > 5) problems.push({ field: "tags", reason: "max 5" });
if (problems.length) throw new AggregateError(problems.map(p => new FieldError(p)));
```

**Go building blocks:**
- implement `Error()` with `fmt.Sprintf`
- collect into `var errs []error`
- `return errors.Join(errs...)`, which returns `nil` when `errs` is empty
- `strings.TrimSpace` and `strings.Contains`

### UserStore — sentinels through layers

**The concept:** a sentinel error (`ErrNotFound`) is created once at the bottom layer. Every layer above wraps it with what *it* was doing. The top layer can still ask "was this a not-found?" with `errors.Is`, however deep it's buried.

```js
// JS: the same idea using error codes
const err = new Error("read settings", { cause: Object.assign(new Error("no such file"), { code: "ENOENT" }) });
const isMissing = (e) => e?.code === "ENOENT" || (e?.cause && isMissing(e.cause));
```

**Go building blocks:**
- `make` the map in the constructor
- `variadic ...User` parameters work like JS rest params
- comma-ok lookup
- `fmt.Errorf` + `%w` at each layer, with each layer adding only *its own* context

### SafeCall — defer + recover at a boundary

**The concept:** a panic is Go's "this should never happen" crash. You almost never catch them, *except* at a boundary (an HTTP handler, a worker in a pool) where one bad request shouldn't kill the whole server.

```js
// JS: an Express error middleware is the same idea, a last-resort catch at the edge
app.use((err, req, res, next) => res.status(500).send("something broke"));
```

**Go building blocks:**
- `recover()` only does anything when called *directly inside a deferred function*. Anywhere else it returns `nil`.
- `defer` accepts an anonymous function literal: `defer func() { ... }()` (note the call parentheses at the end).
- A **named result** (`func f() (err error)`) is a real variable a deferred function can assign to, changing what the caller receives.
- `%v` prints any value, including whatever was passed to `panic`.
