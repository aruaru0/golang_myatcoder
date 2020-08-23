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

const inf = math.MaxInt64 >> 10

type edge struct {
	t, c, r int
}

type maxflow struct {
	node [][]edge
	used []bool
	N    int
}

func newMaxflow(n int) maxflow {
	var ret maxflow
	ret.node = make([][]edge, n)
	ret.N = n
	return ret
}

func (m *maxflow) addEdge(f, t, c int) {
	m.node[f] = append(m.node[f], edge{t, c, len(m.node[t])})
	m.node[t] = append(m.node[t], edge{f, 0, len(m.node[f]) - 1})
}

func (m *maxflow) dfs(v, t, f int) int {
	if v == t {
		return f
	}
	m.used[v] = true
	for i, e := range m.node[v] {
		if m.used[e.t] == false && e.c > 0 {
			d := m.dfs(e.t, t, min(f, e.c))
			if d > 0 {
				m.node[v][i].c -= d
				m.node[e.t][e.r].c += d
				return d
			}
		}
	}
	return 0
}

func (m *maxflow) maxFlow(f, t int) int {
	flow := 0
	for {
		m.used = make([]bool, m.N)
		f := m.dfs(f, t, inf)
		if f == 0 {
			return flow
		}
		flow += f
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	mf := newMaxflow(2*N + 2)
	f := 2 * N
	t := 2*N + 1
	ans := 0
	for i := 0; i < N; i++ {
		b, c := getInt(), getInt()
		ans += b + c
		mf.addEdge(f, i, b)
		mf.addEdge(i, N+i, inf)
		mf.addEdge(N+i, t, c)
	}
	M := getInt()
	for i := 0; i < M; i++ {
		d, e := getInt(), getInt()
		mf.addEdge(d, N+e, inf)
	}

	// for i := 0; i < 2*N+2; i++ {
	// 	out(i, mf.node[i])
	// }
	out(ans - mf.maxFlow(f, t))
}
