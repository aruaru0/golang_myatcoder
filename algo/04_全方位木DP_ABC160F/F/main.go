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

type mint int

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

func (m *mint) plus(a mint) mint {
	x := *m + a
	x %= mod
	return x
}

func (m *mint) minus(a mint) mint {
	x := *m - a
	x %= mod
	if x < 0 {
		x += mod
	}
	return x
}

func (m *mint) mul(a mint) mint {
	x := *m * a
	x %= mod
	return x
}

func (m *mint) pow(t int) mint {
	x := mpow(int(*m), t)
	return mint(x)
}

func (m *mint) inv() mint {
	return m.pow(mod - 2)
}

func (m *mint) div(a mint) mint {
	x := *m * a.inv()
	x %= mod
	return x
}

func comb(n, r int) mint {
	return mint(nCr(n, r))
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

//----------------------------------------------------------
// 全方位木DP
//    ABC 160の解答です。解説をgolangに移植しています
//----------------------------------------------------------
type dp struct {
	dp mint
	t  int
}

func newDp() dp {
	return dp{1, 0}
}

func (d *dp) plus(a dp) {
	d.dp = d.dp.mul(a.dp)
	d.dp = d.dp.mul(comb(d.t+a.t, d.t))
	d.t += a.t
}

func (d *dp) minus(a dp) dp {
	res := *d
	res.t -= a.t
	out("minus ", d.dp, a.dp, res.t+a.t, res.t, comb(res.t+a.t, res.t))
	res.dp = res.dp.div(comb(res.t+a.t, res.t))
	out("minus ", res)
	res.dp = res.dp.div(a.dp)
	out("minus ", res)
	return res
}

func (d *dp) addRoot() dp {
	res := *d
	res.t++
	return res
}

type node struct {
	to []int
}

var used []int

var dsum []dp
var d [][]dp

func dfs(pos int, n []node) dp {
	used[pos] = 1
	d[pos] = make([]dp, len(n[pos].to))
	for i := 0; i < len(n[pos].to); i++ {
		d[pos][i] = dp{1, 0}
	}
	for i, v := range n[pos].to {
		if used[v] == 1 {
			continue
		}
		d[pos][i] = dfs(v, n)
		dsum[pos].plus(d[pos][i])
	}
	return dsum[pos].addRoot()
}

func bfs(v, p int, dpP dp, n []node) {
	if p != -1 {
		out("add=", dpP, dsum[v])
		dsum[v].plus(dpP)
	}
	for i, u := range n[v].to {
		if u == p {
			d[v][i] = dpP
			continue
		}
		d := dsum[v].minus(d[v][i])
		bfs(u, v, d.addRoot(), n)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	n := make([]node, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		n[f].to = append(n[f].to, t)
		n[t].to = append(n[t].to, f)
	}

	dsum = make([]dp, N)
	for i := 0; i < N; i++ {
		dsum[i] = dp{1, 0}
	}
	d = make([][]dp, N)

	used = make([]int, N)
	dfs(0, n)
	bfs(0, -1, dp{1, 0}, n)
	out(d)
	out(dsum)

	for i := 0; i < N; i++ {
		out(dsum[i].dp)
	}
}
