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

var N, T int
var A []int

var dp [51][100100]bool

func rec(cur, v int, S string) {
	if v > T {
		return
	}
	if cur == N {
		if v == T {
			out(S)
			os.Exit(0)
		}
		return
	}
	if dp[cur][v] {
		return
	}
	dp[cur][v] = true
	rec(cur+1, v+A[cur], S+"+")
	rec(cur+1, v*A[cur], S+"*")
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	T = getInt()
	A = getInts(N)
	rec(1, A[0], "")
}
