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

const inf = int(1e15)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	a := make([][]int, M)
	for i := 0; i < M; i++ {
		a[i] = getInts(N)
	}

	bit := 1 << N
	dp := make([][]int, M+1)
	for i := 0; i < M+1; i++ {
		dp[i] = make([]int, bit)
		for j := 0; j < bit; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for i := 0; i < M; i++ {
		v := 0
		for j := 0; j < N; j++ {
			v = (v << 1) | a[i][j]
		}
		for b := 0; b < bit; b++ {
			chmin(&dp[i+1][b], dp[i][b])
			dp[i+1][b|v] = nmin(dp[i+1][b|v], dp[i][b|v], dp[i][b]+1)
			// out(i, b, v, b|v, dp[i][b|v], dp[i][b], min(dp[i][b|v], dp[i][b]+1), "-->", dp[i+1][b|v])
		}
		// out(dp[i+1])
	}
	if dp[M][bit-1] == inf {
		out(-1)
		return
	}
	out(dp[M][bit-1])
}
