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

const inf = 1e18

func calc(pos int) float64 {
	n := 1 << N
	dp := make([][]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}

	// pos -> i へのルートを初期化しておく
	for i := 0; i < N; i++ {
		dist := math.Hypot(x[i]-x[pos], y[i]-y[pos])
		dp[i][1<<i] = dist
	}

	// ビットDP
	for bit := 0; bit < n; bit++ {
		for from := 0; from < N; from++ {
			if bit>>from%2 == 0 {
				continue
			}
			for to := 0; to < N; to++ {
				if bit>>to%2 == 1 {
					continue
				}
				dist := math.Hypot(x[from]-x[to], y[from]-y[to])
				dp[to][bit|1<<to] = math.Min(dp[to][bit|1<<to], dp[from][bit]+dist)
			}
		}
	}

	return dp[pos][n-1]
}

var N int
var x, y []float64

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	x = make([]float64, N)
	y = make([]float64, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getF(), getF()
	}

	ret := calc(0)
	out(ret)
}
