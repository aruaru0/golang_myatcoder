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

type pos struct {
	y      int
	x1, x2 int
	flg    bool
}

type S struct {
	x, l int
}
type F struct {
	a bool
}
type E func() S
type Merger func(a, b S) S
type Mapper func(f F, x S) S
type Comp func(f, g F) F
type Id func() F
type Compare func(v S) bool
type LazySegtree struct {
	n      int
	size   int
	log    int
	d      []S
	lz     []F
	e      E
	merger Merger
	mapper Mapper
	comp   Comp
	id     Id
}

func newLazySegtree(v []S, e E, merger Merger, mapper Mapper, comp Comp, id Id) *LazySegtree {
	lseg := new(LazySegtree)
	lseg.n = len(v)
	lseg.log = lseg.ceilPow2(lseg.n)
	lseg.size = 1 << uint(lseg.log)
	lseg.d = make([]S, 2*lseg.size)
	lseg.e = e
	lseg.lz = make([]F, lseg.size)
	lseg.merger = merger
	lseg.mapper = mapper
	lseg.comp = comp
	lseg.id = id
	for i, _ := range lseg.d {
		lseg.d[i] = lseg.e()
	}
	for i, _ := range lseg.lz {
		lseg.lz[i] = lseg.id()
	}
	for i := 0; i < lseg.n; i++ {
		lseg.d[lseg.size+i] = v[i]
	}
	for i := lseg.size - 1; i >= 1; i-- {
		lseg.Update(i)
	}
	return lseg
}
func (lseg *LazySegtree) Update(k int) {
	lseg.d[k] = lseg.merger(lseg.d[2*k], lseg.d[2*k+1])
}
func (lseg *LazySegtree) AllApply(k int, f F) {
	lseg.d[k] = lseg.mapper(f, lseg.d[k])
	if k < lseg.size {
		lseg.lz[k] = lseg.comp(f, lseg.lz[k])
	}
}
func (lseg *LazySegtree) Push(k int) {
	lseg.AllApply(2*k, lseg.lz[k])
	lseg.AllApply(2*k+1, lseg.lz[k])
	lseg.lz[k] = lseg.id()
}
func (lseg *LazySegtree) Set(p int, x S) {
	p += lseg.size
	for i := lseg.log; i <= 1; i-- {
		lseg.Push(p >> uint(i))
	}
	lseg.d[p] = x
	for i := 1; i <= lseg.log; i++ {
		lseg.Update(p >> uint(i))
	}
}
func (lseg *LazySegtree) Get(p int) S {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.Push(p >> uint(i))
	}
	return lseg.d[p]
}
func (lseg *LazySegtree) Prod(l, r int) S {
	if l == r {
		return lseg.e()
	}
	l += lseg.size
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		if (l>>uint(i))<<uint(i) != l {
			lseg.Push(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg.Push(r >> uint(i))
		}
	}
	sml, smr := lseg.e(), lseg.e()
	for l < r {
		if (l & 1) == 1 {
			sml = lseg.merger(sml, lseg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = lseg.merger(lseg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return lseg.merger(sml, smr)
}
func (lseg *LazySegtree) AllProd() S {
	return lseg.d[1]
}
func (lseg *LazySegtree) Apply(p int, f F) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.Push(p >> uint(i))
	}
	lseg.d[p] = lseg.mapper(f, lseg.d[p])
	for i := 1; i <= lseg.log; i++ {
		lseg.Update(p >> uint(i))
	}
}
func (lseg *LazySegtree) RangeApply(l int, r int, f F) {
	if l == r {
		return
	}
	l += lseg.size
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		if (l>>uint(i))<<uint(i) != l {
			lseg.Push(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg.Push((r - 1) >> uint(i))
		}
	}
	l2, r2 := l, r
	for l < r {
		if l&1 == 1 {
			lseg.AllApply(l, f)
			l++
		}
		if r&1 == 1 {
			r--
			lseg.AllApply(r, f)
		}
		l >>= 1
		r >>= 1
	}
	l, r = l2, r2
	for i := 1; i <= lseg.log; i++ {
		if (l>>uint(i))<<uint(i) != l {
			lseg.Update(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg.Update((r - 1) >> uint(i))
		}
	}
}
func (lseg *LazySegtree) MaxRight(l int, cmp Compare) int {
	if l == lseg.n {
		return lseg.n
	}
	l += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.Push(l >> uint(i))
	}
	sm := lseg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(lseg.merger(sm, lseg.d[l])) {
			for l < lseg.size {
				lseg.Push(l)
				l = 2 * l
				if cmp(lseg.merger(sm, lseg.d[l])) {
					sm = lseg.merger(sm, lseg.d[l])
					l++
				}
			}
			return l - lseg.size
		}
		sm = lseg.merger(sm, lseg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return lseg.n
}
func (lseg *LazySegtree) MinLeft(r int, cmp Compare) int {
	if r == 0 {
		return 0
	}
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.Push(r - 1>>uint(i))
	}
	sm := lseg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(lseg.merger(lseg.d[r], sm)) {
			for r < lseg.size {
				lseg.Push(r)
				r = 2*r + 1
				if cmp(lseg.merger(lseg.d[r], sm)) {
					sm = lseg.merger(lseg.d[r], sm)
					r--
				}
			}
			return r + 1 - lseg.size
		}
		sm = lseg.merger(lseg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}
func (lseg *LazySegtree) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	Q := getI()

	m := make(map[int]int)
	y := make([]pos, 0)
	for i := 0; i < Q; i++ {
		a, b, c, d := getI(), getI(), getI(), getI()
		y = append(y, pos{a, b, d, true})
		y = append(y, pos{c, b, d, false})
		m[b] = 1
		m[d] = 1
	}
	sort.Slice(y, func(i, j int) bool {
		return y[i].y < y[j].y
	})

	x := make([]int, 0)
	for e := range m {
		x = append(x, e)
	}
	sort.Ints(x)
	for i := 0; i < len(x); i++ {
		m[x[i]] = i
	}
	d := make([]int, len(x))
	for i := 1; i < len(x); i++ {
		d[i-1] = x[i] - x[i-1]
	}

	e := func() S {
		return S{0, -1}
	}
	merger := func(a, b S) S {
		return S{a.x + b.x, a.l + b.l}
	}
	mapper := func(f F, x S) S {
		if f.a == false {
			return x
		}
		return S{x.l - x.x, x.l}
	}
	comp := func(f, g F) F {
		if f.a == false {
			return g
		}
		return F{f.a != g.a}
	}
	id := func() F {
		return F{false}
	}

	v := make([]S, len(x))
	for i := 0; i < len(x); i++ {
		v[i].l = d[i]
	}
	seg := newLazySegtree(v, e, merger, mapper, comp, id)
	ans := 0
	for i, j := 0, 0; i < len(y); i = j {
		l := 0
		if i != 0 {
			l = y[i].y - y[i-1].y
		}
		ans += seg.AllProd().x * l
		for j < len(y) && y[i].y == y[j].y {
			yl := y[j].x1
			yr := y[j].x2
			l := lowerBound(x, yl)
			r := lowerBound(x, yr)
			seg.RangeApply(l, r, F{true})
			j++
		}
	}
	out(ans)

}
