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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func chmax(a *float64, b float64) bool {
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, t, m, k := getI(), getI(), getI(), getI()
	dp := make([][]float64, t+1)
	for i := 0; i <= t; i++ {
		dp[i] = make([]float64, k+1)
	}
	dp[t][k] = 1
	for ti := t - 1; ti >= 0; ti-- {
		for kj := 0; kj <= k; kj++ {
			now := 0.0
			dp2 := make([][]float64, m+1)
			for i := 0; i <= m; i++ {
				dp2[i] = make([]float64, m+1)
			}

			for i := 0; i < m; i++ {
				for j := 0; j <= m; j++ {
					for c := 1; j+c <= m; c++ {
						chmax(&dp2[i+1][j+c], dp2[i][j]+dp[ti+1][min(k, kj+c)])
					}
				}
				for i := 0; i <= m; i++ {
					if i > n {
						break
					}
					x := dp2[i][m]
					x += dp[ti+1][kj] * float64(n-i)
					chmax(&now, x)
				}
			}
			dp[ti][kj] = now / float64(n)
		}
	}

	ans := dp[0][0]
	out(ans)
}
