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

var N int
var node [][]int
var a, b []int
var depth []int
var dp []int

func dfs(v, p, d int) {
	depth[v] = d
	for _, e := range node[v] {
		if p == e {
			continue
		}
		dfs(e, v, d+1)
	}
}

func dfs2(v, p, c int) {
	dp[v] += c
	for _, e := range node[v] {
		if e == p {
			continue
		}
		dfs2(e, v, dp[v])
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	node = make([][]int, N)
	a = make([]int, N)
	b = make([]int, N)
	depth = make([]int, N)
	for i := 0; i < N-1; i++ {
		f, t := getI()-1, getI()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
		a[i], b[i] = f, t
	}
	dfs(0, -1, 0)
	dp = make([]int, N)

	Q := getI()
	for i := 0; i < Q; i++ {
		T, e, x := getI(), getI()-1, getI()
		f, t := a[e], b[e]
		if T == 2 {
			f, t = t, f
		}
		// out(t, e, x, "from, to", f, t)
		if depth[f] > depth[t] {
			dp[f] += x
		} else {
			dp[0] += x
			dp[t] -= x
		}
	}
	// out(dp)
	dfs2(0, -1, 0)

	for _, e := range dp {
		out(e)
	}
}
