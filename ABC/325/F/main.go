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
	N := getI()
	d := getInts(N)
	l1, c1, k1 := getI(), getI(), getI()
	l2, c2, k2 := getI(), getI(), getI()

	const inf = int(1e9)
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, k1+1)
		for j := 0; j <= k1; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j <= k1; j++ {
			rem := d[i]
			for nj := j; nj <= k1; nj++ {
				dp[i+1][nj] = min(dp[i+1][nj], dp[i][j]+(rem+l2-1)/l2)
				rem = max(rem-l1, 0)
			}
		}
	}

	const inf2 = int(2e18)
	ans := inf2
	for i := 0; i <= k1; i++ {
		if dp[N][i] > k2 {
			continue
		}
		if dp[N][i] == inf {
			continue
		}
		ans = min(ans, c1*i+c2*dp[N][i])
	}
	if ans == inf2 {
		out(-1)
		return
	}
	out(ans)

}
