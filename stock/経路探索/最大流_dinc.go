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

//------------------------------------------------
// 最大流(Dinic)
//   initG(N) init size
//   addEdge(f,t,c) add edge from(f) to(t) cap(c)
//   maxFlow(s, e) calc cost from s to e
const inf = math.MaxInt64 >> 10

type edge struct {
	to, cap, rev int
}

// G :
var G [][]edge
var level []int
var iter []int

func initG(N int) {
	G = make([][]edge, N)
	level = make([]int, N)
	iter = make([]int, N)
}

func addEdge(from, to, cap int) {
	G[from] = append(G[from], edge{to, cap, len(G[to])})
	G[to] = append(G[to], edge{from, 0, len(G[from]) - 1})
}

func bfs(s int) {
	for i := 0; i < len(level); i++ {
		level[i] = -1
	}
	que := []int{s}
	level[s] = 0
	for len(que) > 0 {
		v := que[0]
		que = que[1:]
		for _, e := range G[v] {
			if e.cap > 0 && level[e.to] < 0 {
				level[e.to] = level[v] + 1
				que = append(que, e.to)
			}
		}
	}
}

func dfs(v, t, f int) int {
	if v == t {
		return f
	}
	for i, e := range G[v] {
		if e.cap > 0 && level[v] < level[e.to] {
			d := dfs(e.to, t, min(f, e.cap))
			if d > 0 {
				G[v][i].cap -= d
				G[e.to][e.rev].cap += d
				return d
			}
		}
	}
	return 0
}

func maxFlow(s, t int) int {
	flow := 0
	N := len(iter)
	for {
		bfs(s)
		if level[t] < 0 {
			return flow
		}
		iter = make([]int, N)
		for {
			f := dfs(s, t, inf)
			if f <= 0 {
				break
			}
			flow += f
		}
	}
}

// sample: yukicoder No.119
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
	}
	M := getInt()
	for i := 0; i < M; i++ {
		d, e := getInt(), getInt()
		addEdge(d+1, e+1+N, inf)
	}
	ans -= maxFlow(s, e)
	out(ans)
}
