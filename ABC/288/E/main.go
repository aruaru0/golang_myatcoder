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
	n, m := getI(), getI()
	a := getInts(n)
	c := getInts(n)
	x := getInts(m)

	// mcは、[l][r]の区間の最小値
	mc := make([][]int, n)
	for i := 0; i < n; i++ {
		mc[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		now := inf
		for r := l; r < n; r++ {
			now = min(now, c[r])
			mc[l][r] = now
		}
	}

	// 購入する必要があるかのフラグ
	must := make(map[int]bool, n)
	for i := 0; i < m; i++ {
		must[x[i]-1] = true
	}

	dp := make([]int, 1)
	for i := 0; i < n; i++ {
		p := make([]int, len(dp)+1)
		for j := 0; j < len(dp)+1; j++ {
			p[j] = inf
		}
		dp, p = p, dp
		for j := 0; j < len(p); j++ {
			if !must[i] {
				dp[j] = min(dp[j], p[j])
			}
			dp[j+1] = min(dp[j+1], p[j]+a[i]+mc[i-j][i])
		}
	}

	ans := inf
	for i := 0; i < len(dp); i++ {
		ans = min(ans, dp[i])
	}
	out(ans)
}
