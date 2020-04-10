package main

import (
	"bufio"
	"fmt"
	"os"
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

var memo [][2]int
var f [][2]int

func rec(p, flg, n int, a []int) int {
	if f[p][flg] == 1 {
		return memo[p][flg]
	}
	if p == n-1 {
		return a[p]
	}
	x0 := a[p] + rec(p+1, 0, n, a)
	a[p+1] = -a[p+1]
	x1 := -a[p] + rec(p+1, 1, n, a)
	a[p+1] = -a[p+1]

	memo[p][flg] = max(x0, x1)
	f[p][flg] = 1
	return max(x0, x1)
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	memo = make([][2]int, N)
	f = make([][2]int, N)

	out(rec(0, 0, N, a))
}
