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
	N := getI()
	s := getS()

	// dp[i][j]
	//  i番目まで決めたとき、i番目の数より大きいものが
	//  j個残っているような場合の数
	dp := make([][]int, N)
	dpsum := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N+1)
		dpsum[i] = make([]int, N+2)
	}
	// 最初はN以外は1
	for i := 0; i < N; i++ {
		dp[0][i] = 1
		dpsum[0][i+1] = dpsum[0][i] + dp[0][i]
	}

	for i := 1; i < N; i++ {
		if s[i-1] == '<' {
			// もし大きいなら、j+1から残り最大個数(N-i)までを加算
			for j := 0; j < N-i; j++ {
				// for k := j + 1; k <= N-i; k++ {
				// 	dp[i][j] = (dp[i][j] + dp[i-1][k]) % mod
				// }

				// 上の式を塁積和で表現（添え字に中）
				dp[i][j] = (dpsum[i-1][N-i+1] - dpsum[i-1][j+1]) % mod
				dp[i][j] = (dp[i][j] + mod) % mod
				dpsum[i][j+1] = (dpsum[i][j] + dp[i][j]) % mod
			}
		} else {
			// もし小さいなら、0から残りjまでを加算
			for j := 0; j < N-i; j++ {
				// for k := 0; k <= j; k++ {
				// 	dp[i][j] = (dp[i][j] + dp[i-1][k]) % mod
				// }

				// 上の式を塁積和で表現（添え字に中）
				dp[i][j] = (dpsum[i-1][j+1] - dpsum[i-1][0]) % mod
				dp[i][j] = (dp[i][j] + mod) % mod
				dpsum[i][j+1] = (dpsum[i][j] + dp[i][j]) % mod
			}
		}
	}
	out(dp[N-1][0])
}
