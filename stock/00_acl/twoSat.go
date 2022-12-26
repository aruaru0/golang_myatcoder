type SccGraph struct {
	n     int
	edges [][2]int
}
type Csr struct {
	start []int
	elist []int
}
type TwoSat struct {
	n        int
	answer   []bool
	sccGraph *SccGraph
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
func newTwoSat(n int) *TwoSat {
	ts := new(TwoSat)
	ts.n = n
	ts.answer = make([]bool, n)
	ts.sccGraph = newSccGraph(n * 2)
	return ts
}
func (ts *TwoSat) AddClause(i int, f bool, j int, g bool) {
	ts.sccGraph.AddEdge(2*i+ts.judge(f, 0, 1), 2*j+ts.judge(g, 1, 0))
	ts.sccGraph.AddEdge(2*j+ts.judge(g, 0, 1), 2*i+ts.judge(f, 1, 0))
}
func (ts *TwoSat) Satisfiable() bool {
	_, id := ts.sccGraph.SccIds()
	for i := 0; i < ts.n; i++ {
		if id[i*2] == id[2*i+1] {
			return false
		}
		ts.answer[i] = id[2*i] < id[2*i+1]
	}
	return true
}
func (ts *TwoSat) judge(f bool, a int, b int) int {
	if f {
		return a
	} else {
		return b
	}
}
func (ts *TwoSat) Answer() []bool {
	return ts.answer
}