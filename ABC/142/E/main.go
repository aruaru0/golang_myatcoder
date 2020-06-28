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

const inf = 1001001001001001

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	a := make([]int, M)
	c := make([]int, M)
	for i := 0; i < M; i++ {
		a[i] = getInt()
		b := getInt()
		for j := 0; j < b; j++ {
			n := getInt() - 1
			c[i] |= (1 << n)
		}
	}

	dp := make([]int, 1<<N)
	for j := 0; j < 1<<N; j++ {
		dp[j] = inf
	}
	dp[0] = 0

	for i := 1; i <= M; i++ {
		for j := 0; j < 1<<N; j++ {
			dp[j|c[i-1]] = min(dp[j|c[i-1]], dp[j]+a[i-1])
		}
	}

	ans := dp[1<<N-1]
	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
