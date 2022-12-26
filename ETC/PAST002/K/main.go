package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	s := getString()
	c := make([]int, N)
	d := make([]int, N)
	for i := 0; i < N; i++ {
		c[i] = getInt()
	}
	for i := 0; i < N; i++ {
		d[i] = getInt()
	}

	const inf = math.MaxInt64 / 4
	dp := make([][]int, N+2)
	for i := 0; i <= N+1; i++ {
		dp[i] = make([]int, N+2)
		for j := 0; j <= N+1; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if s[i] == '(' {
				// dp[i][j]       -> dp[i+1][j+1]
				// dp[i][j] +cost -> dp[i+1][j-1]
				dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j])
				if j != 0 {
					dp[i+1][j-1] = min(dp[i+1][j-1], dp[i][j]+c[i])
				}
			} else {
				// dp[i][j] +cost -> dp[i+1][j+1]
				// dp[i][j]       -> dp[i+1][j-1]
				dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j]+c[i])
				if j != 0 {
					dp[i+1][j-1] = min(dp[i+1][j-1], dp[i][j])
				}
			}
			// dp[i][j]+cost -> dp[i+1][j]
			dp[i+1][j] = min(dp[i+1][j], dp[i][j]+d[i])
		}
	}
	out(dp[N][0])

}
