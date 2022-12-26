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
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	a := getInts(N)

	dp := make([][2]int, N+1)
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < N; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		if a[i] > a[i-1] {
			dp[i][1] = dp[i-1][0] + 1
		} else if a[i] < a[i-1] {
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]+1)
		} else if a[i] == a[i-1] {
			dp[i][1] = dp[i-1][1] + 1
		}
		// out(dp[i])
	}
	out(max(dp[N-1][0], dp[N-1][1]))
}
