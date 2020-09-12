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

// 最小費用流
type edge struct {
	from, to, cap, cost, rev int
	CAP                      int
	idx                      int
}

type minFlow struct {
	V            int
	G            [][]edge
	dist         []int
	prevv, preve []int
	idx          int
}

const inf = int(1e15)

func newMinFlow(N int) *minFlow {
	var ret minFlow
	ret.V = N
	ret.G = make([][]edge, N)
	ret.dist = make([]int, N)
	ret.prevv = make([]int, N)
	ret.preve = make([]int, N)
	ret.idx = 0
	return &ret
}

func (m *minFlow) addEdge(from, to, cap, cost int) {
	m.G[from] = append(m.G[from], edge{from, to, cap, cost, len(m.G[to]), cap, m.idx})
	m.G[to] = append(m.G[to], edge{to, from, 0, -cost, len(m.G[from]) - 1, 0, -1})
	m.idx++
}

func (m *minFlow) minCostFlow(s, t, f int) int {
	res := 0
	for f > 0 {
		for i := 0; i < m.V; i++ {
			m.dist[i] = inf
		}
		m.dist[s] = 0
		update := true
		for update {
			update = false
			for v := 0; v < m.V; v++ {
				if m.dist[v] == inf {
					continue
				}
				for i, e := range m.G[v] {
					if e.cap > 0 && m.dist[e.to] > m.dist[v]+e.cost {
						m.dist[e.to] = m.dist[v] + e.cost
						m.prevv[e.to] = v
						m.preve[e.to] = i
						update = true
					}
				}
			}
		}
		if m.dist[t] == inf {
			return -1
		}
		d := f
		for v := t; v != s; v = m.prevv[v] {
			d = min(d, m.G[m.prevv[v]][m.preve[v]].cap)
		}
		f -= d
		res += d * m.dist[t]
		for v := t; v != s; v = m.prevv[v] {
			e := &m.G[m.prevv[v]][m.preve[v]]
			e.cap -= d
			m.G[v][e.rev].cap += d
		}
	}
	return res
}

func (m *minFlow) edges() []edge {
	edge := make([]edge, m.idx)
	for i := 0; i < m.V; i++ {
		for _, e := range m.G[i] {
			if e.idx >= 0 {
				edge[e.idx].from = e.from
				edge[e.idx].to = e.to
				edge[e.idx].cost = e.cost
				edge[e.idx].cap = e.CAP - e.cap // flow
				edge[e.idx].CAP = e.CAP
			}
		}
	}
	return edge
}

const big = 1000000000

func main() {
	sc.Split(bufio.ScanWords)
	n, k := getInt(), getInt()
	g := newMinFlow(2*n + 2)
	s, t := 2*n, 2*n+1
	g.addEdge(s, t, n*k, big)
	for i := 0; i < n; i++ {
		g.addEdge(s, i, k, 0)
		g.addEdge(n+i, t, k, 0)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a := getInt()
			g.addEdge(i, n+j, 1, big-a)
		}
	}
	result := g.minCostFlow(s, t, n*k)
	out(n*k*big - result)

	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			grid[i][j] = '.'
		}
	}
	edges := g.edges()
	for _, e := range edges {
		if e.from == s || e.to == t || e.cap == 0 {
			continue
		}
		grid[e.from][e.to-n] = 'X'
	}
	for i := 0; i < n; i++ {
		out(string(grid[i]))
	}
}
