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

// DPは不得意・・・・
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	p := getInts(N)
	q := getInts(N)
	v := make([]int, N)
	// １回目がv[i]位だった人の２回目の順位
	for i := 0; i < N; i++ {
		v[p[i]-1] = q[i] - 1
	}
	dp := make([][]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, N+1)
	}
	// dp[今までに選んだ人数][選ばなかった人のうち最も小さい順位]
	dp[0][N] = 1
	for i := 0; i < N; i++ {
		tmp := make([][]int, K+1)
		for j := 0; j < K+1; j++ {
			tmp[j] = make([]int, N+1)
		}
		for x := 0; x <= K; x++ {
			for y := 0; y <= N; y++ {
				// xがKより小さく、２回目の順位がyより小さい
				// ※選べる場合
				if x < K && v[i] < y {
					tmp[x+1][y] += dp[x][y]
					tmp[x+1][y] %= mod
				}
				// 選ばない場合は、yと２回目の順位の小さいほう
				tmp[x][min(y, v[i])] += dp[x][y]
				tmp[x][min(y, v[i])] %= mod
			}
		}
		dp = tmp
	}

	ans := 0
	for _, e := range dp[K] {
		ans += e
		ans %= mod
	}
	out(ans)
}
