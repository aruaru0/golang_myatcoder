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

type edge struct {
	to, c, t int
}

const inf = 1001001001001

func main() {
	sc.Split(bufio.ScanWords)
	N, C, V := getInt(), getInt(), getInt()
	s := getInts(V)
	t := getInts(V)
	y := getInts(V)
	m := getInts(V)
	node := make([][]edge, N)
	for i := 0; i < V; i++ {
		from := s[i] - 1
		to := t[i] - 1
		node[from] = append(node[from], edge{to, y[i], m[i]})
		// node[to] = append(node[to], edge{from, y[i], m[i]})
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][C] = 0
	q := make([]int, 0)
	q = append(q, 0)
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		for _, e := range node[n] {
			for j := e.c; j <= C; j++ {
				if dp[e.to][j-e.c] > dp[n][j]+e.t {
					dp[e.to][j-e.c] = dp[n][j] + e.t
					q = append(q, e.to)
				}
			}
		}
	}
	// for i := 0; i < N; i++ {
	// 	out(dp[i])
	// }
	ans := inf
	for i := 0; i <= C; i++ {
		ans = min(ans, dp[N-1][i])
	}
	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
