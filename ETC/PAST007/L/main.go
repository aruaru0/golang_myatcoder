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

type E func() int
type Merger func(a, b int) int
type Compare func(v int) bool
type Segtree struct {
	n      int
	size   int
	log    int
	d      []int
	e      E
	merger Merger
}

func newSegtree(v []int, e E, m Merger) *Segtree {
	seg := new(Segtree)
	seg.n = len(v)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]int, 2*seg.size)
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
func (seg *Segtree) Set(p, x int) {
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}
func (seg *Segtree) Get(p int) int {
	return seg.d[p+seg.size]
}
func (seg *Segtree) Prod(l, r int) int {
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
func (seg *Segtree) AllProd() int {
	return seg.d[1]
}
func (seg *Segtree) MaxRight(l int, cmp Compare) int {
	if l == seg.n {
		return seg.n
	}
	l += seg.size
	sm := seg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(seg.merger(sm, seg.d[l])) {
			for l < seg.size {
				l = 2 * l
				if cmp(seg.merger(sm, seg.d[l])) {
					sm = seg.merger(sm, seg.d[l])
					l++
				}
			}
			return l - seg.size
		}
		sm = seg.merger(sm, seg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return seg.n
}
func (seg *Segtree) MinLeft(r int, cmp Compare) int {
	if r == 0 {
		return 0
	}
	r += seg.size
	sm := seg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(seg.merger(seg.d[r], sm)) {
			for r < seg.size {
				r = 2*r + 1
				if cmp(seg.merger(seg.d[r], sm)) {
					sm = seg.merger(seg.d[r], sm)
					r--
				}
			}
			return r + 1 - seg.size
		}
		sm = seg.merger(seg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}
func (seg *Segtree) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()
	a := getInts(N)

	m := make(map[int][]int)
	for i := 0; i < N; i++ {
		m[a[i]] = append(m[a[i]], i)
	}
	e := func() int {
		return inf
	}
	merger := func(a, b int) int {
		return min(a, b)
	}
	seg := newSegtree(a, e, merger)

	for q := 0; q < Q; q++ {
		t, x, y := getI(), getI(), getI()
		if t == 1 {
			x--
			a[x] = y
			seg.Set(x, y)
			m[y] = append(m[y], x)
		} else {
			t := seg.Prod(x-1, y)
			// out(t, m[t])
			ans := make(map[int]bool, 0)
			newm := make([]int, 0)
			for _, e := range m[t] {
				if a[e] == t {
					newm = append(newm, e)
					if x-1 <= e && e < y {
						ans[e] = true
					}
				}
				m[t] = newm
			}
			ans2 := make([]int, 0)
			for e := range ans {
				ans2 = append(ans2, e)
			}
			sort.Ints(ans2)
			fmt.Fprint(wr, len(ans2), " ")
			for i := 0; i < len(ans2); i++ {
				fmt.Fprint(wr, ans2[i]+1, " ")
			}
			out()
		}
	}
}
