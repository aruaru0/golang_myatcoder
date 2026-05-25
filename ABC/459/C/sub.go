package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()

	const inf = int(1e18)
	op := func(a, b int) int {
		return min(a, b)
	}
	e := func() int {
		return inf
	}

	const size = int(1e6)

	seg := NewSegTreeFromSlice(make([]int, N), op, e)
	bit := NewFenwickTree[int](size)
	bit.Add(0, N)

	// for i := 0; i <= N; i++ {
	// 	fmt.Fprint(wr, bit.Sum(i, i+1), " ")
	// }
	// out()

	for qi := 0; qi < Q; qi++ {
		t := getI()
		if t == 1 {
			x := getI() - 1
			v := seg.Get(x)
			seg.Set(x, v+1)
			bit.Add(v, -1)
			bit.Add(v+1, 1)
			// out("pos", x, "v", v)
		} else {
			y := getI()
			e := seg.AllProd()
			cnt := bit.Sum(e+y, size)
			// out("show", y, e, cnt)
			out(cnt)
		}
		// out("---------")
		// for i := 0; i < N; i++ {
		// 	fmt.Fprint(wr, seg.Get(i), " ")
		// }
		// out()
		// for i := 0; i <= N; i++ {
		// 	fmt.Fprint(wr, bit.Sum(i, i+1), " ")
		// }
		// out()
	}
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type FenwickTree[T Numeric] struct {
	n    int
	data []T
}

func NewFenwickTree[T Numeric](n int) *FenwickTree[T] {
	if n < 0 {
		panic("n must be non-negative")
	}
	return &FenwickTree[T]{
		n:    n,
		data: make([]T, n),
	}
}

func (f *FenwickTree[T]) Add(p int, x T) {
	if p < 0 || p >= f.n {
		panic("p is out of bounds")
	}
	p++
	for p <= f.n {
		f.data[p-1] += x
		p += p & -p
	}
}

func (f *FenwickTree[T]) Sum(l, r int) T {
	if l < 0 || l > r || r > f.n {
		panic("invalid range")
	}
	return f.sum(r) - f.sum(l)
}

func (f *FenwickTree[T]) sum(r int) T {
	var s T
	for r > 0 {
		s += f.data[r-1]
		r -= r & -r
	}
	return s
}
func segtree_countrZero(n uint32) int {
	return bits.TrailingZeros32(n)
}

func segtree_bitCeil(n uint32) uint32 {
	x := uint32(1)
	for x < n {
		x *= 2
	}
	return x
}

type SegTree[S any] struct {
	n    int
	size int
	log  int
	d    []S
	op   func(S, S) S
	e    func() S
}

func NewSegTree[S any](n int, op func(S, S) S, e func() S) *SegTree[S] {
	v := make([]S, n)
	for i := 0; i < n; i++ {
		v[i] = e()
	}
	return NewSegTreeFromSlice(v, op, e)
}

func NewSegTreeFromSlice[S any](v []S, op func(S, S) S, e func() S) *SegTree[S] {
	n := len(v)
	size := int(segtree_bitCeil(uint32(n)))
	log := segtree_countrZero(uint32(size))
	d := make([]S, 2*size)
	for i := range d {
		d[i] = e()
	}
	for i := 0; i < n; i++ {
		d[size+i] = v[i]
	}
	st := &SegTree[S]{
		n:    n,
		size: size,
		log:  log,
		d:    d,
		op:   op,
		e:    e,
	}
	for i := size - 1; i >= 1; i-- {
		st.update(i)
	}
	return st
}

func (st *SegTree[S]) update(k int) {
	st.d[k] = st.op(st.d[2*k], st.d[2*k+1])
}

func (st *SegTree[S]) Set(p int, x S) {
	if p < 0 || p >= st.n {
		panic("p is out of bounds")
	}
	p += st.size
	st.d[p] = x
	for i := 1; i <= st.log; i++ {
		st.update(p >> i)
	}
}

func (st *SegTree[S]) Get(p int) S {
	if p < 0 || p >= st.n {
		panic("p is out of bounds")
	}
	return st.d[p+st.size]
}

func (st *SegTree[S]) Prod(l, r int) S {
	if l < 0 || r < l || r > st.n {
		panic("invalid range")
	}
	sml := st.e()
	smr := st.e()
	l += st.size
	r += st.size

	for l < r {
		if l&1 == 1 {
			sml = st.op(sml, st.d[l])
			l++
		}
		if r&1 == 1 {
			r--
			smr = st.op(st.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return st.op(sml, smr)
}

func (st *SegTree[S]) AllProd() S {
	return st.d[1]
}

func (st *SegTree[S]) MaxRight(l int, f func(S) bool) int {
	if l < 0 || l > st.n {
		panic("l is out of bounds")
	}
	if !f(st.e()) {
		panic("f(e) must be true")
	}
	if l == st.n {
		return st.n
	}
	l += st.size
	sm := st.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !f(st.op(sm, st.d[l])) {
			for l < st.size {
				l = 2 * l
				if f(st.op(sm, st.d[l])) {
					sm = st.op(sm, st.d[l])
					l++
				}
			}
			return l - st.size
		}
		sm = st.op(sm, st.d[l])
		l++
		if (l & -l) == l {
			break
		}
	}
	return st.n
}

func (st *SegTree[S]) MinLeft(r int, f func(S) bool) int {
	if r < 0 || r > st.n {
		panic("r is out of bounds")
	}
	if !f(st.e()) {
		panic("f(e) must be true")
	}
	if r == 0 {
		return 0
	}
	r += st.size
	sm := st.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !f(st.op(st.d[r], sm)) {
			for r < st.size {
				r = 2*r + 1
				if f(st.op(st.d[r], sm)) {
					sm = st.op(st.d[r], sm)
					r--
				}
			}
			return r + 1 - st.size
		}
		sm = st.op(st.d[r], sm)
		if (r & -r) == r {
			break
		}
	}
	return 0
}
