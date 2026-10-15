# Algorithms 07 · Dynamic Programming

**Builds on:** algorithms 04 (recursion, memoised Fibonacci), language 01 (2D slices).

```sh
go test -timeout 10s ./algorithms/07-dynamic-programming/
```

**Dynamic programming = recursion + never solving the same subproblem twice.** If a brute-force recursive solution keeps asking the same smaller question (`ways(5)` asks `ways(4)` and `ways(3)`, and `ways(4)` asks `ways(3)` *again*), you can store each answer once and build up.

## Complexity

| Kata | Brute force | DP time | DP space |
|------|-------------|---------|----------|
| ClimbStairs | O(2ⁿ) | O(n) | O(n), or O(1) |
| CoinChange | exponential | O(amount × coins) | O(amount) |
| LongestCommonSubsequence | O(2ⁿ⁺ᵐ) | O(n × m) | O(n × m), or O(min) |
| Knapsack | O(2ⁿ) | O(n × capacity) | O(n × capacity), or O(capacity) |

## The recipe

For every problem in this module, answer these four questions *before* writing code:

1. **State:** what does `dp[i]` (or `dp[i][j]`) *mean*, in words? Get this sentence right and the rest follows.
2. **Recurrence:** how is `dp[i]` built from smaller entries?
3. **Base cases:** which entries do you know without computing?
4. **Order:** which direction do you fill the table so that every entry's inputs are already done?

**Top-down** (recursion + memo map, as in algorithms 04) and **bottom-up** (a loop filling a slice) are the same idea. Bottom-up avoids recursion depth and map overhead; top-down only computes the states it needs. Try both on ClimbStairs.

## Go you'll lean on

- 1D table: `dp := make([]int, n+1)`.
- 2D table, which is a slice of slices with each row made separately:
  ```go
  dp := make([][]int, rows)
  for i := range dp { dp[i] = make([]int, cols) }
  ```
- "Infinity" for minimisation: `math.MaxInt`, but adding 1 to it overflows to a huge negative. Use `amount + 1` as a safe "impossible" sentinel instead.
- `min` / `max` built-ins.

---

## The katas

### ClimbStairs — the gentlest DP

**The concept:** to stand on step `n`, your last move came from step `n-1` or step `n-2`. So the ways to reach `n` is the sum of the ways to reach those two. That's a one-dimensional table where each cell looks back a fixed distance.

```js
// JS: the same "look back at earlier answers" shape, a running total table
const prefix = [0];
for (const x of nums) prefix.push(prefix.at(-1) + x);
```

**Go building blocks:** `dp := make([]int, n+1)` and a loop from small `i` to large. The bonus is two rolling variables.

### CoinChange — minimise over choices

**The concept:**
- **State:** `dp[a]` is the fewest coins that make amount `a`.
- To make `a`, the *last* coin was one of the coins `c`, so try each one and take the best `dp[a-c] + 1`.
- Some amounts are impossible, so mark them so they never win a `min`.

```js
// JS: the same "best over every last step" shape, cheapest way to reach each floor with different lift costs
// cost[f] = min over each lift L that ends at f of (cost[f - L.floors] + L.price)
```

**Go building blocks:**
- `make([]int, amount+1)`, filled with an "impossible" sentinel, and `dp[0] = 0`
- an outer loop over amounts, an inner loop over coins
- skip a coin when `c > a`

### LongestCommonSubsequence — the 2D grid

**The concept:**
- **State:** `dp[i][j]` is the LCS length of the first `i` characters of `a` and the first `j` of `b`.
- If those last characters match, extend the diagonal (`dp[i-1][j-1] + 1`).
- If not, take the better of dropping one character from either string.
- An extra row and column of zeros handles the empty-prefix base cases for free.

```
      ""  a  c  e
  ""   0  0  0  0
  a    0  1  1  1
  b    0  1  1  1
  c    0  1  2  2
  d    0  1  2  2
  e    0  1  2  3   ← answer bottom-right
```

**Go building blocks:** a `(len(a)+1) × (len(b)+1)` 2D slice, and compare `a[i-1]` with `b[j-1]` (bytes are fine for ASCII).

### Knapsack — take it or leave it

**The concept:**
- **State:** `dp[i][w]` is the best value using only the first `i` items with capacity `w`.
- For item `i`, either **skip it** (`dp[i-1][w]`) or, if it fits, **take it** (its value + `dp[i-1][w - weight]`). The `i-1` in "take it" is what stops an item being used twice.
- Space bonus: one 1D row is enough if you loop capacity **downwards**. Work out why.

```
items (w,v): (1,1) (3,4) (4,5) (5,7), capacity 7
item (3,4): skip → best of earlier items at w=7 | take → 4 + best of earlier items at w=4
                                                          ↑ "earlier items" is the 0/1 rule
```

**Go building blocks:** a 2D table of `(len(items)+1) × (capacity+1)`, nested loops, and `max`.
