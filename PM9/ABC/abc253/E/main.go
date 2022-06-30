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
	N, M, K := getI(), getI(), getI()

	if K == 0 {
		ans := M
		for i := 1; i < N; i++ {
			ans = ans * M % mod
		}
		out(ans)
		return
	}

	dp := make([]int, M+1)
	for i := 0; i <= M; i++ {
		dp[i] = 1
	}
	dpsum := make([]int, M+1)
	for i := 0; i < M; i++ {
		dpsum[i+1] = dp[i+1] + dpsum[i]
	}
	for i := 1; i < N; i++ {
		tmp := make([]int, M+1)
		for j := 1; j <= M; j++ {
			if j-K >= 0 {
				tmp[j] += dpsum[j-K]
				tmp[j] %= mod
			}
			if j+K <= M {
				tmp[j] += dpsum[M] - dpsum[j+K-1]
				if tmp[j] < 0 {
					tmp[j] += mod
				}
				tmp[j] %= mod
			}
		}
		dp = tmp
		for i := 0; i < M; i++ {
			dpsum[i+1] = dp[i+1] + dpsum[i]
			dpsum[i+1] %= mod
		}
	}
	ans := 0
	for i := 0; i <= M; i++ {
		ans += dp[i]
		ans %= mod
	}
	out(ans)
}
