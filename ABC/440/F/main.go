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

// -----------------------------------------------------------------------------
// セグメント木ライブラリ
// -----------------------------------------------------------------------------
// S はセグメント木のノードが持つデータ
type S struct {
	cnt int   // 区間内の要素数
	sum int64 // 区間内の要素の和
}

type E func() S
type Merger func(a, b S) S
type Compare func(v int) bool

type Segtree[T any] struct {
	n      int
	size   int
	log    int
	d      []S
	e      E
	merger Merger
}

func newSegtree[T any](v []S, e E, m Merger) *Segtree[T] {
	seg := new(Segtree[T])
	seg.n = len(v)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]S, 2*seg.size)
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

func (seg *Segtree[T]) Set(p int, x S) {
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}

func (seg *Segtree[T]) Get(p int) S {
	return seg.d[p+seg.size]
}

func (seg *Segtree[T]) Prod(l, r int) S {
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

func (seg *Segtree[T]) AllProd() S {
	return seg.d[1]
}

func (seg *Segtree[T]) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

// Add は値 p の個数を diff だけ増やし、総和を更新する
func (seg *Segtree[T]) Add(p int, diff int) {
	current := seg.Get(p)
	newS := S{
		cnt: current.cnt + diff,
		sum: current.sum + int64(diff)*int64(p),
	}
	seg.Set(p, newS)
}

// QueryTopSum は大きい方から k 個の要素の和を返す
func (seg *Segtree[T]) QueryTopSum(k int) int64 {
	if k <= 0 {
		return 0
	}
	// 全体の個数より多く要求された場合は全ての和を返す
	if k >= seg.d[1].cnt {
		return seg.d[1].sum
	}
	return seg.queryTopSumRec(1, k)
}

func (seg *Segtree[T]) queryTopSumRec(node, k int) int64 {
	// 葉の場合
	if node >= seg.size {
		val := int64(node - seg.size)
		// この値を持つ要素のうち、必要な個数分だけ足す
		take := k
		if take > seg.d[node].cnt {
			take = seg.d[node].cnt
		}
		return val * int64(take)
	}

	// 右の子（大きい値側）を見る
	right := 2*node + 1
	rightCnt := seg.d[right].cnt

	if rightCnt >= k {
		// 右側だけで足りる場合
		return seg.queryTopSumRec(right, k)
	} else {
		// 右側を全部取って、左側から不足分を取る
		return seg.d[right].sum + seg.queryTopSumRec(2*node, k-rightCnt)
	}
}

// GetKthLargest は大きい方から k 番目の値(インデックス)を返す
func (seg *Segtree[T]) GetKthLargest(k int) int {
	if k <= 0 || k > seg.d[1].cnt {
		return -1
	}
	node := 1
	for node < seg.size {
		right := 2*node + 1
		rightCnt := seg.d[right].cnt
		if rightCnt >= k {
			node = right
		} else {
			k -= rightCnt
			node = 2 * node
		}
	}
	return node - seg.size
}

// -----------------------------------------------------------------------------
// メイン処理
// -----------------------------------------------------------------------------

const MAX_A = 1000005

// 単位元
func e() S {
	return S{0, 0}
}

// マージ関数
func merger(a, b S) S {
	return S{
		cnt: a.cnt + b.cnt,
		sum: a.sum + b.sum,
	}
}

type Horse struct {
	A, B int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N, Q := getI(), getI()

	// セグメント木の初期化
	// 値 A_i は 1 ~ 10^6 なので、サイズは MAX_A 必要
	initialVec := make([]S, MAX_A)
	allTree := newSegtree[any](initialVec, e, merger)
	type2Tree := newSegtree[any](initialVec, e, merger)

	horses := make([]Horse, N+1)
	var sumAll int64 = 0
	countN2 := 0

	for i := 1; i <= N; i++ {
		a, b := getI(), getI()
		horses[i] = Horse{a, b}

		allTree.Add(a, 1)
		sumAll += int64(a)
		if b == 2 {
			type2Tree.Add(a, 1)
			countN2++
		}
	}

	for i := 0; i < Q; i++ {
		w, x, y := getI(), getI(), getI()

		oldH := horses[w]
		allTree.Add(oldH.A, -1)
		sumAll -= int64(oldH.A)
		if oldH.B == 2 {
			type2Tree.Add(oldH.A, -1)
			countN2--
		}

		horses[w] = Horse{x, y}
		allTree.Add(x, 1)
		sumAll += int64(x)
		if y == 2 {
			type2Tree.Add(x, 1)
			countN2++
		}

		if countN2 == 0 {
			// B=2 が一人もいない場合、全員係数1
			out(sumAll)
		} else {
			// B=2 の馬全員の A の和
			// type2TreeにはB=2の馬しか入っていないので、TopSum(countN2)は全員の和になる
			sumType2 := type2Tree.AllProd().sum

			// 全体の上位 countN2 人の A の和
			sumTopN2 := allTree.QueryTopSum(countN2)

			if sumType2 == sumTopN2 {
				sumTopNext := allTree.QueryTopSum(countN2 + 1)
				kthVal := allTree.GetKthLargest(countN2)

				ans := sumAll + sumTopNext - int64(kthVal)
				out(ans)
			} else {
				ans := sumAll + sumTopN2
				out(ans)
			}
		}
	}
}
