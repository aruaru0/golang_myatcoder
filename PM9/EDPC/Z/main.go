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
	N, C := getI(), getI()
	h := getInts(N)

	dp := make([]int, N)
	f := func(j, x int) int {
		return -2*h[j]*x + h[j]*h[j] + dp[j]
	}
	check := func(f1, f2, f3 int) bool {
		a1, b1 := -2*h[f1], h[f1]*h[f1]+dp[f1]
		a2, b2 := -2*h[f2], h[f2]*h[f2]+dp[f2]
		a3, b3 := -2*h[f3], h[f3]*h[f3]+dp[f3]
		return (a2-a1)*(b3-b2) >= (b2-b1)*(a3-a2)
	}

	q := []int{0}
	for i := 1; i < N; i++ {
		for len(q) > 1 && f(q[0], h[i]) >= f(q[1], h[i]) {
			q = q[1:]
		}
		dp[i] = f(q[0], h[i]) + C + h[i]*h[i]
		for len(q) > 1 && check(q[len(q)-2], q[len(q)-1], i) {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	out(dp[N-1])
}
