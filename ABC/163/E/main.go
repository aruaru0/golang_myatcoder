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

type data struct {
	val, pos int
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]data, N)
	for i := 0; i < N; i++ {
		a[i].val = getInt()
		a[i].pos = i
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].val > a[j].val
	})
	// out(a)

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, N+1)
	}

	ans := 0
	for l := 0; l <= N; l++ {
		for r := 0; r <= N; r++ {
			if l+r == N {
				// out(dp[l])
				ans = max(ans, dp[l][r])
				break
			}
			v, i := a[l+r].val, a[l+r].pos
			dp[l+1][r] = max(dp[l+1][r], dp[l][r]+asub(l, i)*v)
			dp[l][r+1] = max(dp[l][r+1], dp[l][r]+asub(N-1-r, i)*v)
		}
	}
	out(ans)
}
