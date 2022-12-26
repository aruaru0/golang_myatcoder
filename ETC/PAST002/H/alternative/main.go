package main

import (
	"bufio"
	"fmt"
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

type pos struct {
	y, x int
}

const inf = 1001001001001001

//          9223372036854775807

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getString()
	}

	a := make([][]pos, 11)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			c := s[i][j]
			if c == 'G' {
				a[10] = append(a[10], pos{i, j})
			} else if c == 'S' {
				a[0] = append(a[0], pos{i, j})
			} else {
				a[int(c-'0')] = append(a[int(c-'0')], pos{i, j})
			}
		}
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			dp[i][j] = inf
		}
	}

	dp[a[0][0].y][a[0][0].x] = 0

	for i := 1; i < 11; i++ {
		for _, to := range a[i] {
			// out("-", to)
			for _, from := range a[i-1] {
				// out("-->", from)
				cost := abs(to.y-from.y) + abs(to.x-from.x)
				dp[to.y][to.x] = min(dp[to.y][to.x], dp[from.y][from.x]+cost)
			}
			// for j := 0; j < N; j++ {
			// 	out(dp[j])
			// }
		}
	}

	ans := dp[a[10][0].y][a[10][0].x]
	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}

}
