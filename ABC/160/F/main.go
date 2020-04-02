package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// mod付き整数演算
const mod = 1000000007

func mpow(x, n int) int {
	ans := 1
	for n != 0 {
		if n&1 == 1 {
			ans = ans * x % mod
		}
		x = x * x % mod
		n = n >> 1
	}
	return ans
}

func mplus(a, b int) int {
	x := a + b
	x %= mod
	return x
}

func minus(a, b int) int {
	x := a - b
	x %= mod
	if x < 0 {
		x += mod
	}
	return x
}

func mmul(a, b int) int {
	x := a * b
	x %= mod
	return x
}

func minv(a int) int {
	return mpow(a, mod-2)
}

func mdiv(a, b int) int {
	x := a * minv(b)
	x %= mod
	return x
}

func comb(n, r int) int {
	return nCr(n, r)
}

// mod付きコンビネーション
var fracMemo = []int{1, 1}

func mfrac(n int) int {
	if len(fracMemo) > n {
		return fracMemo[n]
	}
	if len(fracMemo) == 0 {
		fracMemo = append(fracMemo, 1)
	}
	for len(fracMemo) <= n {
		size := len(fracMemo)
		fracMemo = append(fracMemo, fracMemo[size-1]*size%mod)
	}
	return fracMemo[n]
}

var ifracMemo = []int{1, 1}

func mifrac(n int) int {
	if len(ifracMemo) > n {
		return ifracMemo[n]
	}
	if len(ifracMemo) == 0 {
		fracMemo = append(ifracMemo, 1)
	}
	for len(ifracMemo) <= n {
		size := len(ifracMemo)
		ifracMemo = append(ifracMemo, ifracMemo[size-1]*mpow(size, mod-2)%mod)
	}
	return ifracMemo[n]
}

func nCr(n, r int) int {
	if n == r {
		return 1
	}
	if n < r || r < 0 {
		return 0
	}
	ret := 1
	ret *= mfrac(n)
	ret %= mod
	ret *= mifrac(r)
	ret %= mod
	ret *= mifrac(n - r)
	ret %= mod
	return (ret)
}

//------
type node []int

type sDP struct {
	d int
	t int
}

var dp [][]sDP
var dsum []sDP

func calc(x, n, r int) int {
	res := comb(n, r)
	res = mmul(res, x)
	return res
}

func icalc(y, x, n, r int) int {
	res := mdiv(y, comb(n, r))
	res = mdiv(res, x)
	return res
}

func dfs(v, p int, n []node) sDP {
	N := len(n[v])
	dp[v] = make([]sDP, N)
	for i := 0; i < N; i++ {
		dp[v][i].d = 1
	}
	for i, u := range n[v] {
		if u == p {
			continue
		}
		dp[v][i] = dfs(u, v, n)
		dp[v][i].t++
		dsum[v].t += dp[v][i].t
		x := calc(dsum[u].d, dsum[v].t, dp[v][i].t)
		dsum[v].d = mmul(dsum[v].d, x)
	}
	return dsum[v]
}

func bfs(v, p int, dpP sDP, n []node) {
	if p != -1 {
		x := calc(dpP.d, dsum[v].t+dpP.t, dsum[v].t)
		dsum[v].d = mmul(dsum[v].d, x)
		dsum[v].t += dpP.t
	}
	for i, u := range n[v] {
		if u == p {
			dp[v][i] = dpP
			continue
		}
		var d sDP
		d.d = icalc(dsum[v].d, dp[v][i].d,
			dsum[v].t, dsum[v].t-dp[v][i].t)
		d.t = dsum[v].t - dp[v][i].t + 1
		bfs(u, v, d, n)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	n := make([]node, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		n[f] = append(n[f], t)
		n[t] = append(n[t], f)
	}
	// out(n, N)

	dp = make([][]sDP, N)
	dsum = make([]sDP, N)
	for i := 0; i < N; i++ {
		dsum[i].d = 1
	}
	dfs(0, -1, n)
	bfs(0, -1, sDP{1, 0}, n)
	// out(dp)
	// out(dsum)
	for i := 0; i < N; i++ {
		out(dsum[i].d)
	}
}
