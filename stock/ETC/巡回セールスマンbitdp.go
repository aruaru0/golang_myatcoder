package main

import (
	"bufio"
	"fmt"
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

const inf = int(1e18)

var x, y, z []int

func cost(i, j int) int {
	return abs(x[i]-x[j]) + abs(y[i]-y[j]) + max(0, z[j]-z[i])
}

func rec(s, v int) int {
	if dp[s][v] >= 0 {
		return dp[s][v]
	}
	if s == (1<<N)-1 && v == 0 {
		dp[s][v] = 0
		return dp[s][v]
	}
	tmp := inf
	for u := 0; u < N; u++ {
		if ((s >> u) & 1) == 0 {
			tmp = min(tmp, rec(s|1<<u, u)+d[v][u])
		}
	}
	dp[s][v] = tmp
	return dp[s][v]
}

var dp [][]int
var d [][]int
var N int

// Atcoder ABC 180E 問題
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	x = make([]int, N)
	y = make([]int, N)
	z = make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i], z[i] = getI(), getI(), getI()
	}

	dp = make([][]int, 1<<N)
	for i := 0; i < 1<<N; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = -1
		}
	}
	d = make([][]int, N)
	for i := 0; i < N; i++ {
		d[i] = make([]int, N)
		for j := 0; j < N; j++ {
			d[i][j] = cost(i, j)
		}
	}

	out(rec(0, 0))
}
