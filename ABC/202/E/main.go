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

/*
Lowest Common Ansestor
*/
type lca struct {
	G     [][]int
	vs    []int
	depth []int
	id    []int
	n     int
	k     int
	seg   *SegmentTree
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

	l.seg = SegtreeInit(n*2-1, Data{inf, 0})
	for i, e := range l.depth {
		l.seg.Set(i, Data{e, i})
	}
	l.seg.Update()
	return &l
}

func (l *lca) lca(u, v int) int {
	idx := l.seg.Query(min(l.id[u], l.id[v]), max(l.id[u], l.id[v]))
	return l.vs[idx.idx]
}

func (l *lca) dfs(v, p, d int) {
	l.id[v] = l.k
	l.vs[l.k] = v
	l.depth[l.k] = d
	l.k++
	// out(v, p, l.G[v], l.k)
	for i := 0; i < len(l.G[v]); i++ {
		if l.G[v][i] != p {
			l.dfs(l.G[v][i], v, d+1)
			l.vs[l.k] = v
			l.depth[l.k] = d
			l.k++
		}
	}
}

/*
  セグメント木(2020.05.24作成)
	SegtreeInitで初期化
	Setで値設定し、Updateで木作成
	Getで値を取得
	UpdateAtで個別のアイテムを更新
	Queryで区間の値を取得
	compareで比較方法変更可能
*/

// Data :
// Data型をstructすれば複数データが持てる
// compareも変更すること！
type Data struct {
	val, idx int
}

// SegmentTree :
type SegmentTree struct {
	inf    Data
	d      []Data
	offset int
}

// SegtreeInit :　nが要素数、valが初期値
func SegtreeInit(n int, val Data) *SegmentTree {
	var ret SegmentTree
	size := 1
	for size < n {
		size *= 2
	}
	ret.d = make([]Data, size*2)
	for i := 1; i < size*2; i++ {
		ret.d[i] = val
	}
	ret.offset = size
	ret.inf = val
	return &ret
}

// Set : 要素に値をセット（※木は更新されない）
func (s *SegmentTree) Set(idx int, val Data) {
	s.d[s.offset+idx] = val
}

// Get : 要素に値を取得
func (s *SegmentTree) Get(idx int) Data {
	return s.d[s.offset+idx]
}

// Update :
func (s *SegmentTree) Update() {
	N := s.offset
	off := s.offset

	for N > 1 {
		for i := off; i < off+N; i += 2 {
			p := i / 2
			l := i
			r := i + 1
			s.d[p] = s.compare(s.d[l], s.d[r])
		}
		off /= 2
		N /= 2
	}
}

// querySub :
// a, b ... 範囲
func (s *SegmentTree) querySub(a, b, k, l, r int) Data {
	if r <= a || b <= l {
		return s.inf
	}
	if a <= l && r <= b {
		return s.d[k]
	}
	return s.compare(
		s.querySub(a, b, k*2, l, (l+r)/2),
		s.querySub(a, b, k*2+1, (l+r)/2, r))
}

// Query :
// a, b ... 範囲 a <= x < bの範囲で検索
// [a, b)となっているのに注意
func (s *SegmentTree) Query(a, b int) Data {
	return s.querySub(a, b, 1, 0, s.offset)
}

// UpdateAt :
func (s *SegmentTree) UpdateAt(n int, val Data) {
	pos := s.offset + n
	s.d[pos] = val
	for pos > 1 {
		p := pos / 2
		l := p * 2
		r := p*2 + 1
		s.d[p] = s.compare(s.d[l], s.d[r])
		pos /= 2
	}
}

// compare :
// 比較関数（ここで比較方法を設定）
// ※min,maxを入れ替えるときなどは、Initの設定注意
func (s *SegmentTree) compare(l, r Data) Data {
	// 区間の合計の場合はinitを0にして下記
	// return l + r

	// 区間のminの場合はinfに最大値以上を設定して下記
	if l.val < r.val {
		return l
	}
	return r
}

func dfs(c, p, cnt int) {
	dist[c] = cnt
	route = append(route, c)
	for _, e := range node[c] {
		if e == p {
			continue
		}
		dfs(e, c, cnt+1)
	}
	route = append(route, c)
}

var node [][]int
var dist []int
var route []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	p := getInts(N - 1)
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		f := p[i] - 1
		t := i + 1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
	}

	route = make([]int, 0)
	dist = make([]int, N)
	dfs(0, -1, 0)

	inout := make([][]int, N)
	depth := make(map[int][]int)
	used := make([]bool, N)
	for i, e := range route {
		inout[e] = append(inout[e], i)
		if !used[e] {
			depth[dist[e]] = append(depth[dist[e]], i)
			used[e] = true
		}
	}
	// out(inout)
	// out(depth)
	Q := getI()
	for i := 0; i < Q; i++ {
		u, d := getI()-1, getI()
		l, r := inout[u][0], inout[u][1]
		// out("----")
		// out(l, r, d)
		x := lowerBound(depth[d], l)
		y := lowerBound(depth[d], r)
		// out(x, y)
		out(y - x)
	}
}
