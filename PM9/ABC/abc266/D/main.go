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

const tmax int = 1e5
const inf int = 1e18

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	X := make([]int, tmax+1)
	A := make([]int, tmax+1)
	for i := 0; i < N; i++ {
		t, x, a := getI(), getI(), getI()
		X[t] = x
		A[t] = a
	}

	dp := make([][]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = make([]int, tmax+1)
		for j := 0; j <= tmax; j++ {
			dp[i][j] = -inf
		}
	}
	dp[0][0] = 0
	for t := 0; t < tmax; t++ {
		for i := 0; i < 5; i++ {
			// 移動しない
			dp[i][t+1] = dp[i][t]
			if i != 0 {
				// １つ左に移動
				chmax(&dp[i][t+1], dp[i-1][t])
			}
			if i != 4 {
				// １つ右に移動
				chmax(&dp[i][t+1], dp[i+1][t])
			}
		}
		// t+1秒目に出現する位置X[t+1]に、スコアを加算
		dp[X[t+1]][t+1] += A[t+1]
	}

	ans := 0
	for i := 0; i < 5; i++ {
		ans = max(ans, dp[i][tmax])
	}
	out(ans)
}
