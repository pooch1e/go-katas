# 12 · Concurrency

**Builds on:** 05 (errors), 06 (closures), 08 (queues), and Promises / async-await from JS.

```sh
go test ./language/12-concurrency/
go test -race ./language/12-concurrency/   # always use -race for this module
```

This is where Go is genuinely different from JS. JS has **one thread and an event loop**, so two callbacks can never run at the same instant. Go runs goroutines **truly in parallel** across CPU cores, so two goroutines *can* touch the same variable at the same instant. The race detector (`-race`) catches that.

## Coming from JS

| JS | Go | Why it matters |
|----|----|----------------|
| `const p = doWork()` (starts immediately) | `go doWork()` | starts a **goroutine**. It returns nothing, so results come back via channels |
| `await p` | `<-ch` | receiving from a channel blocks until a value arrives |
| `resolve(v)` | `ch <- v` | sending blocks until someone receives (unless the channel is buffered) |
| `Promise.all(ps)` | `sync.WaitGroup` + write into `results[i]` | |
| `Promise.race(ps)` | first value from a shared channel | |
| `Promise.any` with a timeout | `select { case v := <-ch: ... case <-time.After(d): ... }` | `select` waits on several channels at once |
| `AbortController` / `signal` | `context.Context` | `ctx.Done()` is a channel that closes on cancel |
| `AbortSignal.timeout(ms)` | `context.WithTimeout(ctx, d)` | always `defer cancel()` |
| `setTimeout(r, ms)` as a promise | `time.After(d)` | a channel that receives once after `d` |
| *(not needed: single thread)* | `sync.Mutex`, `atomic.Int64` | protect shared state |

## New syntax you'll meet

- **`go f()`** runs `f` concurrently. The `go` statement returns immediately. If `main` (or a test) returns, running goroutines are killed.
- **Channels.**
  - `ch := make(chan int)` is unbuffered: a send waits for a receiver.
  - `make(chan int, n)` has a buffer of `n`: sends only wait once it's full.
  - `close(ch)` signals "no more values", and `for v := range ch` stops on close.
- **`sync.WaitGroup`.** The pattern is `wg.Add(1)` before `go`, and `defer wg.Done()` inside. Then `wg.Wait()`. Go 1.25+ also has `wg.Go(func() { ... })`, which does all three.
- **Loop variables are per-iteration** (since Go 1.22), so `for i := range xs { go func() { use(i) }() }` is safe. Old blog posts warn about this; the warning no longer applies.
- **Writing to *different* indexes of a pre-made slice from different goroutines is safe.** Writing the *same* variable (`total += x`) from several goroutines is a data race.
- **`select`** waits on several channel operations and runs whichever is ready first.

---

## The katas

### SumParallel — fan-out, fan-in

**The concept:** split work into independent pieces, process them in parallel, then combine. Each goroutine computes a *partial* result in its own variable, so nothing is shared until the combine step.

```js
// JS: the same shape with worker threads, one partial sum per worker
const partials = await Promise.all(chunks.map((c) => runInWorker("sum", c)));
const total = partials.reduce((a, b) => a + b, 0);
```

**Go building blocks:**
- chunking (module 06)
- `partials := make([]int, workers)`, where each goroutine writes only its own index
- `sync.WaitGroup`, then sum `partials`

### FetchAll — Promise.all

**The concept:** start everything at once, wait for all of it, and keep results in *input* order (not completion order). Each goroutine knows its index, so it writes to `results[i]`.

```js
// JS
const pages = await Promise.all(urls.map((u) => fetch(u).then((r) => r.text())));
```

**Go building blocks:**
- `results := make([]string, len(urls))` and `errs := make([]error, len(urls))`
- a `WaitGroup`
- after `Wait()`, `errors.Join(errs...)` (nil when they're all nil)
- Going further: `golang.org/x/sync/errgroup` is what real code uses for this.

### FirstResult — Promise.race

**The concept:** every function sends its result to one shared channel, and you take the first value that arrives. The losers must still be able to send without blocking forever, or they leak.

```js
// JS
const fastest = await Promise.race(mirrors.map((m) => fetch(m)));
```

**Go building blocks:** one shared channel, one goroutine per function sending into it, and a single receive. The test's hint asks the key question: what happens to a sender when nobody receives?

### WithTimeout — context deadlines

**The concept:** don't let a slow operation hang forever. Give it a **context** with a deadline. Well-behaved code watches `ctx.Done()` and gives up when it closes. The caller gets `context.DeadlineExceeded`.

```js
// JS
const res = await fetch(url, { signal: AbortSignal.timeout(2000) });
```

**Go building blocks:**
- `context.Background()` is the root context to derive from
- `context.WithTimeout` returns a new context *and* a `cancel` func. Always arrange for `cancel` to be called, or the timer leaks.
- The test's `fn` already respects the context, so your job is only the plumbing.

### WorkerPool — bounded concurrency

**The concept:** don't start 10,000 goroutines for 10,000 jobs. Start `workers` goroutines that each pull jobs from a shared channel until it's closed. This caps how much runs at once (database connections, API rate limits).

```js
// JS: p-limit style, allowing at most 3 uploads at a time
const limit = pLimit(3);
await Promise.all(files.map((f) => limit(() => upload(f))));
```

**Go building blocks:**
- a channel carrying job *indexes*, so each result knows which slot to write
- `for x := range ch` in each worker keeps receiving until the channel is closed
- the order of the last three steps matters: start workers, send everything, `close`, wait

### Retry — loops, delays and cancellation

**The concept:** retry a flaky operation a fixed number of times with a pause between tries. The pause must be **cancellable**: if the caller gives up, you stop immediately instead of sleeping on.

```js
// JS
for (let i = 0; i < attempts; i++) {
  try { return await op(); }
  catch (e) { last = e; await sleep(delay, { signal }); }
}
throw last;
```

**Go building blocks:**
- remember the most recent error so you can wrap it at the end
- a `select` with two cases, the context's `Done()` channel and `time.After(delay)`, is a sleep that can be interrupted
- `ctx.Err()` tells you why a context ended
