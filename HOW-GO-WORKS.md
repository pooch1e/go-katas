# How Go actually runs (for a JS dev on Linux)

## Compile first, then run, but fast

**JS:** hand a `.js` file to Node and it runs straight away. Mistakes show up when that line executes.

**Go:** the compiler reads every `.go` file in a package, checks all the types, and produces a native binary. If *anything* is wrong (a type mismatch, an unused variable, an unused import), it refuses to build and nothing runs.

```
  your .go files ──(compile)──▶ a native binary ──(run)──▶ output
                    go build                       ./program
                    └────────── go run / go test do both ──────────┘
```

The closest JS comparison is TypeScript with `strict` and `noUnusedLocals`, except it's mandatory and very fast. Go was designed for quick compiles, so `go test` usually feels as instant as vitest.

Unlike Node, the output is a **standalone binary** with no runtime to install. Copy it to another Linux machine and it runs.

## What's installed

Installed globally with [mise](https://mise.jdx.dev), the same way your Node is:

| Thing | What it is | JS equivalent |
|-------|-----------|---------------|
| `go` | compiler, test runner, formatter, package manager in one | `node` + `npm` + `tsc` + `prettier` |
| `gopls` | the language server (autocomplete, go-to-definition, auto-imports) | the TypeScript language server |
| `gotestsum` | nicer test output plus `--watch` | `vitest` watch mode UI |
| `dlv` (Delve) | the debugger | Node inspector / Chrome DevTools |

Check with `go version`, `go env GOPATH`, and `mise ls`.

The repo's `mise.toml` pins the same Go version for any other machine: clone, `mise install`, done.

## The project files, mapped to JS

```
go-katas/
├── go.mod          ← like package.json: module name, Go version, dependencies
├── go.sum          ← like package-lock.json: checksums. Commit it, never edit it
├── mise.toml       ← like .nvmrc / .tool-versions
└── language/01-slices/
    ├── slices.go           ← package slicekata
    └── sum_all_test.go     ← package slicekata_test
```

- **A package is a folder.** Every `.go` file in a folder declares the same `package` name, and they all see each other's code with no imports. File names don't matter to the compiler.
- **Imports are by path, not file.** The tests do `import slicekata "github.com/pooch1e/go-katas/language/01-slices"`. The path is the `go.mod` module name plus the folder. The name before it (`slicekata`) is what you type in code.
- **Why `slicekata`, not `slices`?** A package named `slices` would clash with the standard library's `slices` package, which the tests also use. Every module's package name ends in `kata` for that reason.
- **Capital letter = exported.** `SumAll` is visible to other packages; `sumAll` isn't. That's Go's whole `export` system, and it applies to functions, types, struct fields and methods.
- **`_test.go` files** are only compiled by `go test`. A `package slicekata_test` (note the suffix) is a separate package in the same folder. The tests can only use exported names, which is why they act like real callers.
- **No `node_modules`.** Dependencies download to a shared cache (`$(go env GOMODCACHE)`). `go get github.com/x/y` is `npm install`, and `go mod tidy` cleans up `go.mod` to match your imports.

## Everyday commands

Run these from the repo root:

| Command | What it does | JS equivalent |
|---------|--------------|---------------|
| `go test ./...` | build + test every package | `npx vitest run` |
| `go test ./language/01-slices/` | one package | `vitest run path/` |
| `go test -run TestFindMax ./language/01-slices/` | matching tests only | `it.only` / `vitest -t` |
| `go test -v ...` | show every test, including passes and skips | `--reporter=verbose` |
| `go test -race ...` | detect data races (module 12) | — |
| `go test -bench=. ...` | run benchmarks | `vitest bench` |
| `gotestsum --watch` | rerun tests on save | `vitest` |
| `go vet ./...` | catch suspicious code that compiles | `eslint` |
| `gofmt -w .` / `go fmt ./...` | format everything (there are no options; everyone's Go looks the same) | `prettier --write` |
| `go build ./...` | compile only | `tsc --noEmit` |
| `go run .` | build + run a `package main` program | `node index.js` |
| `go doc slices.SortFunc` | read the docs for anything | MDN |

## Where `console.log` went

- **In code:** `fmt.Println(x)`, or `fmt.Printf("%+v\n", x)` for structs with field names, or `%#v` for Go syntax. Remove `fmt` from imports when you're done, because an unused import is a compile error.
- **In tests:** `t.Log(x)` or `t.Logf("got %v", x)`. These only print for failing tests, or with `-v`, which keeps passing runs quiet.

## Reading a compile error

```
language/01-slices/slices.go:23:2: declared and not used: total
language/01-slices/slices.go:27:9: cannot use best (variable of type float64) as int value in return statement
```

That's `file:line:column: message`, the same as tsc. Two Go-specific ones you'll hit early:

- **`declared and not used`.** Go refuses to compile unused local variables. Use it, or name it `_`.
- **`"fmt" imported and not used`.** Same rule for imports. gopls can add and remove imports for you on save; if Zed isn't doing it, run the "organize imports" code action.

## Zed setup

Zed has Go support built in. With `gopls` on your PATH (mise put it there), it just works:

- **Format on save** runs `gofmt` automatically, so you never think about formatting.
- **Auto-imports.** gopls offers to add the import when you use a package like `slices`, and the "organize imports" code action tidies unused ones.
- **Inline errors** appear as you type, just like TypeScript.
- **Run a test:** a ▶ appears next to each `func TestXxx` in a `_test.go` file. Click it to run just that test, or the whole package.
- **Debug a test:** Zed's debugger uses Delve for Go. Set a breakpoint by clicking the gutter, then start the debugger on the test from the ▶ menu or the command palette (`debugger: start`). You get step over/into, variables and a call stack, just like Chrome DevTools.

### Debugging from the terminal (any editor)

```sh
dlv test ./language/01-slices/ -- -test.run TestFindMax
(dlv) break slicekata.FindMax
(dlv) continue
(dlv) print nums
(dlv) next
```

### VS Code

Install the official **Go** extension (`golang.go`). It uses the same `gopls` and `dlv`. "run test | debug test" links appear above each test function.
