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

const inf = 99999

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, X, Y := getI(), getI(), getI()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getI(), getI()
	}

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, X+1)
		for j := 0; j < X+1; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0
	for i := 0; i < N; i++ {
		tmp := make([][]int, N+1)
		for j := 0; j < N+1; j++ {
			tmp[j] = make([]int, X+1)
			copy(tmp[j], dp[j])
		}

		for j := 0; j <= i; j++ {
			for k := 0; k <= X-a[i]; k++ {
				tmp[j+1][k+a[i]] = min(tmp[j+1][k+a[i]], dp[j][k]+b[i])
			}
		}
		// out(tmp)
		dp = tmp
	}

	ans := 0
	for i := 0; i < N+1; i++ {
		// out(i, dp[i])
		for j := 0; j <= X; j++ {
			if dp[i][j] <= Y {
				ans = max(ans, i)
			}
		}
	}
	out(min(N, ans+1))
}
