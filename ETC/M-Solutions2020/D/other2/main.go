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
	N := getInt()
	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		a[i] = getInt()
	}
	dp := make([]int, N+1)

	dp[0] = 1000
	for k := 1; k <= N; k++ {
		for i := 0; i < k-1; i++ {
			m := inf
			for j := i + 1; j < k; j++ {
				m = min(m, a[j])
			}
			dp[k] = max(dp[k], dp[i]/m*a[k]+dp[i]%m)
		}
		// out(dp)
	}
	ans := 0
	for i := 0; i <= N; i++ {
		ans = max(ans, dp[i])
	}
	out(ans)
}
