package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

type SccGraph struct {
	n     int
	edges [][2]int
}
type Csr struct {
	start []int
	elist []int
}

func newSccGraph(n int) *SccGraph {
	scc := new(SccGraph)
	scc.n = n
	return scc
}
func (scc *SccGraph) NumVertices() int {
	return scc.n
}
func (scc *SccGraph) AddEdge(from int, to int) {
	scc.edges = append(scc.edges, [2]int{from, to})
}
func (c *Csr) csr(n int, edges [][2]int) {
	c.start = make([]int, n+1)
	c.elist = make([]int, len(edges))
	for _, e := range edges {
		c.start[e[0]+1]++
	}
	for i := 1; i <= n; i++ {
		c.start[i] += c.start[i-1]
	}
	counter := make([]int, n+1)
	copy(counter, c.start)
	for _, e := range edges {
		c.elist[counter[e[0]]] = e[1]
		counter[e[0]]++
	}
}
func (scc *SccGraph) SccIds() (int, []int) {
	g := new(Csr)
	g.csr(scc.n, scc.edges)
	nowOrd, groupNum := 0, 0
	visited, low := make([]int, 0, scc.n), make([]int, scc.n)
	ord, ids := make([]int, scc.n), make([]int, scc.n)
	for i := 0; i < scc.n; i++ {
		ord[i] = -1
	}
	var dfs func(v int)
	dfs = func(v int) {
		low[v], ord[v] = nowOrd, nowOrd
		nowOrd++
		visited = append(visited, v)
		for i := g.start[v]; i < g.start[v+1]; i++ {
			to := g.elist[i]
			if ord[to] == -1 {
				dfs(to)
				low[v] = scc.min(low[v], low[to])
			} else {
				low[v] = scc.min(low[v], ord[to])
			}
		}
		if low[v] == ord[v] {
			for {
				u := visited[len(visited)-1]
				visited = visited[:len(visited)-1]
				ord[u] = scc.n
				ids[u] = groupNum
				if u == v {
					break
				}
			}
			groupNum++
		}
	}
	for i := 0; i < scc.n; i++ {
		if ord[i] == -1 {
			dfs(i)
		}
	}
	for i := 0; i < len(ids); i++ {
		ids[i] = groupNum - 1 - ids[i]
	}
	return groupNum, ids
}
func (scc *SccGraph) min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func (scc *SccGraph) Scc() [][]int {
	groupNum, ids := scc.SccIds()
	counts := make([]int, groupNum)
	for _, x := range ids {
		counts[x]++
	}
	groups := make([][]int, groupNum)
	for i := 0; i < groupNum; i++ {
		groups[i] = make([]int, 0, counts[i])
	}
	for i := 0; i < scc.n; i++ {
		groups[ids[i]] = append(groups[ids[i]], i)
	}
	return groups
}

type MinCostFlow struct {
	n   int
	pos [][2]int
	g   [][]_Edge
}
type _Edge struct {
	to   int
	rev  int
	capa int
	cost int
}
type Edge struct {
	from int
	to   int
	capa int
	flow int
	cost int
}

