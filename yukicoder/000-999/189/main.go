package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

func condAB(c bool, a, b int) int {
	if c {
		return a
	}
	return b
}

const mod = 1000000009

func main() {
	sc.Split(bufio.ScanWords)
	M, D := getString(), getString()

	var dp1 [210][2][2020]int
	var dp2 [210][2][2020]int

	var f = func(dp *[210][2][2020]int, n int, d string) {
		dp[0][0][0] = 1
		for i := 0; i < n; i++ {
			for j := 0; j < 2; j++ {
				for k := 0; k <= 2010; k++ {
					for a := 0; a <= condAB(j != 0, 9, int(d[i]-'0')); a++ {
						jj := j
						if a < int(d[i]-'0') {
							jj = 1
						}
						// cond1
						dp[i+1][jj][k+a] += dp[i][j][k]
						dp[i+1][jj][k+a] %= mod
					}
				}
			}
		}
	}
	f(&dp1, len(M), M)
	f(&dp2, len(D), D)

	ans := 0
	for i := 1; i <= 2000; i++ {
		x := dp1[len(M)][1][i] + dp1[len(M)][0][i]
		y := dp2[len(D)][1][i] + dp2[len(D)][0][i]
		// out(i, x, y)
		if x != 0 && y != 0 {
			ans += x * y
			ans %= mod
		}
	}
	out(ans)
	// out(dp1[len(M)])
	// out(dp2[len(D)])
}
