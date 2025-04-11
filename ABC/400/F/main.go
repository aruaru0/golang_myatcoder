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
	c := getInts(N)
	for i := 0; i < N; i++ {
		c[i]--
	}
	x := getInts(N)

	const inf = int(1e18)
	dp := make([][]int, N)
	dp2 := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp2[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = inf
			dp2[i][j] = inf
		}
	}

	for i := 0; i < N; i++ {
		dp2[i][i] = x[c[i]]
	}

	get := func(l, r int) int {
		l = (l + N) % N
		r = (r + N) % N
		return dp[l][r]
	}

	get2 := func(l, r int) int {
		l = (l + N) % N
		r = (r + N) % N
		return dp2[l][r]
	}

	for w := 1; w <= N; w++ {
		for l := 0; l < N; l++ {
			r := (l + w - 1) % N
			if w > 1 && c[l] == c[r] {
				chmin(&dp2[l][r], get2(l, r-1))
			}
			for i := 0; i < w-1; i++ {
				chmin(&dp2[l][r], get2(l, l+i)+get(l+i+1, r))
			}

			chmin(&dp[l][r], dp2[l][r]+w)
			for i := 0; i < w-1; i++ {
				chmin(&dp[l][r], get(l, l+i)+get(l+i+1, r))
			}
		}
	}
	ans := inf
	for l := 0; l < N; l++ {
		chmin(&ans, get(l, l+N-1))
	}

	out(ans)
}
