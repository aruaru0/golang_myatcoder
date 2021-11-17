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
	N, X := getI(), getI()
	a := getInts(N)

	// Xをちょうど支払う枚数
	x := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		x[i] = X / a[i]
		X %= a[i]
	}

	// 下のコイン何枚で上のコインになるか
	t := make([]int, N)
	for i := 0; i < N-1; i++ {
		t[i] = a[i+1] / a[i]
	}
	t[N-1] = 1e18

	// out(x)
	// out(t)
	dp := make([]int, 2)
	dp[0] = 1
	for i := 0; i < N; i++ {
		p := make([]int, 2)
		p, dp = dp, p
		for j := 0; j < 2; j++ {
			for nj := 0; nj < 2; nj++ {
				// z : おつり　y ： 支払い
				// y + nj*t = x + z + j
				// z = y(=0) + nj*t - x - j
				// i枚目のコインの支払いを０にする
				z := 0 + t[i]*nj - x[i] - j
				if 0 <= z && z < t[i] {
					dp[nj] += p[j]
				}
				// y + nj*t = x + z + j
				// y = x + z(=0) + j - nj*t
				// i枚目のコインのおつりを０にする
				y := x[i] + 0 + j - t[i]*nj
				if 0 < y && y < t[i] {
					{
						dp[nj] += p[j]
					}
				}
			}
		}
	}
	out(dp[0])
}
