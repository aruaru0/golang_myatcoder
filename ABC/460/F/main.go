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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
}

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

func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

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

func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

// -----------------------------
// ジェネリクス対応の SegTree
// -----------------------------
type Segtree[T any] struct {
	n      int
	size   int
	log    int
	d      []T
	e      func() T
	merger func(a, b T) T
}

func newSegtree[T any](v []T, e func() T, m func(a, b T) T) *Segtree[T] {
	seg := new(Segtree[T])
	seg.n = len(v)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]T, 2*seg.size)
	seg.e = e
	seg.merger = m
	for i := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		seg.d[seg.size+i] = v[i]
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.Update(i)
	}
	return seg
}

func (seg *Segtree[T]) Update(k int) {
	seg.d[k] = seg.merger(seg.d[2*k], seg.d[2*k+1])
}

func (seg *Segtree[T]) Set(p int, x T) {
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}

func (seg *Segtree[T]) Get(p int) T {
	return seg.d[p+seg.size]
}

func (seg *Segtree[T]) Prod(l, r int) T {
	sml, smr := seg.e(), seg.e()
	l += seg.size
	r += seg.size
	for l < r {
		if (l & 1) == 1 {
			sml = seg.merger(sml, seg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = seg.merger(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return seg.merger(sml, smr)
}

func (seg *Segtree[T]) AllProd() T {
	return seg.d[1]
}

func (seg *Segtree[T]) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

// -----------------------------
// Lowest Common Ancestor
// -----------------------------
type LcaData struct {
	val, idx int
}

type lca struct {
	G     [][]int
	vs    []int
	depth []int
	id    []int
	n     int
	k     int
	seg   *Segtree[LcaData]
}

const inf = int(1e16)

func newLCA(root, n int, node [][]int) *lca {
	var l lca
	l.n = n
	l.G = node
	l.vs = make([]int, n*2-1)
	l.depth = make([]int, n*2-1)
	l.id = make([]int, n)
	l.k = 0
	l.dfs(root, -1, 0)

	initData := make([]LcaData, n*2-1)
	for i, e := range l.depth {
		initData[i] = LcaData{val: e, idx: i}
	}
	e := func() LcaData {
		return LcaData{val: inf, idx: -1}
	}
	merger := func(a, b LcaData) LcaData {
		if a.val < b.val {
			return a
		}
		return b
	}
	l.seg = newSegtree(initData, e, merger)
	return &l
}

func (l *lca) lca(u, v int) int {
	res := l.seg.Prod(min(l.id[u], l.id[v]), max(l.id[u], l.id[v])+1)
	return l.vs[res.idx]
}

func (l *lca) dist(u, v int) int {
	lc := l.lca(u, v)
	return l.depth[l.id[u]] + l.depth[l.id[v]] - 2*l.depth[l.id[lc]]
}

func (l *lca) dfs(v, p, d int) {
	l.id[v] = l.k
	l.vs[l.k] = v
	l.depth[l.k] = d
	l.k++
	for i := 0; i < len(l.G[v]); i++ {
		if l.G[v][i] != p {
			l.dfs(l.G[v][i], v, d+1)
			l.vs[l.k] = v
			l.depth[l.k] = d
			l.k++
		}
	}
}

// -----------------------------
// 直径管理用の型定義
// -----------------------------
type DiamData struct {
	u, v int
	d    int
	has  bool
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N := getI()
	node := make([][]int, N)
	for i := 0; i < N-1; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	lcaObj := newLCA(0, N, node)

	isBlack := make([]bool, N)
	for i := 0; i < N; i++ {
		isBlack[i] = true
	}

	// 直径管理セグメント木の初期化
	initData := make([]DiamData, N)
	for i := 0; i < N; i++ {
		initData[i] = DiamData{u: i, v: i, d: 0, has: true}
	}

	e := func() DiamData {
		return DiamData{has: false}
	}

	merger := func(l, r DiamData) DiamData {
		if !l.has {
			return r
		}
		if !r.has {
			return l
		}

		best := l
		if r.d > best.d {
			best = r
		}

		if cand := lcaObj.dist(l.u, r.u); cand > best.d {
			best = DiamData{u: l.u, v: r.u, d: cand, has: true}
		}
		if cand := lcaObj.dist(l.u, r.v); cand > best.d {
			best = DiamData{u: l.u, v: r.v, d: cand, has: true}
		}
		if cand := lcaObj.dist(l.v, r.u); cand > best.d {
			best = DiamData{u: l.v, v: r.u, d: cand, has: true}
		}
		if cand := lcaObj.dist(l.v, r.v); cand > best.d {
			best = DiamData{u: l.v, v: r.v, d: cand, has: true}
		}

		return best
	}

	seg := newSegtree(initData, e, merger)

	Q := getI()
	for qi := 0; qi < Q; qi++ {
		x := getI() - 1
		isBlack[x] = !isBlack[x]

		if isBlack[x] {
			seg.Set(x, DiamData{u: x, v: x, d: 0, has: true})
		} else {
			seg.Set(x, DiamData{has: false})
		}

		out(seg.AllProd().d)
	}
}
