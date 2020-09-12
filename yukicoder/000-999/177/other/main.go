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

func main() {
	sc.Split(bufio.ScanWords)
	W := getInt()
	N := getInt()
	J := getInts(N)
	M := getInt()
	C := getInts(M)
	s := 0
	e := M + N + 1
	n := 1
	m := N + 1
	initG(N + M + 2)
	for i := 0; i < N; i++ {
		addEdge(s, n+i, J[i])
	}
	for i := 0; i < M; i++ {
		addEdge(m+i, e, C[i])
	}
	for i := 0; i < M; i++ {
		Q := getInt()
		s := make([]bool, N)
		for j := 0; j < Q; j++ {
			x := getInt() - 1
			s[x] = true
		}
		for j := 0; j < N; j++ {
			if s[j] == false {
				addEdge(n+j, m+i, inf)
			}
		}
	}
	// for i := 0; i < M+N+2; i++ {
	// 	out(i, G[i])
	// }
	c := maxFlow(s, e)
	// out(c)
	if W <= c {
		out("SHIROBAKO")
		return
	}
	out("BANSAKUTSUKITA")
}
