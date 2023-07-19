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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	solve()
}

func solve() {
	n := getI()
	box := make([][]int, n)
	cnt := map[int]int{}
	cc := make([]int, n)
	for i := 0; i < n; i++ {
		box[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			box[i][j] = getI()
		}
		// h, w, d を h < w < dの関係にソート
		sort.Ints(box[i])
		cc[i] = box[i][1]
		cnt[box[i][0]]++
	}

	// 箱をｈでソートする
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})
	// 座標圧縮
	cp := newCompress(cc)
	t := NewTree(len(cp))

	for i := 0; i < n; {
		// hが同じものに対して繰り返す
		for j := 0; j < cnt[box[i][0]]; j++ {
			b := box[i+j]
			w := cp.Get(b[1])
			if t.Prod(0, w) < b[2] {
				out("Yes")
				return
			}
		}
		// 同じｈのものについて、最小を登録する
		for j := 0; j < cnt[box[i][0]]; j++ {
			b := box[i+j]
			w := cp.Get(b[1])
			old := t.Get(w)
			if old > b[2] {
				t.Set(w, b[2])
			}
		}
		i += cnt[box[i][0]]
	}
	out("No")
}

type compress map[int]int

func newCompress(aa []int) compress {
	n := len(aa)
	bb := make([]int, n)
	copy(bb, aa)
	sort.SliceStable(bb, func(i, j int) bool {
		return bb[i] < bb[j]
	})
	c := compress{}
	l := 0
	i := 0
	for r := 0; r < n; r++ {
		for r > l && bb[r] != bb[l] {
			l++
		}
		if r == l {
			c[bb[l]] = i
			i++
		}
	}
	return c
}
func (c compress) Get(x int) int {
	return c[x]
}

type S = int

func op(l, r S) S {
	return min(l, r)
}

func e() S {
	return math.MaxInt64
}

func NewTree(n int) Tree {
	t := make(Tree, n<<1)
	for i := 0; i < n<<1; i++ {
		t[i] = e()
	}
	return t
}

type Tree []S

func (t Tree) Set(p int, x S) {
	SegmentTree(len(t)>>1).Set(p, func(i int) {
		t[i] = x
	}, func(i, l, r int) {
		t[i] = op(t[l], t[r])
	})
}
func (t Tree) Get(p int) S {
	return t[SegmentTree(len(t)>>1).Get(p)]
}
func (t Tree) Prod(l, r int) S {
	sl := e()
	sr := e()
	SegmentTree(len(t)>>1).Prod(l, r, func(i int) {
		sl = op(sl, t[i])
	}, func(i int) {
		sr = op(t[i], sr)
	})
	return op(sl, sr)
}
func (t Tree) MaxRight(l int, f func(S) bool) int {
	sm := e()
	return SegmentTree(len(t)>>1).MaxRight(l, func(i int) {
		sm = op(sm, t[i])
	}, func(i int) bool {
		return f(op(sm, t[i]))
	})
}
func (t Tree) MinLeft(r int, f func(S) bool) int {
	sm := e()
	return SegmentTree(len(t)>>1).MinLeft(r, func(i int) {
		sm = op(t[i], sm)
	}, func(i int) bool {
		return f(op(t[i], sm))
	})
}

type SegmentTree int

func (s SegmentTree) Set(p int, set func(int), update func(int, int, int)) {
	p = s.Get(p)
	set(p)
	for p > 1 {
		p >>= 1
		update(p, p<<1, p<<1+1)
	}
}
func (s SegmentTree) Get(p int) int {
	return p + int(s)
}
func (s SegmentTree) Prod(l, r int, lop, rop func(int)) {
	l = s.Get(l)
	r = s.Get(r)
	for l < r {
		if l&1 == 1 {
			lop(l)
			l++
		}
		if r&1 == 1 {
			r--
			rop(r)
		}
		l >>= 1
		r >>= 1
	}
}
func (s SegmentTree) MaxRight(l int, update func(int), f func(int) bool) int {
	n := s.Get(0)
	if l == n {
		return n
	}
	subtree := func(root int) int {
		for root < n {
			root <<= 1
			if f(root) {
				update(root)
				root++
			}
		}
		return root - n
	}
	l = s.Get(l)
	r := s.Get(n)
	stack := make([]int, 0)
	for l < r {
		if l&1 == 1 {
			if !f(l) {
				return subtree(l)
			}
			update(l)
			l++
		}
		if r&1 == 1 {
			r--
			stack = append(stack, r)
		}
		l >>= 1
		r >>= 1
	}
	for len(stack) > 0 {
		l = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !f(l) {
			return subtree(l)
		}
		update(l)
	}
	return n
}
func (s SegmentTree) MinLeft(r int, update func(int), f func(int) bool) int {
	if r == 0 {
		return 0
	}
	n := s.Get(0)
	l := s.Get(0)
	r = s.Get(r)
	subtree := func(root int) int {
		for root < n {
			root = root<<1 + 1
			if f(root) {
				update(root)
				root--
			}
		}
		return root + 1 - n
	}
	stack := make([]int, 0)
	for l < r {
		if l&1 == 1 {
			stack = append(stack, l)
			l++
		}
		if r&1 == 1 {
			r--
			if !f(r) {
				return subtree(r)
			}
			update(r)
		}
		l >>= 1
		r >>= 1
	}
	for len(stack) > 0 {
		r = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !f(r) {
			return subtree(r)
		}
		update(r)
	}
	return 0
}
