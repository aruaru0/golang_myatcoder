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

type S struct {
	l, r, d int
}
type E func() S
type Merger func(a, b S) S
type Compare func(v int) bool
type Segtree struct {
	n      int
	size   int
	log    int
	d      []S
	e      E
	merger Merger
}

func newSegtree(v []S, e E, m Merger) *Segtree {
	seg := new(Segtree)
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
func (seg *Segtree) Update(k int) {
	seg.d[k] = seg.merger(seg.d[2*k], seg.d[2*k+1])
}
func (seg *Segtree) Set(p int, x S) {
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}
func (seg *Segtree) Get(p int) S {
	return seg.d[p+seg.size]
}
func (seg *Segtree) Prod(l, r int) S {
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
func (seg *Segtree) AllProd() S {
	return seg.d[1]
}

func (seg *Segtree) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

func (p S) f(s int) (int, int) {
	if p.d == -1 {
		if s < p.l {
			return p.l, p.l - s
		} else if s > p.r {
			return p.r, s - p.r
		}
		return s, 0
	} else {
		return p.r, p.d + abs(s-p.l)
	}
}

const INF = 1001001001

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()

	e := func() S {
		return S{0, INF, -1}
	}
	merger := func(a, b S) S {
		if a.d == -1 {
			if b.d == -1 {
				if a.r < b.l {
					return S{a.r, b.l, b.l - a.r}
				}
				if a.l > b.r {
					return S{a.l, b.r, a.l - b.r}
				}
				return S{max(a.l, b.l), min(a.r, b.r), -1}
			} else {
				t, d := a.f(b.l)
				return S{t, b.r, b.d + d}
			}
		} else {
			t, d := b.f(a.r)
			return S{a.l, t, a.d + d}
		}
	}

	v := make([]S, n)
	for i := 0; i < n; i++ {
		v[i] = S{0, INF, -1}
	}
	t := newSegtree(v, e, merger)
	for i := 0; i < n; i++ {
		l, u := getI(), getI()
		t.Set(i, S{l, u, -1})
	}

	q := getI()
	for qi := 0; qi < q; qi++ {
		sx, sy, tx, ty := getI(), getI(), getI(), getI()
		if sx > tx {
			sx, tx = tx, sx
			sy, ty = ty, sy
		}
		sx--
		tx--
		y, d := t.Prod(sx, tx).f(sy)
		d += abs(y - ty)
		d += tx - sx
		out(d)
	}
}
