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

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInts(N)
	}

	dp := make([][][]int, 64)
	for i := 0; i < 64; i++ {
		dp[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]int, N)
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			dp[0][i][j] = a[i][j]
		}
	}

	// ２回で行ける回数、４回で行ける回数、８回で行ける回数・・・2^62で行ける回数まで計算しておく
	for i := 0; i < 63; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				for l := 0; l < N; l++ {
					dp[i+1][k][l] += dp[i][k][j] * dp[i][j][l] % mod
					dp[i+1][k][l] %= mod
				}
			}
		}
	}

	// ansは単位行列
	ans := make([][]int, N)
	for i := 0; i < N; i++ {
		ans[i] = make([]int, N)
		ans[i][i] = 1
	}
	idx := 0
	for K > 0 {
		if K%2 == 1 {
			tmp := make([][]int, N)
			for i := 0; i < N; i++ {
				tmp[i] = make([]int, N)
			}
			// out(idx, dp[idx], ans)
			// K%2==1なら、該当する回数を書ける（行列累乗）
			for j := 0; j < N; j++ {
				for k := 0; k < N; k++ {
					for l := 0; l < N; l++ {
						tmp[k][l] += ans[k][j] * dp[idx][j][l] % mod
						tmp[k][l] %= mod
					}
				}
			}
			ans = tmp
		}
		K /= 2
		idx++
	}
	
	// すべて加算したものが答え
	tot := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			tot += ans[i][j]
			tot %= mod
		}
	}
	out(tot)
}
