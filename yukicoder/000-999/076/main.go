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

	var dp [1000100]float64
	const off = 10
	s := []float64{1.0 / 12.0, 2.0 / 12.0, 3.0 / 12.0, 1.0 / 12.0, 3.0 / 12.0, 2.0 / 12.0}
	for i := 1; i <= 1000006; i++ {
		dp[i+off] = 1 + (dp[i+off-1]*s[0] +
			dp[i+off-2]*s[1] +
			dp[i+off-3]*s[2] +
			dp[i+off-4]*s[3] +
			dp[i+off-5]*s[4] +
			dp[i+off-6]*s[5])
	}
	T := getInt()
	for i := 0; i < T; i++ {
		N := getInt()
		out(dp[N+off])
	}
}
