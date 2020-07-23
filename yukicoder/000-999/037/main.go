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

func main() {
	sc.Split(bufio.ScanWords)
	T, N := getInt(), getInt()
	c := getInts(N)
	v := getInts(N)
	// ｖが１になるまで分解してく
	// ｖは５００なので１つが最大９個程度
	for i := 0; i < N; i++ {
		cost := c[i]
		val := v[i] / 2
		for val > 0 {
			c = append(c, cost)
			v = append(v, val)
			val /= 2
		}
	}
	// 普通のナップサックＤＰ
	n := len(c)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 10100)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= 10000; j++ {
			if j-c[i-1] >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-c[i-1]]+v[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	out(dp[n][T])
}
