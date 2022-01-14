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

var node [][]int
var route []int
var dist []int

func dfs(cur, prev, cnt int) {
	dist[cur] = cnt
	route = append(route, cur)
	for _, nxt := range node[cur] {
		if nxt == prev {
			continue
		}
		dfs(nxt, cur, cnt+1)
	}
	route = append(route, cur)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	p := getInts(N - 1)
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		f := p[i] - 1
		t := i + 1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
	}
	route = make([]int, 0)
	dist = make([]int, N)
	// 行き掛けと帰りがけにrouteにノード番号をメモしつつ、深さを記録する
	dfs(0, -1, 0)
	// out(route, dist)

	// route中の行き、帰りの位置と、深さごとのroute中の位置を記録する
	io := make([][]int, N)
	depth := make(map[int][]int)
	used := make([]bool, N)
	for i, e := range route {
		io[e] = append(io[e], i)
		if !used[e] {
			depth[dist[e]] = append(depth[dist[e]], i)
			used[e] = true
		}
	}
	// out(io)
	// out(depth)
	Q := getI()
	for i := 0; i < Q; i++ {
		u, d := getI()-1, getI()
		// uがrouteで登場する位置を調べる
		l, r := io[u][0], io[u][1]
		// 深さがdのものが、区間[l,r)にあるかどうか調べ、範囲を出力
		x := lowerBound(depth[d], l)
		y := lowerBound(depth[d], r)
		out(y - x)
	}
}
