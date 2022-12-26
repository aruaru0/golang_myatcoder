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
	}
	return flow
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()

	initG(N*2 + 2)
	s := 0
	e := N*2 + 2 - 1
	ans := 0
	for i := 0; i < N; i++ {
		b, c := getInt(), getInt()
		addEdge(s, i+1, b)
		addEdge(i+N+1, e, c)
		addEdge(i+1, i+N+1, inf)
		ans += b + c
		// out(s, i+1, b)
		// out(i+N+1, e, c)
		// out(i+1, i+N+1, inf)
	}
	M := getInt()
	for i := 0; i < M; i++ {
		d, e := getInt(), getInt()
		addEdge(d+1, e+1+N, inf)
		// out(d+1, e+1+N, inf)
	}
	// out(s, e, 2*N+2)
	ans -= fordFulkerson(s, e, 2*N+2)
	out(ans)
}
