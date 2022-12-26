type S struct {
	a [4][4]int
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
