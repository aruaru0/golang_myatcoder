package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

type OP struct {
	t    int
	x, y int
}

type S struct {
	val int
}
type F struct {
	val int
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
	for i := range lseg.d {
		lseg.d[i] = lseg.e()
	}
	for i := range lseg.lz {
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
	// ループ条件を i <= 1 から i >= 1 に修正
	for i := lseg.log; i >= 1; i-- {
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
			// Prodの右端は開区間なので、r-1の親をPushする必要がある
			lseg.Push((r - 1) >> uint(i))
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

	Q := getI()

	op := make([]OP, Q)
	for qi := 0; qi < Q; qi++ {
		t := getI()
		if t == 1 {
			op[qi] = OP{t, getI(), -1}
		} else {
			op[qi] = OP{t, getI(), getI()}
		}
	}

	// 削除がないと仮定した場合の連結リストを構築
	next := make([]int, Q+1)
	for i := 0; i <= Q; i++ {
		next[i] = -1
	}

	for i := 0; i < Q; i++ {
		if op[i].t == 1 {
			// 値(i+1)をxの次に挿入
			next[i+1] = next[op[i].x]
			next[op[i].x] = i + 1
		}
	}

	// 連結リストを配列に変換
	p := []int{0}
	cur := 0
	for next[cur] != -1 {
		p = append(p, next[cur])
		cur = next[cur]
	}

	// 各値が配列pのどこにあるかを記録
	pos := make([]int, Q+1)
	for i, e := range p {
		pos[e] = i
	}

	// セグメント木の設定
	e := func() S { return S{0} }
	merger := func(a, b S) S { return S{a.val + b.val} }
	// f.val=1なら削除(値を0に)、0なら何もしない
	mapper := func(f F, x S) S {
		if f.val == 1 {
			return S{0}
		}
		return x
	}
	// 削除(1)が優先されるように合成
	comp := func(f, g F) F { return F{max(f.val, g.val)} }
	id := func() F { return F{0} }

	// セグメント木をpの長さで初期化
	seg := newLazySegtree(make([]S, len(p)), e, merger, mapper, comp, id)
	// 初期値0をセット
	seg.Set(pos[0], S{0})

	// クエリのシミュレーション
	for qi := 0; qi < Q; qi++ {
		if op[qi].t == 1 {
			// 値(qi+1)を対応する位置にセット
			seg.Set(pos[qi+1], S{qi + 1})
		} else {
			px, py := pos[op[qi].x], pos[op[qi].y]
			if px > py {
				px, py = py, px
			}
			if px+1 == py {
				// 間に要素がない
				out(0)
			} else {
				// 区間[px+1, py)の合計を求めて出力
				res := seg.Prod(px+1, py)
				out(res.val)
				// 同区間を削除(値を0に)
				seg.RangeApply(px+1, py, F{1})
			}
		}
	}
}
