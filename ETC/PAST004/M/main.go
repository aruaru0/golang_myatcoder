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

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	n            int
}

// Dsu :
func Dsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
	}
	return &d
}

// Merge :
func (d DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

// Same :
func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader :
func (d DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

// Size :
func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

// Groups : original implement
func (d DSU) Groups() [][]int {
	m := make(map[int][]int)
	for i := 0; i < d.n; i++ {
		x := d.Leader(i)
		if x < 0 {
			m[i] = append(m[i], i)
		} else {
			m[x] = append(m[x], i)
		}
	}
	ret := make([][]int, len(m))
	idx := 0
	for _, e := range m {
		ret[idx] = make([]int, len(e))
		copy(ret[idx], e)
		idx++
	}
	return ret
}

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

type edge struct {
	a, b int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()
	lca := newLCA(N)
	m := make(map[edge]int)
	for i := 0; i < N-1; i++ {
		a, b := getI()-1, getI()-1
		lca.addEdge(a, b)
		m[edge{a, b}] = i
		m[edge{b, a}] = i
	}
	lca.build(0)

	uu := make([]int, Q)
	vv := make([]int, Q)
	cc := make([]int, Q)
	for i := 0; i < Q; i++ {
		uu[i], vv[i], cc[i] = getI()-1, getI()-1, getI()
	}

	ans := make([]int, N-1)
	root := make([]int, N)
	for i := 0; i < N; i++ {
		root[i] = i
	}

	un := Dsu(N)

	for i := Q - 1; i >= 0; i-- {
		u := uu[i]
		v := vv[i]
		c := cc[i]
		w := lca.lca(u, v)
		for j := 0; j < 2; j++ {
			for lca.depth(u) > lca.depth(w) {
				iu := lca.parent[0][u]
				if un.Same(u, iu) {
					u = root[un.Leader(u)]
					continue
				}
				ans[m[edge{u, iu}]] = c
				l1 := un.Leader(u)
				l2 := un.Leader(iu)
				newRoot := root[l1]
				if lca.depth(root[l1]) > lca.depth(root[l2]) {
					newRoot = root[l2]
				}
				un.Merge(u, iu)
				root[un.Leader(u)] = newRoot
				u = iu
			}
			u, v = v, u
		}
	}

	for _, e := range ans {
		out(e)
	}
}
