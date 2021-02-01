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

func solve() {
	N := getI()
	a := getInts(N)
	a = append(a, a[0:2]...)
	c := make([]int, N+2)
	for i := 2; i < N+2; i++ {
		x, y, z := a[i-2], a[i-1], a[i]
		if x != y && x != z && y != z && (nmin(x, y, z) == y || nmax(x, y, z) == y) {
			c[i] = x
		}
	}

	// out(a, c)

	ans := 0
	for k := 0; k < 3; k++ {
		dp := make([]int, N)
		for i := 2; i < N; i++ {
			// out(i, k)
			if i-3 >= 0 {
				dp[i] = nmax(dp[i-1], dp[i-2], dp[i-3]+c[i+k])
			} else {
				dp[i] = nmax(dp[i-1], dp[i-2], c[i+k])
			}
		}
		// out(dp)
		ans = max(ans, dp[N-1])
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		// out("--------------------------------")
		solve()
	}
}
