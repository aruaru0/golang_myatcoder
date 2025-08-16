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

const (
	INF int = 1e18
	mod     = 998244353
	eps     = 1e-12
)

func modAdd(a, b int64) int64 {
	x := a + b
	if x >= mod {
		x -= mod
	}
	return x
}
func modSub(a, b int64) int64 {
	x := a - b
	if x < 0 {
		x += mod
	}
	return x
}

func modMul(a, b int64) int64 { return (a * b) % mod }

func modPow(a, e int64) int64 {
	res := int64(1)
	for e > 0 {
		if e&1 == 1 {
			res = modMul(res, a)
		}
		a = modMul(a, a)
		e >>= 1
	}
	return res
}

func modInv(a int64) int64 { return modPow(a, mod-2) }

type Comb struct {
	fac    []int64
	facInv []int64
}

func initComb(n int) *Comb {
	f := make([]int64, n+1)
	g := make([]int64, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = modMul(f[i-1], int64(i))
	}
	g[n] = modInv(f[n])
	for i := n; i > 0; i-- {
		g[i-1] = modMul(g[i], int64(i))
	}
	return &Comb{fac: f, facInv: g}
}

func (c *Comb) C(n, k int64) int64 {
	if k < 0 || n < 0 || k > n {
		return 0
	}
	return modMul(modMul(c.fac[n], c.facInv[k]), c.facInv[n-k])
}

func initFib(n int) []int64 {
	F := make([]int64, n+1)
	if n >= 0 {
		F[0] = 0
	}
	if n >= 1 {
		F[1] = 1
	}
	for i := 2; i <= n; i++ {
		F[i] = modAdd(F[i-1], F[i-2])
	}
	return F
}

type Mat struct {
	a00, a01, a10, a11 int64
}

var Iden = Mat{1, 0, 0, 1}

func mul(A, B Mat) Mat {
	return Mat{
		a00: (A.a00*B.a00 + A.a01*B.a10) % mod,
		a01: (A.a00*B.a01 + A.a01*B.a11) % mod,
		a10: (A.a10*B.a00 + A.a11*B.a10) % mod,
		a11: (A.a10*B.a01 + A.a11*B.a11) % mod,
	}
}

func buildEdgeMat(L, r int64, C *Comb) Mat {
	return Mat{
		a00: C.C(L-r, r),
		a01: C.C(L-r, r-1),
		a10: C.C(L-r-1, r),
		a11: C.C(L-r-1, r-1),
	}
}

type SegTree struct {
	n    int
	data []Mat
}

func NewSegTree(n int) *SegTree {
	st := &SegTree{
		n:    n,
		data: make([]Mat, 4*n+5),
	}
	st.build(1, 1, n)
	return st
}

func (st *SegTree) build(idx, l, r int) {
	if l == r {
		st.data[idx] = Iden
		return
	}
	m := (l + r) >> 1
	st.build(idx<<1, l, m)
	st.build(idx<<1|1, m+1, r)
	st.data[idx] = mul(st.data[idx<<1], st.data[idx<<1|1])
}

func (st *SegTree) update(pos int, val Mat) {
	st._update(1, 1, st.n, pos, val)
}
func (st *SegTree) _update(idx, l, r, pos int, val Mat) {
	if l == r {
		st.data[idx] = val
		return
	}
	m := (l + r) >> 1
	if pos <= m {
		st._update(idx<<1, l, m, pos, val)
	} else {
		st._update(idx<<1|1, m+1, r, pos, val)
	}
	st.data[idx] = mul(st.data[idx<<1], st.data[idx<<1|1])
}

func (st *SegTree) AllProduct() Mat { return st.data[1] }

type node struct {
	key  int
	pri  uint64
	l, r *node
}

type TreapSet struct {
	root *node
	seed uint64
}

func NewTreapSet() *TreapSet { return &TreapSet{seed: 88172645463393265} }

func (t *TreapSet) rnd() uint64 {
	t.seed ^= t.seed << 7
	t.seed ^= t.seed >> 9
	return t.seed
}

func rotateRight(p *node) *node {
	q := p.l
	p.l = q.r
	q.r = p
	return q
}
func rotateLeft(p *node) *node {
	q := p.r
	p.r = q.l
	q.l = p
	return q
}

