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

const N = 100005

var (
	n int
	a []int
	b []int
	c []int
	g [][]int
)

func mdy(i, w int) {
	for ; i <= n; i += i & -i {
		c[i] += w
	}
}

func qry(x int) int {
	p := 0
	for i := 16; i >= 0; i-- {
		if p+(1<<i) < n && c[p|(1<<i)] < x {
			p |= 1 << i
			x -= c[p]
		}
	}
	return b[p+1]
}

func dfs(u, p, d int) int {
	mdy(a[u], 1)
	w := 1 << 60
	if d&1 != 0 {
		w = 0
	}
	wv := 0
	if len(g[u]) == 1 && u != 1 {
		w = (qry(d/2) + qry(d/2+1)) / 2
		if d&1 != 0 {
			w = qry((d + 1) / 2)
		}
	}
	for _, v := range g[u] {
		if v != p {
			wv = dfs(v, u, d+1)
			if d&1 > 0 {
				w = max(w, wv)
			} else {
				w = min(w, wv)
			}
		}
	}
	mdy(a[u], -1)
	return w
}

// 写した
// multisetを使った解答はgoでは厳しい
// bitを使った方法でもよいがとりあえず
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	n = getI()
	a = make([]int, n+1)
	b = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		a[i] = getI()
		b[i] = a[i]
	}
	// bをソート
	sort.Ints(b[1 : n+1])

	// a[i]が何番目かをチェックし、aに格納する
	for i := 1; i < n+1; i++ {
		pos := lowerBound(b[1:n+1], a[i])
		a[i] = pos + 1
	}

	c = make([]int, n+1)
	// グラフを作成
	g = make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		u, v := getI(), getI()
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	fmt.Println(dfs(1, 0, 1))
}
