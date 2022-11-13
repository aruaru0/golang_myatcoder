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
	N, S := getI(), getI()
	a := getInts(N)
	nmax := 0
	for i := 0; i < N; i++ {
		nmax += a[i]
	}

	if nmax < S {
		out(-1)
		return
	}

	dp := make([][]bool, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]bool, nmax+1)
	}

	dp[0][0] = true
	for i := 0; i < N; i++ {
		for j := 0; j <= nmax; j++ {
			dp[i+1][j] = dp[i+1][j] || dp[i][j]
			if j+a[i] <= nmax {
				// out(i, j, a[i], dp[i][j], dp[i+1][j+a[i]] || dp[i][j])
				dp[i+1][j+a[i]] = dp[i+1][j+a[i]] || dp[i][j]
			}
		}
		// out(dp[i+1])
	}

	if dp[N][S] == false {
		out(-1)
		return
	}
	pos := S
	ans := make([]int, 0)
	for i := N - 1; i >= 0; i-- {
		if pos-a[i] >= 0 && dp[i+1][pos-a[i]] {
			pos -= a[i]
			ans = append(ans, i+1)
		}
	}
	out(len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(wr, ans[i], " ")
	}
	out()
	// tot := 0
	// for _, e := range ans {
	// 	tot += a[e-1]
	// }
	// out(tot)
}
