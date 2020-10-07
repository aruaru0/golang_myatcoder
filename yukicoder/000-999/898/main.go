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

var node [][]int
var cost [][]int
var N int

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N = getInt()
	node = make([][]int, N)
	cost = make([][]int, N)
	for i := 0; i < N-1; i++ {
		f, t, c := getInt(), getInt(), getInt()
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
		cost[f] = append(cost[f], c)
		cost[t] = append(cost[t], c)
	}
	// out(node)
	// out(cost)
	dist = make([]int, N)
	dfs(0, -1, 0)
	l := newLCA(0, N, node)

	Q := getInt()
	// out(dist)
	for i := 0; i < Q; i++ {
		a, b, c := getInt(), getInt(), getInt()
		x := l.lca(a, b)
		y := l.lca(a, c)
		z := l.lca(b, c)
		p := dist[a] + dist[b] + dist[c] - dist[x] - dist[y] - dist[z]
		out(p)
	}
	// out(l)
	// out(dist)
}

var dist []int

func dfs(v, p, c int) {
	dist[v] = c
	for i, e := range node[v] {
		if e == p {
			continue
		}
		dfs(e, v, c+cost[v][i])
	}
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
