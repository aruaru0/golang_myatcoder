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
	c := getInts(N)

	type Data struct {
		sum  int
		size int
	}

	initialData := make([]Data, N)
	for i := 0; i < N; i++ {
		initialData[i] = Data{sum: c[i], size: 1}
	}

	e := func() Data { return Data{sum: 0, size: 0} }
	merger := func(a, b Data) Data {
		return Data{sum: a.sum + b.sum, size: a.size + b.size}
	}
	mapper := func(f int, x Data) Data {
		return Data{sum: x.sum + f*x.size, size: x.size}
	}
	comp := func(f, g int) int { return f + g }
	id := func() int { return 0 }

	seg := NewLazySegTreeFromSlice(initialData, merger, e, mapper, comp, id)

	for qi := 0; qi < Q; qi++ {
		t := getI()
		if t == 1 {
			l, r, v := getI()-1, getI(), getI()
			seg.ApplyRange(l, r, v)
		} else {
			l, r := getI()-1, getI()
			out(seg.Prod(l, r).sum)
		}
	}
}
func lazysegtree_countrZero(n uint32) int {
	return bits.TrailingZeros32(n)
}

func lazysegtree_bitCeil(n uint32) uint32 {
	x := uint32(1)
	for x < n {
		x *= 2
	}
	return x
}

type LazySegTree[S any, F any] struct {
	n           int
	size        int
	log         int
	d           []S
	lz          []F
	op          func(S, S) S
	e           func() S
	mapping     func(F, S) S
	composition func(F, F) F
	id          func() F
}

func NewLazySegTree[S any, F any](
	n int,
	op func(S, S) S,
	e func() S,
	mapping func(F, S) S,
	composition func(F, F) F,
	id func() F,
) *LazySegTree[S, F] {
	v := make([]S, n)
	for i := 0; i < n; i++ {
		v[i] = e()
	}
	return NewLazySegTreeFromSlice(v, op, e, mapping, composition, id)
}

func NewLazySegTreeFromSlice[S any, F any](
	v []S,
	op func(S, S) S,
	e func() S,
	mapping func(F, S) S,
	composition func(F, F) F,
	id func() F,
) *LazySegTree[S, F] {
	n := len(v)
	size := int(lazysegtree_bitCeil(uint32(n)))
	log := lazysegtree_countrZero(uint32(size))
	d := make([]S, 2*size)
	for i := range d {
		d[i] = e()
	}
	lz := make([]F, size)
	for i := range lz {
		lz[i] = id()
	}
	for i := 0; i < n; i++ {
		d[size+i] = v[i]
	}
	st := &LazySegTree[S, F]{
		n:           n,
		size:        size,
		log:         log,
		d:           d,
		lz:          lz,
		op:          op,
		e:           e,
		mapping:     mapping,
		composition: composition,
		id:          id,
	}
	for i := size - 1; i >= 1; i-- {
		st.update(i)
	}
	return st
}

func (st *LazySegTree[S, F]) update(k int) {
	st.d[k] = st.op(st.d[2*k], st.d[2*k+1])
}

func (st *LazySegTree[S, F]) allApply(k int, f F) {
	st.d[k] = st.mapping(f, st.d[k])
	if k < st.size {
		st.lz[k] = st.composition(f, st.lz[k])
	}
}

func (st *LazySegTree[S, F]) push(k int) {
	st.allApply(2*k, st.lz[k])
	st.allApply(2*k+1, st.lz[k])
	st.lz[k] = st.id()
}

func (st *LazySegTree[S, F]) Set(p int, x S) {
	if p < 0 || p >= st.n {
		panic("p is out of bounds")
	}
	p += st.size
	for i := st.log; i >= 1; i-- {
		st.push(p >> i)
	}
	st.d[p] = x
	for i := 1; i <= st.log; i++ {
		st.update(p >> i)
	}
}

func (st *LazySegTree[S, F]) Get(p int) S {
	if p < 0 || p >= st.n {
		panic("p is out of bounds")
	}
	p += st.size
	for i := st.log; i >= 1; i-- {
		st.push(p >> i)
	}
	return st.d[p]
}

func (st *LazySegTree[S, F]) Prod(l, r int) S {
	if l < 0 || r < l || r > st.n {
		panic("invalid range")
	}
	if l == r {
		return st.e()
	}
	l += st.size
	r += st.size

	for i := st.log; i >= 1; i-- {
		if ((l >> i) << i) != l {
			st.push(l >> i)
		}
		if ((r >> i) << i) != r {
			st.push((r - 1) >> i)
		}
	}

	sml := st.e()
	smr := st.e()
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

func (st *LazySegTree[S, F]) AllProd() S {
	return st.d[1]
}

func (st *LazySegTree[S, F]) Apply(p int, f F) {
	if p < 0 || p >= st.n {
		panic("p is out of bounds")
	}
	p += st.size
	for i := st.log; i >= 1; i-- {
		st.push(p >> i)
	}
	st.d[p] = st.mapping(f, st.d[p])
	for i := 1; i <= st.log; i++ {
		st.update(p >> i)
	}
}

func (st *LazySegTree[S, F]) ApplyRange(l, r int, f F) {
	if l < 0 || r < l || r > st.n {
		panic("invalid range")
	}
	if l == r {
		return
	}

	l += st.size
	r += st.size

	for i := st.log; i >= 1; i-- {
		if ((l >> i) << i) != l {
			st.push(l >> i)
		}
		if ((r >> i) << i) != r {
			st.push((r - 1) >> i)
		}
	}

	l2, r2 := l, r
	for l < r {
		if l&1 == 1 {
			st.allApply(l, f)
			l++
		}
		if r&1 == 1 {
			r--
			st.allApply(r, f)
		}
		l >>= 1
		r >>= 1
	}
	l, r = l2, r2

	for i := 1; i <= st.log; i++ {
		if ((l >> i) << i) != l {
			st.update(l >> i)
		}
		if ((r >> i) << i) != r {
			st.update((r - 1) >> i)
		}
	}
}

func (st *LazySegTree[S, F]) MaxRight(l int, g func(S) bool) int {
	if l < 0 || l > st.n {
		panic("l is out of bounds")
	}
	if !g(st.e()) {
		panic("g(e) must be true")
	}
	if l == st.n {
		return st.n
	}
	l += st.size
	for i := st.log; i >= 1; i-- {
		st.push(l >> i)
	}
	sm := st.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !g(st.op(sm, st.d[l])) {
			for l < st.size {
				st.push(l)
				l = 2 * l
				if g(st.op(sm, st.d[l])) {
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

func (st *LazySegTree[S, F]) MinLeft(r int, g func(S) bool) int {
	if r < 0 || r > st.n {
		panic("r is out of bounds")
	}
	if !g(st.e()) {
		panic("g(e) must be true")
	}
	if r == 0 {
		return 0
	}
	r += st.size
	for i := st.log; i >= 1; i-- {
		st.push((r - 1) >> i)
	}
	sm := st.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !g(st.op(st.d[r], sm)) {
			for r < st.size {
				st.push(r)
				r = 2*r + 1
				if g(st.op(st.d[r], sm)) {
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
