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

// Sはdataの構造体
// sumの場合weightが必要
type S struct {
	value int
	// weight int
}

// Fはlazyの構造体
// updateの場合flagが必要
type F struct {
	value int
	// flag  bool
}

type SegTree[S, F comparable] struct {
	length       int
	binaryLength int
	data         []S
	lazy         []F
	e            S
	id           F
	op           func(a, b S) S
	mapping      func(f F, x S) S
	composition  func(f, g F) F
}

func NewSegTree[S, F comparable](leaves []S, e S, id F, op func(a, b S) S, mapping func(f F, x S) S, composition func(f, g F) F) SegTree[S, F] {
	binaryLength := 1
	for binaryLength < len(leaves) {
		binaryLength *= 2
	}
	st := SegTree[S, F]{length: len(leaves), binaryLength: binaryLength, data: make([]S, binaryLength*2-1), lazy: make([]F, binaryLength*2-1), e: e, id: id, op: op, mapping: mapping, composition: composition}
	for i := 0; i < len(st.data)-(binaryLength-1); i++ {
		if i < len(leaves) {
			st.data[i+binaryLength-1] = leaves[i]
		} else {
			st.data[i+binaryLength-1] = e
		}
	}
	st.build()
	for i := range st.lazy {
		st.lazy[i] = st.id
	}
	return st
}

func (st *SegTree[S, F]) Set(l, r int, value F) {
	st.setSegment(l, r, 0, 0, st.binaryLength, value)
}

func (st *SegTree[S, F]) Get(l, r int) S {
	return st.getSegment(l, r, 0, 0, st.binaryLength)
}

func (st *SegTree[S, F]) MaxRight(l int, isOk func(S) bool) int {
	if l >= st.length-1 {
		return st.length - 1
	}
	if l < 0 {
		l = 0
	}
	l += st.binaryLength - 1
	stack := make([]int, 0)
	for i := l; i >= 0; i = (i - 1) >> 1 {
		stack = append(stack, i)
	}
	for i := range stack {
		st.eval(stack[len(stack)-1-i])
	}
	now := st.e
	for {
		for l%2 == 1 {
			l = (l - 1) >> 1
		}
		st.eval(l)
		if !isOk(st.op(now, st.data[l])) {
			for l < st.binaryLength-1 {
				l = l<<1 + 1
				st.eval(l)
				if isOk(st.op(now, st.data[l])) {
					now = st.op(now, st.data[l])
					l++
					st.eval(l)
				}
			}
			return l - st.binaryLength
		}
		now = st.op(now, st.data[l])
		l++
		if l&(l+1) == 0 {
			break
		}
	}
	return st.length - 1
}

func (st *SegTree[S, F]) MinLeft(r int, isOk func(S) bool) int {
	if r <= 0 {
		return 0
	}
	if r >= st.length {
		r = st.length - 1
	}
	r += st.binaryLength - 1
	// r += 1
	now := st.e
	for {
		r--
		for r > 0 && r%2 == 0 {
			r = (r - 1) >> 1
		}
		st.eval(r)
		if !isOk(st.op(now, st.data[r])) {
			for r < st.binaryLength-1 {
				r = r<<1 + 2
				st.eval(r)
				if isOk(st.op(now, st.data[r])) {
					now = st.op(now, st.data[r])
					r--
				}
			}
			return r + 1 - st.binaryLength
		}
		now = st.op(now, st.data[r])
		if r&(r+1) == 0 {
			break
		}
	}
	return 0
}

func (st *SegTree[S, F]) build() {
	for i := st.binaryLength - 2; i >= 0; i-- {
		st.data[i] = st.op(st.data[1+i<<1], st.data[2+i<<1])
	}
}

func (st *SegTree[S, F]) eval(idx int) {
	if st.lazy[idx] == st.id {
		return
	}
	if idx < st.binaryLength-1 {
		st.lazy[1+idx<<1] = st.composition(st.lazy[idx], st.lazy[1+idx<<1])
		st.lazy[2+idx<<1] = st.composition(st.lazy[idx], st.lazy[2+idx<<1])
	}
	st.data[idx] = st.mapping(st.lazy[idx], st.data[idx])
	st.lazy[idx] = st.id
}

