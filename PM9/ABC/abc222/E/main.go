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

func dfs(cur, prev, end int) bool {
	if cur == end {
		return true
	}
	for _, e := range node[cur] {
		if e.to == prev {
			continue
		}
		ret := dfs(e.to, cur, end)
		if ret {
			l[e.idx]++
			return true
		}
	}
	return false
}

type edge struct {
	to, idx int
}

var N, M, K int
var node [][]edge
var l []int

const mod = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, K = getI(), getI(), getI()
	a := make([]int, M)
	for i := 0; i < M; i++ {
		a[i] = getI() - 1
	}
	node = make([][]edge, N)
	for i := 0; i < N-1; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], edge{v, i})
		node[v] = append(node[v], edge{u, i})
	}

	// 辺を利用する回数を数える
	l = make([]int, N-1)
	for i := 0; i < M-1; i++ {
		dfs(a[i], -1, a[i+1])
	}

	// 各辺についてDP（配列よりmapのが軽そう）
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < N-1; i++ {
		p := make(map[int]int)
		for v, e := range m {
			p[v+l[i]] += e
			p[v+l[i]] %= mod
			p[v-l[i]] += e
			p[v-l[i]] %= mod
		}
		m = p
	}

	out(m[K])
}
