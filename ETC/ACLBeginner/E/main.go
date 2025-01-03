package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

const mod = 998244353

func f(s string) {
	x := 0
	for _, e := range s {
		x = x*10 + int(e-'0')
		x %= mod
	}
	out(x)
}

func mpow(a, b int) int {
	if b == 0 {
		return 1
	} else if b%2 == 0 {
		tmp := mpow(a, b/2)
		return tmp * tmp % mod
	}
	return mpow(a, b-1) * a % mod
}

// 逆元を使った割り算（MOD）
// mod. m での a の逆元 a^{-1} を計算する
func modinv(a, m int) int {
	b := m
	u := 1
	v := 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= m
	if u < 0 {
		u += m
	}
	return u
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	inv9 := modinv(9, mod)
	// 初期値
	AffineE := func() S {
		return S{0, 1}
	}
	// Ｓ同士
	AffileMerger := func(l, r S) S {
		a := l.a * r.w % mod
		a += r.a
		a %= mod
		w := l.w * r.w
		w %= mod
		return S{a, w}
	}
	// 変換処理
	AffineMapper := func(l F, r S) S {
		if l.d == 0 {
			return r
		}
		a := (r.w - 1)
		if a < 0 {
			a += mod
		}
		a *= inv9
		a %= mod
		a *= l.d
		a %= mod
		return S{a, r.w}
	}
	//　FoG処理
	AffineComp := func(l F, r F) F {
		if l.d == 0 {
			return r
		}
		return l
	}
	// 変換しても変化しない係数
	AffineID := func() F {
		return F{0}
	}

	N, Q := getInt(), getInt()
	a := make([]S, N)
	for i := 0; i < N; i++ {
		a[i] = S{1, 10}
	}
	seg := newLazySegtree(a, AffineE, AffileMerger, AffineMapper, AffineComp, AffineID)
	for i := 0; i < Q; i++ {
		l, r, d := getInt()-1, getInt(), getInt()
		seg.RangeApply(l, r, F{d})
		out(seg.AllProd().a)
	}
}

type S struct {
	a, w int
}

type F struct {
	d int
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