func (st *SegTree[S, F]) setSegment(l, r, idx, nowL, nowR int, value F) {
	st.eval(idx)
	if nowL >= r || nowR <= l {
		return
	} else if nowL >= l && nowR <= r {
		st.lazy[idx] = st.composition(value, st.lazy[idx])
		st.eval(idx)
	} else if nowL < r && nowR > l {
		st.setSegment(l, r, 1+idx<<1, nowL, (nowL+nowR)/2, value)
		st.setSegment(l, r, 2+idx<<1, (nowL+nowR)/2, nowR, value)
		st.data[idx] = st.op(st.data[1+idx<<1], st.data[2+idx<<1])
	}
}

func (st *SegTree[S, F]) getSegment(l, r, idx, nowL, nowR int) S {
	st.eval(idx)
	if nowL >= r || nowR <= l {
		return st.e
	} else if nowL >= l && nowR <= r {
		return st.data[idx]
	} else {
		v1 := st.getSegment(l, r, 1+idx<<1, nowL, (nowL+nowR)/2)
		v2 := st.getSegment(l, r, 2+idx<<1, (nowL+nowR)/2, nowR)
		return st.op(v1, v2)
	}
}

type Pair struct {
	x, y int
}

// Coodinate Compression 構造体と遅延評価用の関数群を追加
type CC struct {
	initialized bool
	xs          []Pair
}

func (cc *CC) Add(x Pair) {
	cc.xs = append(cc.xs, x)
}

func (cc *CC) Init() {
	sort.Slice(cc.xs, func(i, j int) bool {
		if cc.xs[i].x != cc.xs[j].x {
			return cc.xs[i].x < cc.xs[j].x
		}
		return cc.xs[i].y < cc.xs[j].y
	})
	if len(cc.xs) == 0 {
		cc.initialized = true
		return
	}
	res := make([]Pair, 0, len(cc.xs))
	res = append(res, cc.xs[0])
	for i := 1; i < len(cc.xs); i++ {
		if cc.xs[i] != res[len(res)-1] {
			res = append(res, cc.xs[i])
		}
	}
	cc.xs = res
	cc.initialized = true
}

func (cc *CC) Get(x Pair) int {
	if !cc.initialized {
		cc.Init()
	}
	idx := sort.Search(len(cc.xs), func(i int) bool {
		if cc.xs[i].x != x.x {
			return cc.xs[i].x >= x.x
		}
		return cc.xs[i].y >= x.y
	})
	return idx
}

func (cc *CC) Size() int {
	if !cc.initialized {
		cc.Init()
	}
	return len(cc.xs)
}

const INF = 1000000000000000000

func opVal(a, b int) int {
	return max(a, b)
}

func mappingVal(f, x int) int {
	return f + x
}

func compositionVal(f, g int) int {
	return f + g
}

type QueryInfo struct {
	typ, i, x int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()
	a := getInts(N)
	b := getInts(N)

	p := make([]Pair, 0)

	bid := make([]int, N)
	cc := &CC{xs: p}
	for i := 0; i < N; i++ {
		bid[i] = i
		cc.Add(Pair{b[i], i})
	}

	qs := make([]QueryInfo, Q)
	for qi := 0; qi < Q; qi++ {
		typ := getI()
		i := getI() - 1
		x := getI()
		qs[qi] = QueryInfo{typ, i, x}
		if typ == 2 {
			cc.Add(Pair{x, N + qi})
		}
	}

	m := cc.Size()
	leaves := make([]int, m)
	for i := 0; i < m; i++ {
		leaves[i] = -INF
	}

	st := NewSegTree[int, int](leaves, -INF, 0, opVal, mappingVal, compositionVal)

	add := func(i int) {
		j := cc.Get(Pair{b[i], bid[i]})
		valJ := st.Get(j, j+1)
		d := valJ + INF
		newVal := b[i] + d
		// 1点上書きの代わりに、目標値と現在値の差分を加算することで代用
		st.Set(j, j+1, newVal-valJ)
		st.Set(0, j+1, a[i])
	}

	del := func(i int) {
		j := cc.Get(Pair{b[i], bid[i]})
		st.Set(0, j+1, -a[i])
		valJ := st.Get(j, j+1)
		d := valJ - b[i]
		newVal := d - INF
		st.Set(j, j+1, newVal-valJ)
	}

	for i := 0; i < N; i++ {
		add(i)
	}

	for qi := 0; qi < Q; qi++ {
		typ := qs[qi].typ
		i := qs[qi].i
		x := qs[qi].x
		del(i)
		if typ == 1 {
			a[i] = x
		} else {
			b[i] = x
			bid[i] = N + qi
		}
		add(i)
		out(st.Get(0, m))
	}
}