func (t *TreapSet) insert(p **node, key int) {
	if *p == nil {
		*p = &node{key: key, pri: t.rnd()}
		return
	}
	if key < (*p).key {
		t.insert(&((*p).l), key)
		if (*p).l.pri > (*p).pri {
			*p = rotateRight(*p)
		}
	} else if key > (*p).key {
		t.insert(&((*p).r), key)
		if (*p).r.pri > (*p).pri {
			*p = rotateLeft(*p)
		}
	}
}

func (t *TreapSet) delete(p **node, key int) {
	if *p == nil {
		return
	}
	if key < (*p).key {
		t.delete(&((*p).l), key)
	} else if key > (*p).key {
		t.delete(&((*p).r), key)
	} else {
		if (*p).l == nil {
			*p = (*p).r
		} else if (*p).r == nil {
			*p = (*p).l
		} else {
			if (*p).l.pri > (*p).r.pri {
				*p = rotateRight(*p)
				t.delete(&((*p).r), key)
			} else {
				*p = rotateLeft(*p)
				t.delete(&((*p).l), key)
			}
		}
	}
}

func (t *TreapSet) Insert(key int) { t.insert(&t.root, key) }
func (t *TreapSet) Remove(key int) { t.delete(&t.root, key) }
func (t *TreapSet) Has(key int) bool {
	cur := t.root
	for cur != nil {
		if key < cur.key {
			cur = cur.l
		} else if key > cur.key {
			cur = cur.r
		} else {
			return true
		}
	}
	return false
}
func (t *TreapSet) Prev(key int) int {
	ans := 0
	cur := t.root
	for cur != nil {
		if cur.key < key {
			ans = cur.key
			cur = cur.r
		} else {
			cur = cur.l
		}
	}
	return ans
}
func (t *TreapSet) Next(key int) int {
	ans := 0
	cur := t.root
	for cur != nil {
		if cur.key > key {
			ans = cur.key
			cur = cur.l
		} else {
			cur = cur.r
		}
	}
	return ans
}
func (t *TreapSet) Max() int {
	if t.root == nil {
		return 0
	}
	cur := t.root
	for cur.r != nil {
		cur = cur.r
	}
	return cur.key
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, q := getI(), getI()
	comb := initComb(n + 5)
	fib := initFib(n + 5)
	seg := NewSegTree(n)

	a := make([]int64, n+1)
	st := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		a[i] = -1
	}

	set := NewTreapSet()

	getB := func(idx int) int64 {
		if idx == 0 {
			return 0
		}
		return a[idx]
	}

	updateEdgeAt := func(idx int) {
		if idx <= 0 || idx > n {
			return
		}
		if !st[idx] {
			seg.update(idx, Iden)
			return
		}
		pre := set.Prev(idx)
		L := int64(idx - pre)
		r := getB(idx) - getB(pre)
		mat := buildEdgeMat(L, r, comb)
		seg.update(idx, mat)
	}

	for i := 0; i < q; i++ {
		x, y := getI(), getI()
		if y == -1 {
			if st[x] {
				pre := set.Prev(x)
				suc := set.Next(x)
				set.Remove(x)
				st[x] = false
				a[x] = -1
				seg.update(x, Iden)

				if suc != 0 {
					L := int64(suc - pre)
					r := getB(suc) - getB(pre)
					seg.update(suc, buildEdgeMat(L, r, comb))
				}
			}
		} else {
			yy := int64(y)
			if st[x] {
				a[x] = yy
				updateEdgeAt(x)
				suc := set.Next(x)
				if suc != 0 {
					updateEdgeAt(suc)
				}
			} else {
				pre := set.Prev(x)
				suc := set.Next(x)
				set.Insert(x)
				st[x] = true
				a[x] = yy
				L := int64(x - pre)
				r := getB(x) - getB(pre)
				seg.update(x, buildEdgeMat(L, r, comb))
				if suc != 0 {
					L2 := int64(suc - x)
					r2 := getB(suc) - getB(x)
					seg.update(suc, buildEdgeMat(L2, r2, comb))
				}
			}
		}

		P := seg.AllProduct()
		last := set.Max()
		f1 := fib[n-last+2]
		f2 := fib[n-last+1]

		r0 := P.a00
		r1 := P.a01
		ans := (r0*f1 + r1*f2) % mod
		out((ans + mod) % mod)
	}
}
