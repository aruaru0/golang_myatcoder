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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()

	c := make([]int, n)
	p := make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i], p[i] = getI(), getI()
		a[i] = getInts(p[i])
	}

	// dp[r] = (dp[r] ++ dp[xx] + dp[xxx] + ... ) / p[i] + c[i]
	// (1-1/p[i])dp[r] = (dp[xx] + ... + )/p[i] + c[i]
	// つまり、
	// not_zero = c[i] + (dp[..] + dp[...] + ... + dp[....])/p[i]
	// dp[r] = dp[r] * 1/p[i] * zero + not_zero
	// (1-1/p[i]*zero)dp[r] = not_zero
	// dp[r] = not_zero / (1-1/p[i]*zero) -> 1/p[i]*zero = b
	// dp[r] = not_zero / (1-b)

	dp := make([]float64, m+1)
	for r := 1; r <= m; r++ {
		dp[r] = 1e18
		for i := 0; i < n; i++ {
			now, b := 0.0, 0.0
			for _, d := range a[i] {
				if d != 0 {
					now += dp[max(0, r-d)]
				} else {
					b += 1.0
				}
			}
			now /= float64(p[i])
			b /= float64(p[i])
			now += float64(c[i])
			now /= 1 - b
			dp[r] = math.Min(dp[r], now)
		}
	}

	out(dp[m])
}
