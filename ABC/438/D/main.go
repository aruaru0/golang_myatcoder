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
	val int
}

// 宣言はこんな感じ
// seg := newSegtree[S](v, e, merger)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	b := getInts(N)
	c := getInts(N)

	for i := 0; i < N-1; i++ {
		a[i+1] += a[i]
		b[i+1] += b[i]

	}

	for i := N - 1; i > 0; i-- {
		c[i-1] += c[i]
	}

	b = append([]int{0}, b...)
	c = append(c, 0)
	v := make([]S, N+1)

	for i := 0; i < N+1; i++ {
		v[i] = S{b[i] + c[i]}
	}

	const inf = int(1e18)
	e := func() S {
		return S{-inf}
	}

	merger := func(a, b S) S {
		if a.val > b.val {
			return a
		}
		return b
	}

	seg := newSegtree[S](v, e, merger)
	ans := -inf

	for i := 0; i < N-2; i++ {
		pre := a[i] - b[i+1]
		post := seg.Prod(i+2, N).val

		ans = max(ans, pre+post)
	}

	out(ans)
}
