type Edge struct {
	from int
	to   int
	capa int
	flow int
}
type _Edge struct {
	to   int
	rev  int
	capa int
}
type MaxFlow struct {
	n   int
	pos [][2]int
	g   [][]_Edge
}

func newMaxFlow(n int) *MaxFlow {
	return &MaxFlow{
		n: n,
		g: make([][]_Edge, n),
	}
}
func (mf *MaxFlow) smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (mf *MaxFlow) AddEdge(from, to, capa int) int {
	m := len(mf.pos)
	mf.pos = append(mf.pos, [2]int{from, len(mf.g[from])})
	mf.g[from] = append(mf.g[from], _Edge{to, len(mf.g[to]), capa})
	mf.g[to] = append(mf.g[to], _Edge{from, len(mf.g[from]) - 1, 0})
	return m
}
func (mf *MaxFlow) GetEdge(i int) Edge {
	_e := mf.g[mf.pos[i][0]][mf.pos[i][1]]
	_re := mf.g[_e.to][_e.rev]
	return Edge{mf.pos[i][0], _e.to, _e.capa + _re.capa, _re.capa}
}
func (mf *MaxFlow) EdgesList() []Edge {
	m := len(mf.pos)
	result := make([]Edge, 0, m)
	for i := 0; i < m; i++ {
		result = append(result, mf.GetEdge(i))
	}
	return result
}
func (mf *MaxFlow) ChangeEdge(i, newCapa, newFlow int) {
	_e := &mf.g[mf.pos[i][0]][mf.pos[i][1]]
	_re := &mf.g[_e.to][_e.rev]
	_e.capa = newCapa - newFlow
	_re.capa = newFlow
}
func (mf *MaxFlow) Flow(s, t int) int {
	return mf.FlowL(s, t, int(1e+18))
}
func (mf *MaxFlow) FlowL(s, t, flowLim int) int {
	level := make([]int, mf.n)
	iter := make([]int, mf.n)
	bfs := func() {
		for i, _ := range level {
			level[i] = -1
		}
		level[s] = 0
		q := make([]int, 0, mf.n)
		q = append(q, s)
		for len(q) != 0 {
			v := q[0]
			q = q[1:]
			for _, e := range mf.g[v] {
				if e.capa == 0 || level[e.to] >= 0 {
					continue
				}
				level[e.to] = level[v] + 1
				if e.to == t {
					return
				}
				q = append(q, e.to)
			}
		}
	}
	var dfs func(v, up int) int
	dfs = func(v, up int) int {
		if v == s {
			return up
		}
		res := 0
		lv := level[v]
		for ; iter[v] < len(mf.g[v]); iter[v]++ {
			e := &mf.g[v][iter[v]]
			if lv <= level[e.to] || mf.g[e.to][e.rev].capa == 0 {
				continue
			}
			d := dfs(e.to, mf.smaller(up-res, mf.g[e.to][e.rev].capa))
			if d <= 0 {
				continue
			}
			mf.g[v][iter[v]].capa += d
			mf.g[e.to][e.rev].capa -= d
			res += d
			if res == up {
				break
			}
		}
		return res
	}
	flow := 0
	for flow < flowLim {
		bfs()
		if level[t] == -1 {
			break
		}
		for i, _ := range iter {
			iter[i] = 0
		}
		for flow < flowLim {
			f := dfs(t, flowLim-flow)
			if f == 0 {
				break
			}
			flow += f
		}
	}
	return flow
}
func (mf *MaxFlow) MinCut(s int) []bool {
	visited := make([]bool, mf.n)
	q := make([]int, 0, mf.n)
	q = append(q, s)
	for len(q) != 0 {
		p := q[0]
		q = q[1:]
		visited[p] = true
		for _, e := range mf.g[p] {
			if e.capa > 0 && !visited[e.to] {
				visited[e.to] = true
				q = append(q, e.to)
			}
		}
	}
	return visited
}
