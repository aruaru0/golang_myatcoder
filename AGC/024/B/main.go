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

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getInt()
	}

	const inf = 1001001001
	dp := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = 0
	}
	for i := 0; i < N; i++ {
		dp[p[i]] += dp[p[i]-1] + 1
	}

	m := 0
	for i := 0; i <= N; i++ {
		m = max(m, dp[i])
	}

	// out(p)
	// out(dp)
	// out(N, m)
	out(N - m)
}
