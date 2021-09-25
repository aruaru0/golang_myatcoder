package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type pair struct {
	f, s int
}

// P :
type P []pair

func lower_bound(a P, x pair) int {
	// ２分探索
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		// 0, 1のフラグも比較している点に注意
		if a[m].f >= x.f || (a[m].f == x.f && a[m].s >= x.s) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func upper_bound(a P, x pair) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m].f <= x.f || (a[m].f == x.f && a[m].s <= x.s) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

type G struct {
	to []pair
}

// MX :
const MX = 100005

var x [MX]int
var used [MX]bool
var g [MX]G

var ans []int

func dfs(v int) int {
	used[v] = true
	res := x[v]
	for _, e := range g[v].to {
		if used[e.f] {
			continue
		}
		r := dfs(e.f)
		if r == 1 {
			ans = append(ans, e.s)
		}
		res ^= r
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)

	n, m := getInt(), getInt()
	a := make(P, n)
	// 爆弾の座標
	for i := 0; i < n; i++ {
		a[i] = pair{getInt(), getInt()}
	}
	sort.Slice(a, func(i, j int) bool {
		// 距離でソート（距離が同じなら0を先に）
		if a[i].f == a[j].f {
			return a[i].s < a[j].s
		}
		return a[i].f < a[j].f
	})

	// xorを取って変化する部分を見つける
	x[0] = a[0].s
	for i := 0; i < n-1; i++ {
		x[i+1] = a[i].s ^ a[i+1].s
	}
	x[n] = a[n-1].s

	// l,rの位置でグラフを作成
	for i := 0; i < m; i++ {
		l, r := getInt(), getInt()
		l = lower_bound(a, pair{l, 0})
		r = upper_bound(a, pair{r, 1})
		g[l].to = append(g[l].to, pair{r, i + 1})
		g[r].to = append(g[r].to, pair{l, i + 1})
	}

	for i := 0; i <= n; i++ {
		if used[i] {
			continue
		}
		if dfs(i) == 1 {
			out(-1)
			return
		}
	}

	out(len(ans))
	sort.Ints(ans)
	for _, v := range ans {
		fmt.Print(v, " ")
	}
	out("")
}
