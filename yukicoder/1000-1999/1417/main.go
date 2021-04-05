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
	s := getS()
	n := len(s)
	dp := make([][3][3][2]int, n+1)
	for i := 0; i < n; i++ {
		dp[i+1][0][0][0] = 1
	}
	dp[0][0][0][1] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 2; l++ {
					lim := 9
					if l == 1 {
						lim = int(s[i] - '0')
					}
					for d := 1; d <= lim; d++ {
						flg := 0
						if l == 1 && d == lim {
							flg = l & 1
						}
						if d%4 == 0 {
							dp[i+1][min(j+2, 2)][k][flg] += dp[i][j][k][l]
							dp[i+1][min(j+2, 2)][k][flg] %= mod
						} else if d%2 == 0 {
							dp[i+1][min(j+1, 2)][k][flg] += dp[i][j][k][l]
							dp[i+1][min(j+1, 2)][k][flg] %= mod
						} else if d%5 == 0 {
							dp[i+1][j][min(k+1, 2)][flg] += dp[i][j][k][l]
							dp[i+1][j][min(k+1, 2)][flg] %= mod

						} else {
							dp[i+1][j][k][flg] += dp[i][j][k][l]
							dp[i+1][j][k][flg] %= mod
						}
					}
				}
			}
		}
		// out(dp[n][2][2][0], dp[n][2][2][1])
	}
	out((dp[n][2][2][0] + dp[n][2][2][1]) % mod)
}
