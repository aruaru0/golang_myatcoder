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
	for i, _ := range seg.d {
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

// Sは書き換える
type S struct {
	ans    int
	lc, rc byte
	l, r   int
	same   bool
}

// 宣言はこんな感じ
// seg := newSegtree[S](v, e, merger)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	e := func() S {
		return S{
			ans:  0,
			lc:   '?',
			rc:   '?',
			l:    0,
			r:    0,
			same: false,
		}
	}
	merger := func(a, b S) S {

		var res S
		// 答えはどちらかの大きい方
		res.ans = max(a.ans, b.ans)
		// もし右端と左端が同じなら、連結して答えを更新
		if a.rc == b.lc {
			res.ans = max(res.ans, a.r+b.l)
		}
		// それぞれ代入
		res.lc = a.lc
		res.l = a.l
		res.rc = b.rc
		res.r = b.r
		// aが全て同じでbの左端と同じなら左を伸ばす
		if a.same && a.lc == b.lc {
			res.l += b.l
		}
		// bが全て一致でaの右端が同じなら右を伸ばす
		if b.same && a.rc == b.rc {
			res.r += a.r
		}
		/// 全て一致しているならsameをtrueにする
		res.same = a.same && b.same && a.lc == b.lc
		return res
	}

	n, q := getI(), getI()
	s := getS()

	makeS := func(c byte) S {
		// １文字のセグメント
		return S{
			ans:  1, // １文字
			lc:   c, // 左端と右端は同じ
			rc:   c,
			l:    1, // 左からも右からも１文字
			r:    1,
			same: true,
		}
	}

	v := make([]S, n)
	for i := 0; i < n; i++ {
		v[i] = makeS(s[i])
	}

	seg := newSegtree[S](v, e, merger)

	for qi := 0; qi < q; qi++ {
		t := getI()
		if t == 1 {
			i, x := getI()-1, getS()[0]
			seg.Set(i, makeS(byte(x)))
		}
		if t == 2 {
			l, r := getI()-1, getI()
			ret := seg.Prod(l, r)
			out(ret.ans)
		}
	}
}
