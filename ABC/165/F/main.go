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

// 整数用
func lowerBound(a []int, x int) int {
	l := 0
	r := len(a) - 1
	for l <= r {
		m := (l + r) / 2
		if a[m] >= x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(a []int, x int) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m] <= x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

var a []int
var edge [][]int
var dp []int
var ans []int

const inf = 1001001001001

func dfs(x, prev int) {
	pos := lowerBound(dp, a[x])
	val := dp[pos]
	dp[pos] = a[x]
	ans[x] = lowerBound(dp, inf)
	for _, v := range edge[x] {
		if v == prev {
			continue
		}
		dfs(v, x)
	}
	dp[pos] = val

}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a = make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	edge = make([][]int, N)
	for i := 0; i < N-1; i++ {
		u, v := getInt()-1, getInt()-1
		edge[u] = append(edge[u], v)
		edge[v] = append(edge[v], u)
	}

	dp = make([]int, N+1)
	ans = make([]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = inf
	}
	dfs(0, -1)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := 0; i < N; i++ {
		fmt.Fprintln(w, ans[i])
	}
}
