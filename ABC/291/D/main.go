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

const mod = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getI(), getI()
	}

	a = append([]int{-1}, a...)
	b = append([]int{-1}, b...)

	// dp[i][j] i番目まで見て、最後のカードがj向きの時の選び方の数
	dp := make([][2]int, N+1)
	dp[0][0] = 1

	for i := 0; i < N; i++ {
		if a[i+1] != b[i] { // １つ前の裏が表と異なるなら
			dp[i+1][0] += dp[i][1]
			dp[i+1][0] %= mod
		}
		if b[i+1] != a[i] { // １つ前の表が裏と異なるなら
			dp[i+1][1] += dp[i][0]
			dp[i+1][1] %= mod
		}
		if a[i+1] != a[i] { // 1つ前の表が、表と異なるなら
			dp[i+1][0] += dp[i][0]
			dp[i+1][0] %= mod
		}
		if b[i+1] != b[i] { // 1つ前の裏が、裏と異なるなら
			dp[i+1][1] += dp[i][1]
			dp[i+1][1] %= mod
		}
	}
	out((dp[N][0] + dp[N][1]) % mod)
}
