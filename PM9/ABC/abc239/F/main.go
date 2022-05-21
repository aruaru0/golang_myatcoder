package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

const maxn = 200010

var fa = make([]int, maxn)

func Find(p int) int {
	if p == fa[p] {
		return p
	}
	fa[p] = Find(fa[p])
	return fa[p]
}

func Merge(x, y int) bool {
	x = Find(x)
	y = Find(y)
	if x == y {
		return false
	}
	fa[x] = y
	return true
}

// 解き方は分かったが、実装したらバグったので
// 丸写しに変更（今日は時間がないので）
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()

	d := make([]int, n+1)
	u := 0
	for i := 1; i <= n; i++ {
		d[i] = getI()
		u += d[i]
		fa[i] = i
	}

	// 数が合わない場合は-1
	if u != ((n - 1) << 1) {
		out(-1)
		return
	}

	for i := 1; i <= m; i++ {
		u, v := getI(), getI()
		if Merge(u, v) {
			d[u]--
			d[v]--
			// 負になるのは許されない
			if d[u] < 0 || d[v] < 0 {
				out(-1)
				return
			}
		} else { // 既につながっていたら-1
			out(-1)
			return
		}
	}

	s := make([]int, n+1)
	dg := make([]int, n+1)
	nxt := make([]int, n+1)
	h := make([]int, n+1)
	j := 0
	for i := 1; i <= n; i++ {
		u = Find(i)
		if s[u] == 0 {
			j++
			s[u] = j
		}
		if d[i] != 0 {
			dg[s[u]] += d[i]
			nxt[i] = h[s[u]]
			h[s[u]] = i
		}
	}

	type node struct{ ind, idx int }
	N := make([]node, j+1)
	t := 0
	for i := 1; i <= j; i++ {
		N[i] = node{dg[i], i}
	}
	sort.Slice(N, func(i, j int) bool {
		return N[i].ind < N[j].ind
	})
	N = append(N, node{0, 0})

	stk := make([]int, j+1)
	var i int
	for i = 1; i <= j; i++ {
		if N[i].ind == 1 {
			t++
			stk[t] = N[i].idx
		} else {
			break
		}
	}

	t1 := 0
	a1 := make([]int, n+1)
	a2 := make([]int, n+1)
	for t > 0 {
		u = stk[t]
		t--
		if N[i].idx == 0 {
			continue
		}
		t1++
		a1[t1] = h[u]
		a2[t1] = h[N[i].idx]
		d[a2[t1]]--
		if d[a2[t1]] == 0 {
			h[N[i].idx] = nxt[a2[t1]]
		}
		N[i].ind--
		if N[i].ind == 1 {
			t++
			stk[t] = N[i].idx
			i++
		}
	}

	// 最終処理
	a1[t1+1] = h[stk[1]]
	if stk[2] != 0 {
		t1++
		a2[t1] = h[stk[2]]
	}
	if t1+m != n-1 {
		out(-1)
	} else {
		for i := 1; i <= t1; i++ {
			out(a1[i], a2[i])
		}
	}
}
