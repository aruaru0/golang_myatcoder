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

// 値を圧縮した配列を返す
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

// ==========================================
// 【追加】セグメント木上の二分探索 (MaxRight)
// ==========================================
// l から右に向かって「区間積が条件 f を満たし続ける最大の右端 r」を返します。
// つまり、条件 f が初めて false になるインデックス（＝探したい要素の左端）が返ります。
// 最後まで条件が true の場合は seg.n を返します。計算量: O(log N)
func (seg *Segtree[T]) MaxRight(l int, f func(T) bool) int {
	if l == seg.n {
		return seg.n
	}
	l += seg.size
	sm := seg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !f(seg.merger(sm, seg.d[l])) {
			for l < seg.size {
				l = 2 * l
				if f(seg.merger(sm, seg.d[l])) {
					sm = seg.merger(sm, seg.d[l])
					l++
				}
			}
			return l - seg.size
		}
		sm = seg.merger(sm, seg.d[l])
		l++
		if (l & -l) == l {
			break
		}
	}
	return seg.n
}

// type Node struct {
// 	maxVal int
// 	index  int
// }
// func main() {
// 	// (入出力の準備やNの読み込みなどは省略)
// 	a := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3} // 例
// 	N := len(a)

// 	const inf = int(1e18)

// 	e := func() Node {
// 		return Node{maxVal: -inf, index: -1}
// 	}

// 	merger := func(a, b Node) Node {
// 		if a.maxVal >= b.maxVal {
// 			return a
// 		}
// 		return b
// 	}

// 	v := make([]Node, N)
// 	for i := 0; i < N; i++ {
// 		v[i] = Node{maxVal: a[i], index: i}
// 	}

// 	seg := newSegtree[Node](v, e, merger)

// 	// 例えば区間 [2, 7) の最大値とそのインデックスを取得
// 	res := seg.Prod(2, 7)

// 	// res は Node 型で返ってくるので、そのまま中身にアクセス可能！
// 	fmt.Printf("区間内の最大値は %d (元の配列のインデックス: %d) です\n", res.maxVal, res.index)
//     // 出力結果: 区間内の最大値は 9 (元の配列のインデックス: 5) です
// }

// 宣言はこんな感じ
// seg := newSegtree[S](v, e, merger)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	h := getInts(N)

	e := func() int {
		return -1
	}
	merger := func(a, b int) int {
		return max(a, b)
	}

	seg := newSegtree(h, e, merger)

	const inf = int(1e18)
	e1 := func() int {
		return inf
	}
	merger1 := func(a, b int) int {
		return min(a, b)
	}

	segMin := newSegtree(h, e1, merger1)

	ans := 0
	for i := 0; i < N; i++ {
		idx := seg.MaxRight(i, func(v int) bool {
			return v <= h[i]
		})
		if idx != N {
			v := segMin.Prod(i, idx)
			ans += v
			// out(idx, v)
		}
	}
	out(ans)
}
