package main

import (
	"bufio"
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

// UnionFind高速
type UnionFind struct {
	d []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.d = make([]int, N)
	for i := 0; i < N; i++ {
		u.d[i] = -1
	}
	return u
}

func (p *UnionFind) root(x int) int {
	if p.d[x] < 0 {
		return x
	}
	p.d[x] = p.root(p.d[x]) // ※親を検索したら、リンクを更新。書き換え時要注意
	return p.d[x]
}

func (p *UnionFind) unite(x, y int) bool {
	x = p.root(x)
	y = p.root(y)
	if x == y {
		return false
	}
	if p.d[x] > p.d[y] {
		x, y = y, x
	}
	p.d[x] += p.d[y]
	p.d[y] = x
	return true
}

func (p *UnionFind) same(x, y int) bool {
	return p.root(x) == p.root(y)
}

func (p *UnionFind) size(x int) int {
	return -p.d[p.root(x)]
}

var n int
var tree [][]int
var dist []int
var dist2 []int
var ans []int
var cost []int
var used []bool

func dfs1(v, p int) {
	used[v] = true
	res := 0
	res2 := cost[v]
	for _, u := range tree[v] {
		if u == p {
			continue
		}
		dfs1(u, v)
		res += dist[u] + dist2[u]
		res2 += dist2[u]
	}
	dist2[v] = res2
	dist[v] = res
}

func dfs2(v, d_par, d_par2, p int) int {
	used[v] = true
	ans[v] = d_par + dist[v] + d_par2
	res := ans[v]
	var f = func(a, b int) int {
		return a + b
	}
	lf := make([]int, len(tree[v]))
	ri := make([]int, len(tree[v]))
	lf2 := make([]int, len(tree[v]))
	ri2 := make([]int, len(tree[v]))

	for i := 1; i < len(tree[v]); i++ {
		if tree[v][i-1] == p {
			lf[i] = f(d_par, lf[i-1])
		} else {
			lf[i] = f(dist[tree[v][i-1]], lf[i-1])
		}
		if tree[v][i-1] == p {
			lf2[i] = f(d_par2, lf2[i-1])
		} else {
			lf2[i] = f(dist2[tree[v][i-1]], lf2[i-1])
		}
	}
	for i := len(tree[v]) - 2; i >= 0; i-- {
		if tree[v][i+1] == p {
			ri[i] = f(d_par, ri[i+1])
		} else {
			ri[i] = f(dist[tree[v][i+1]], ri[i+1])
		}
		if tree[v][i+1] == p {
			ri2[i] = f(d_par2, ri2[i+1])
		} else {
			ri2[i] = f(dist2[tree[v][i+1]], ri2[i+1])
		}
	}
	for i := 0; i < len(tree[v]); i++ {
		if tree[v][i] == p {
			continue
		}
		res = min(res, dfs2(tree[v][i], f(lf[i], ri[i])+f(lf2[i], ri2[i]), f(lf2[i], ri2[i])+cost[v], v))
	}
	return res
}

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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n = getI()
	m, q := getI(), getI()
	tree = make([][]int, n)
	dist = make([]int, n)
	dist2 = make([]int, n)
	ans = make([]int, n)
	cost = make([]int, n)
	uf := newUnionFind(n)
	uf2 := newUnionFind(n)
	lca := newLCA(n)
	sum := 0

	for i := 0; i < m; i++ {
		u, v := getI()-1, getI()-1
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
		uf.unite(u, v)
		uf2.unite(u, v)
		lca.addEdge(u, v)
	}
	for i := 1; i < n; i++ {
		if !uf2.same(0, i) {
			uf2.unite(0, i)
			lca.addEdge(0, i)
		}
	}
	lca.build(0)
	for i := 0; i < q; i++ {
		a, b := getI()-1, getI()-1
		if uf.same(a, b) {
			sum += lca.dist(a, b)
		} else {
			cost[a]++
			cost[b]++
		}
	}
	used = make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		dfs1(i, -1)
	}
	used = make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		sum += dfs2(i, 0, 0, -1)
	}
	out(sum)
}
