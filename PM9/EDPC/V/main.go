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

func dfs1(cur, prev int) {
	dp1[cur] = 1
	for _, e := range node[cur] {
		if e == prev {
			continue
		}
		dfs1(e, cur)
		dp1[cur] *= dp1[e] + 1 // 子が黒の場合 + 白の場合（１通り）
		dp1[cur] %= M
	}
}

func dfs2(cur, prev int) {
	dp2[cur] = 1
	for _, e := range node[cur] {
		dp2[cur] *= dp1[e] + 1
		dp2[cur] %= M
	}

	n := len(node[cur])
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = (dp1[node[cur][i]] + 1) % M
		r[i] = l[i]
	}
	for i := 1; i < n; i++ {
		l[i] *= l[i-1]
		l[i] %= M
	}
	for i := n - 2; i >= 0; i-- {
		r[i] *= r[i+1]
		r[i] %= M
	}

	for i := 0; i < n; i++ {
		if node[cur][i] == prev {
			continue
		}
		dp1[cur] = 1
		if i != 0 {
			dp1[cur] *= l[i-1]
			dp1[cur] %= M
		}
		if i+1 < n {
			dp1[cur] *= r[i+1]
			dp1[cur] %= M
		}
		dfs2(node[cur][i], cur)
	}
}

var N, M int
var node [][]int
var dp1 []int
var dp2 []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		x, y := getI()-1, getI()-1
		node[x] = append(node[x], y)
		node[y] = append(node[y], x)
	}

	dp1 = make([]int, N)
	dp2 = make([]int, N)
	dfs1(0, -1)
	dfs2(0, -1)

	for i := 0; i < N; i++ {
		out(dp2[i])
	}
}
