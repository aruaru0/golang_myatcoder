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
	val int
	cnt int
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()

	a_max := 0
	a := make([]int, N)
	m := make(map[int]bool)
	for i := 0; i < N; i++ {
		a[i] = getI()
		a_max = max(a_max, a[i])
		m[a[i]] = true
	}

	for i := 0; i <= N; i++ {
		m[i] = true
	}

	idx := make([]int, Q)
	x := make([]int, Q)
	for q := 0; q < Q; q++ {
		idx[q], x[q] = getI()-1, getI()
		m[x[q]] = true
	}

	idx2val := make([]int, 0)
	for e := range m {
		idx2val = append(idx2val, e)
	}
	sort.Ints(idx2val)

	val2idx := make(map[int]int)
	for i, e := range idx2val {
		val2idx[e] = i
	}

	max_idx := len(idx2val)
	v := make([]S, max_idx)
	for i := 0; i < max_idx; i++ {
		v[i].val = i
	}
	for i := 0; i < N; i++ {
		idx := val2idx[a[i]]
		v[idx].cnt++
	}

	const inf = int(1e6)
	e := func() S { return S{-1, inf} }
	merger := func(a, b S) S {
		if a.cnt == 0 && b.cnt == 0 {
			if a.val < b.val {
				return a
			} else {
				return b
			}
		}
		if a.cnt != 0 {
			return b
		}
		return a
	}

	seg := newSegtree(v, e, merger)

	for i := 0; i < Q; i++ {
		prev := a[idx[i]]
		a[idx[i]] = x[i]
		ppos, npos := val2idx[prev], val2idx[x[i]]
		e := seg.Get(ppos)
		e.cnt--
		seg.Set(ppos, e)
		e = seg.Get(npos)
		e.cnt++
		seg.Set(npos, e)

		e = seg.AllProd()

		out(e.val)
	}
}
