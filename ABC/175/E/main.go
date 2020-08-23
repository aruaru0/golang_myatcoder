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

var R, C, K int
var v [][]int
var dp [][][4]int

func main() {
	sc.Split(bufio.ScanWords)
	R, C, K = getInt(), getInt(), getInt()
	v = make([][]int, R)
	dp = make([][][4]int, R)
	for i := 0; i < R; i++ {
		v[i] = make([]int, C)
		dp[i] = make([][4]int, C)
	}
	for i := 0; i < K; i++ {
		r, c, p := getInt()-1, getInt()-1, getInt()
		v[r][c] = p
	}
	// for i := 0; i < R; i++ {
	// 	out(v[i])
	// }
	// out("-------")
	dp[0][0][1] = v[0][0]
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			for k := 0; k < 4; k++ {
				if i > 0 {
					dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][k])
					dp[i][j][1] = max(dp[i][j][1], dp[i-1][j][k]+v[i][j])
				}
				if j > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k])
					if k > 0 {
						// dp[i][j][k] = max(dp[i][j][k], dp[i][j][k-1])
						dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k-1]+v[i][j])
					}
				}
			}
			// out(i, j, dp[i][j])
		}
	}

	ma := 0
	for i := 0; i < 4; i++ {
		ma = max(ma, dp[R-1][C-1][i])
	}
	out(ma)
}
