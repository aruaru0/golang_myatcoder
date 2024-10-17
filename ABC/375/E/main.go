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

	a := make([]int, N)
	b := make([]int, N)
	x := 0

	for i := 0; i < N; i++ {
		a[i], b[i] = getI()-1, getI()
		x += b[i]
	}

	if x%3 != 0 {
		out(-1)
		return
	}

	x /= 3

	const inf = int(1e18)
	dp := make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		dp[i] = make([]int, x+1)
		for j := 0; j < x+1; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for i := 0; i < N; i++ {
		tmp := make([][]int, x+1)
		for i := 0; i < x+1; i++ {
			tmp[i] = make([]int, x+1)
			for j := 0; j < x+1; j++ {
				tmp[i][j] = inf
			}
		}
		cost1, cost2, cost3 := 1, 1, 1
		switch a[i] {
		case 0:
			cost1 = 0
		case 1:
			cost2 = 0
		case 2:
			cost3 = 0
		}
		for j1 := 0; j1 <= x; j1++ {
			for j2 := 0; j2 <= x; j2++ {
				if j1+b[i] <= x {
					chmin(&tmp[j1+b[i]][j2], dp[j1][j2]+cost1)
				}
				if j2+b[i] <= x {
					chmin(&tmp[j1][j2+b[i]], dp[j1][j2]+cost2)
				}
				chmin(&tmp[j1][j2], dp[j1][j2]+cost3)
			}
		}
		dp = tmp
	}

	ans := dp[x][x]
	if ans == inf {
		ans = -1
	}

	out(ans)
}
