package main

import (
	"bufio"
	"fmt"
	"math"
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

// 最大流を求めるプログラム　Ford-Fulkerson法
type edge struct {
	to, cap, rev int
}

type node struct {
	to []edge
}

// G :
var G []node
var used []bool

// initG : グラフの初期化
func initG(N int) {
	G = make([]node, N)
}

func addEdge(from, to, cap int) {
	G[from].to = append(G[from].to, edge{to, cap, len(G[to].to)})
	G[to].to = append(G[to].to, edge{from, 0, len(G[from].to) - 1})
}

func dfs(v, t, f int) int {
	if v == t {
		return f
	}
	used[v] = true
	for i, e := range G[v].to {
		if used[e.to] || e.cap <= 0 {
			continue
		}
		d := dfs(e.to, t, min(f, e.cap))
		if d > 0 {
			G[v].to[i].cap -= d
			G[e.to].to[e.rev].cap += d
			return d
		}
	}
	return 0
}

const inf = math.MaxInt64 >> 10

func fordFulkerson(s, t, N int) int {
	flow := 0
	for {
		used = make([]bool, N)
		f := dfs(s, t, inf)
		if f == 0 {
			break
		}
		flow += f
		// out(f)
	}
	return flow
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	d := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getInt(), getInt()
	}
	for i := 0; i < N; i++ {
		c[i], d[i] = getInt(), getInt()
	}

	initG(N*2 + 2)

	for i := 0; i < N; i++ {
		addEdge(0, i+1, 1)
		addEdge(i+N+1, 2*N+1, 1)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if a[i] < c[j] && b[i] < d[j] {
				addEdge(i+1, j+N+1, 1)
			}
		}
	}
	res := fordFulkerson(0, 2*N+1, 2*N+2)
	out(res)
}
