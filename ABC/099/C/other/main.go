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

const inf = 1001001001

var dp []int

func rec(n int, a []int) int {
	if dp[n] != inf {
		return dp[n]
	}
	ret := inf
	for _, v := range a {
		if n == v {
			return 1
		}
		if n < v {
			continue
		}
		ret = min(ret, rec(n-v, a))
	}
	ret++
	dp[n] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)

	a := make([]int, 0)
	n := 1
	for n < 100000 {
		n *= 6
		a = append(a, n)
	}
	n = 1
	for n < 100000 {
		n *= 9
		a = append(a, n)
	}
	a = append(a, 1)
	sort.Ints(a)

	N := getInt()

	dp = make([]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = inf
	}
	out(rec(N, a))
}
