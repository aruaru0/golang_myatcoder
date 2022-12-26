import "math"

//
// LCAの簡易バージョン（２点の距離計算のついたバージョン）
//

// l := newLCA(n)
// l.addEdge(u,v)
// l.build(root) ※実行が必要
// root := l.lca(u,v)
// dist := l.dist(u,v)

type lca struct {
	n      int
	log    int
	parent [][]int
	dep    []int
	G      [][]int
}

func newLCA(n int) *lca {
	var ret lca
	ret.n = n
	ret.log = int(math.Log2(float64(n))) + 1
	ret.parent = make([][]int, ret.log)
	for i := 0; i < ret.log; i++ {
		ret.parent[i] = make([]int, n)
	}
	ret.dep = make([]int, n)
	ret.G = make([][]int, n)
	return &ret
}

func (l *lca) dfs(v, p, d int) {
	l.parent[0][v] = p
	l.dep[v] = d
	for _, to := range l.G[v] {
		if to == p {
			continue
		}
		l.dfs(to, v, d+1)
	}
}

func (l *lca) addEdge(from, to int) {
	l.G[from] = append(l.G[from], to)
	l.G[to] = append(l.G[to], from)
}

func (l *lca) build(root int) {
	l.dfs(root, -1, 0)
	for k := 0; k+1 < l.log; k++ {
		for v := 0; v < l.n; v++ {
			if l.parent[k][v] < 0 {
				l.parent[k+1][v] = -1
			} else {
				l.parent[k+1][v] = l.parent[k][l.parent[k][v]]
			}
		}
	}
}

func (l *lca) depth(v int) int {
	return l.dep[v]
}

func (l *lca) lca(u, v int) int {
	if l.dep[u] > l.dep[v] {
		u, v = v, u
	}
	for k := 0; k < l.log; k++ {
		if (l.dep[v]-l.dep[u])>>k&1 == 1 {
			v = l.parent[k][v]
		}
	}
	if u == v {
		return u
	}
	for k := l.log - 1; k >= 0; k-- {
		if l.parent[k][u] != l.parent[k][v] {
			u = l.parent[k][u]
			v = l.parent[k][v]
		}
	}
	return l.parent[0][u]
}

func (l *lca) dist(u, v int) int {
	return l.dep[u] + l.dep[v] - 2*l.dep[l.lca(u, v)]
}
