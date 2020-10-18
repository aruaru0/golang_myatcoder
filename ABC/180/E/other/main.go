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

const inf = int(1e10)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	x := make([]int, N)
	y := make([]int, N)
	z := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i], z[i] = getInt(), getInt(), getInt()
	}

	cost := make([][]int, N)
	for i := 0; i < N; i++ {
		cost[i] = make([]int, N)
		for j := 0; j < N; j++ {
			cost[i][j] = abs(x[j]-x[i]) + abs(y[j]-y[i]) + max(0, z[j]-z[i])
		}
	}
	n := 1 << N
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = inf
		}
	}
	dp[1][0] = 0
	for i := 0; i < n; i++ {
		for from := 0; from < N; from++ {
			if (i>>from)&1 != 1 {
				continue
			}
			for to := 0; to < N; to++ {
				if (i>>to)&1 != 0 {
					continue
				}
				dp[i|1<<to][to] = min(dp[i|1<<to][to], dp[i][from]+cost[from][to])
			}
		}
	}

	ans := inf
	for i := 1; i < N; i++ {
		ans = min(ans, dp[n-1][i]+cost[i][0])
	}
	out(ans)
}
