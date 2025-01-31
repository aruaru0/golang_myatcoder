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

const maxN = 201

func main() {
	sc.Split(bufio.ScanWords)
	K := getInt()

	x := []int{2, 3, 5, 7, 11, 13}
	y := []int{4, 6, 8, 9, 10, 12}
	dp := make([]int, maxN)
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			pos := x[i] * y[j]
			dp[pos]++
		}
	}
	sum := 0
	for i := 0; i < maxN; i++ {
		sum += dp[i]
	}
	out(float64(dp[K]) / float64(sum))
}
