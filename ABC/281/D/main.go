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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K, D := getI(), getI(), getI()
	a := getInts(N)

	// dp[i][k][d] : i個目のうち、k個の和を計算した値をDで割ったときのあまりがdの場合の最大値
	dp := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int, N+1)
		for j := 0; j <= N; j++ {
			dp[i][j] = make([]int, D)
			for k := 0; k < D; k++ {
				dp[i][j][k] = -inf
			}
		}
	}
	dp[0][0][0] = 0
	for i := 0; i < N; i++ {
		for k := 0; k <= i; k++ {
			for d := 0; d < D; d++ {
				if dp[i][k][d] == -inf {
					continue
				}
				sum := dp[i][k][d] + a[i]
				chmax(&dp[i+1][k][d], dp[i][k][d])
				chmax(&dp[i+1][k+1][sum%D], sum)
			}
		}
		// out(i, "-----------")
		// for k := 0; k < N; k++ {
		// 	out(k, dp[i+1][k])
		// }
	}

	ans := dp[N][K][0]
	if ans == -inf {
		out(-1)
	} else {
		out(ans)
	}
}