func newMinCostFlow(n int) *MinCostFlow {
	return &MinCostFlow{n: n, g: make([][]_Edge, n)}
}
func (mcf *MinCostFlow) AddEdge(from, to, capa, cost int) int {
	m := len(mcf.pos)
	mcf.pos = append(mcf.pos, [2]int{from, len(mcf.g[from])})
	mcf.g[from] = append(mcf.g[from], _Edge{to, len(mcf.g[to]), capa, cost})
	mcf.g[to] = append(mcf.g[to], _Edge{from, len(mcf.g[from]) - 1, 0, -cost})
	return m
}
func (mcf *MinCostFlow) GetEdge(i int) Edge {
	e := mcf.g[mcf.pos[i][0]][mcf.pos[i][1]]
	re := mcf.g[e.to][e.rev]
	return Edge{mcf.pos[i][0], e.to, e.capa + re.capa, re.capa, e.cost}
}
func (mcf *MinCostFlow) Edges() []Edge {
	m := len(mcf.pos)
	res := make([]Edge, m)
	for i := 0; i < m; i++ {
		res[i] = mcf.GetEdge(i)
	}
	return res
}
func (mcf *MinCostFlow) Flow(s, t int) [2]int {
	res := mcf.Slope(s, t)
	return res[len(res)-1]
}
func (mcf *MinCostFlow) FlowL(s, t, flowLim int) [2]int {
	res := mcf.SlopeL(s, t, flowLim)
	return res[len(res)-1]
}
func (mcf *MinCostFlow) Slope(s, t int) [][2]int {
	return mcf.SlopeL(s, t, int(1e+18))
}
func (mcf *MinCostFlow) SlopeL(s, t, flowLim int) [][2]int {
	dual, dist := make([]int, mcf.n), make([]int, mcf.n)
	pv, pe := make([]int, mcf.n), make([]int, mcf.n)
	vis := make([]bool, mcf.n)
	dualRef := func() bool {
		for i := 0; i < mcf.n; i++ {
			dist[i], pv[i], pe[i] = int(1e+18), -1, -1
			vis[i] = false
		}
		pq := make(PriorityQueue, 0)
		heap.Init(&pq)
		item := &Item{value: s, priority: 0}
		dist[s] = 0
		heap.Push(&pq, item)
		for pq.Len() != 0 {
			v := heap.Pop(&pq).(*Item).value
			if vis[v] {
				continue
			}
			vis[v] = true
			if v == t {
				break
			}
			for i := 0; i < len(mcf.g[v]); i++ {
				e := mcf.g[v][i]
				if vis[e.to] || e.capa == 0 {
					continue
				}
				cost := e.cost - dual[e.to] + dual[v]
				if dist[e.to]-dist[v] > cost {
					dist[e.to] = dist[v] + cost
					pv[e.to] = v
					pe[e.to] = i
					item := &Item{value: e.to, priority: dist[e.to]}
					heap.Push(&pq, item)
				}
			}
		}
		if !vis[t] {
			return false
		}
		for v := 0; v < mcf.n; v++ {
			if !vis[v] {
				continue
			}
			dual[v] -= dist[t] - dist[v]
		}
		return true
	}
	flow, cost, prevCost := 0, 0, -1
	res := make([][2]int, 0, mcf.n)
	res = append(res, [2]int{flow, cost})
	for flow < flowLim {
		if !dualRef() {
			break
		}
		c := flowLim - flow
		for v := t; v != s; v = pv[v] {
			c = mcf.Min(c, mcf.g[pv[v]][pe[v]].capa)
		}
		for v := t; v != s; v = pv[v] {
			mcf.g[pv[v]][pe[v]].capa -= c
			mcf.g[v][mcf.g[pv[v]][pe[v]].rev].capa += c
		}
		d := -dual[s]
		flow += c
		cost += c * d
		if prevCost == d {
			res = res[:len(res)-1]
		}
		res = append(res, [2]int{flow, cost})
		prevCost = cost
	}
	return res
}
func (mcf *MinCostFlow) Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Item struct {
	value    int
	priority int
	index    int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type pair struct {
	first, second int
}

// 負のコストの部分でつまずいた
// 解説通りの実装
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m, k := getI(), getI(), getI()
	x := make([]int, 0)
	es := make([]pair, 0)
	v0 := -1
	{
		// SCCを用いて、DAGに変換する
		g := newSccGraph(n)
		es_ := make([]pair, 0)
		for i := 0; i < m; i++ {
			a, b := getI()-1, getI()-1
			g.AddEdge(a, b)
			es_ = append(es_, pair{a, b})
		}
		x_ := getInts(n)
		gr := g.Scc()
		gi := make([]int, n)
		x = make([]int, len(gr))
		for i := 0; i < len(gr); i++ {
			for _, v := range gr[i] {
				gi[v] = i
				x[i] += x_[v]
			}
		}
		for _, p := range es_ {
			a, b := p.first, p.second
			if gi[a] == gi[b] {
				continue
			}
			es = append(es, pair{gi[a], gi[b]})
		}
		v0 = gi[0]
		n = len(x)
	}

	sv := n * 2
	tv := sv + 1

	g := newMinCostFlow(tv + 1)
	for i := 0; i < n; i++ {
		g.AddEdge(i, n+i, 1, 0)
		g.AddEdge(i, n+i, k, x[i])
	}

	xs := make([]int, n+1)
	for i := 0; i < n; i++ {
		xs[i+1] = xs[i] + x[i]
	}

	g.AddEdge(sv, v0, k, xs[v0])
	for i := 0; i < n; i++ {
		g.AddEdge(n+i, tv, k, xs[n]-xs[i+1])
	}
	for _, p := range es {
		a, b := p.first, p.second
		g.AddEdge(n+a, b, k, xs[b]-xs[a+1])
	}

	ans := xs[n]*k - g.FlowL(sv, tv, k)[1]
	out(ans)
}
