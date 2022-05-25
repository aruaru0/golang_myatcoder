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

type rbtreemultisetnode struct {
	left, right, up int32
	red             bool
	key             int
	count           int
}
type rbtreemultiset struct {
	lessthan func(a, b int) bool
	tree     []rbtreemultisetnode
	root     int32
	recycler []int32
	sz       int
	minidx   int32
	maxidx   int32
}

type rbtreemultisetIterator interface {
	Next() (ok bool)
	Prev() (ok bool)
	Key() int
	Count() int
}
type rbtreemultisetiter struct {
	cur    int32
	key    int
	count  int
	rbtree *rbtreemultiset
}

func (i *rbtreemultisetiter) Key() int   { return i.key }
func (i *rbtreemultisetiter) Count() int { return i.count }
func (i *rbtreemultisetiter) Next() bool {
	rbtree := i.rbtree
	v := rbtree.nextidx(i.cur)
	if v == 0 {
		return false
	}
	i.cur, i.key, i.count = v, rbtree.tree[v].key, rbtree.tree[v].count
	return true
}
func (i *rbtreemultisetiter) Prev() bool {
	rbtree := i.rbtree
	v := rbtree.previdx(i.cur)
	if v == 0 {
		return false
	}
	i.cur, i.key, i.count = v, rbtree.tree[v].key, rbtree.tree[v].count
	return true
}
func Newrbtreemultiset(lessthan func(a, b int) bool) *rbtreemultiset {
	q := &rbtreemultiset{lessthan, make([]rbtreemultisetnode, 2), int32(0), make([]int32, 0), 0, 0, 0}
	q.tree[0].left, q.tree[0].right, q.tree[0].up, q.tree[0].red = 0, 0, 0, false
	q.recycler = append(q.recycler, 1)
	return q
}
func (q *rbtreemultiset) Add(k int) {
	if q.sz == 0 {
		z := q.getNewNodenum()
		tree := q.tree
		q.minidx, q.maxidx, q.sz, q.root = z, z, q.sz+1, z
		tree[z].key, tree[z].count, tree[z].up, tree[z].left, tree[z].right, tree[z].red = k, 1, 0, 0, 0, false
		return
	}
	y, cmp := q.findInsertionPoint(k)
	if cmp == 0 {
		q.tree[y].count++
		q.sz += 1
		return
	}
	z := q.getNewNodenum()
	q.sz += 1
	tree := q.tree
	tree[z].key, tree[z].count, tree[z].up, tree[z].left, tree[z].right, tree[z].red = k, 1, y, 0, 0, true
	if cmp < 0 {
		tree[y].left = z
	} else {
		tree[y].right = z
	}
	if q.sz == 0 || q.lessthan(k, tree[q.minidx].key) {
		q.minidx = z
	}
	if q.sz == 0 || q.lessthan(tree[q.maxidx].key, k) {
		q.maxidx = z
	}
	var p, g, u int32
	for p = tree[z].up; tree[p].red; p = tree[z].up {
		g = tree[p].up
		if g == 0 {
			break
		}
		if p == tree[g].left {
			u = tree[g].right
			if tree[u].red {
				tree[p].red, tree[u].red, tree[g].red, z = false, false, true, g
				continue
			}
			if z == tree[p].right {
				z = p
				q.rotleft(z)
				p = tree[z].up
			}
			q.rotright(g)
			tree[g].red, tree[p].red = true, false
		} else {
			u = tree[g].left
			if tree[u].red {
				tree[p].red, tree[u].red, tree[g].red, z = false, false, true, g
				continue
			}
			if z == tree[p].left {
				z = p
				q.rotright(z)
				p = tree[z].up
			}
			q.rotleft(g)
			tree[g].red, tree[p].red = true, false
		}
	}
	tree[q.root].red = false
}
func (q *rbtreemultiset) Delete(k int) bool {
	if q.sz == 0 {
		return false
	}
	z, cmp := q.findInsertionPoint(k)
	if cmp != 0 {
		return false
	} else if q.tree[z].count > 1 {
		q.tree[z].count--
		q.sz--
		return true
	}
	q.sz--
	q.recycler = append(q.recycler, z)
	if q.sz > 0 && !q.lessthan(q.tree[q.minidx].key, k) {
		q.minidx = q.nextidx(q.minidx)
	}
	if q.sz > 0 && !q.lessthan(k, q.tree[q.maxidx].key) {
		q.maxidx = q.previdx(q.maxidx)
	}
	if q.sz == 0 {
		q.root = 0
		return true
	}
	tree := q.tree
	var x int32
	y, y_orig_red := z, tree[z].red
	if tree[z].left == 0 {
		x = tree[z].right
		q.rbTransplant(z, x)
	} else if tree[z].right == 0 {
		x = tree[z].left
		q.rbTransplant(z, x)
	} else {
		y = q.findminidx(tree[z].right)
		y_orig_red = tree[y].red
		x = tree[y].right
		if tree[y].up == z {
			tree[x].up = y
		} else {
			q.rbTransplant(y, x)
			tree[y].right = tree[z].right
			tree[tree[y].right].up = y
		}
		q.rbTransplant(z, y)
		tree[y].left = tree[z].left
		tree[tree[y].left].up = y
		tree[y].red = tree[z].red
	}
	if !y_orig_red {
		for q.root != x && !tree[x].red {
			p := tree[x].up
			if tree[p].left == x {
				s := tree[p].right
				if tree[s].red {
					tree[s].red = false
					tree[p].red = true
					q.rotleft(p)
					s = tree[p].right
				}
				c := tree[s].left
				d := tree[s].right
				if !tree[c].red && !tree[d].red {
					tree[s].red = true
					x = p
				} else {
					if !tree[d].red {
						tree[c].red = false
						tree[s].red = true
						q.rotright(s)
						s = tree[p].right
					}
					tree[s].red = tree[p].red
					tree[p].red = false
					tree[tree[s].right].red = false
					q.rotleft(p)
					x = q.root
				}
			} else {
				s := tree[p].left
				if tree[s].red {
					tree[s].red = false
					tree[p].red = true
					q.rotright(p)
					s = tree[p].left
				}
				c := tree[s].right
				d := tree[s].left
				if !tree[c].red && !tree[d].red {
					tree[s].red = true
					x = p
				} else {
					if !tree[d].red {
						tree[c].red = false
						tree[s].red = true
						q.rotleft(s)
						s = tree[p].left
					}
					tree[s].red = tree[p].red
					tree[p].red = false
					tree[tree[s].left].red = false
					q.rotright(p)
					x = q.root
				}
			}
		}
		tree[x].red = false
	}
	return true
}
func (q *rbtreemultiset) Clear() {
	q.tree, q.root, q.recycler, q.sz = q.tree[:2], 0, q.recycler[:0], 0
	q.recycler = append(q.recycler, int32(1))
}
func (q *rbtreemultiset) IsEmpty() bool       { return q.sz == 0 }
func (q *rbtreemultiset) Contains(k int) bool { _, cmp := q.findInsertionPoint(k); return cmp == 0 }
func (q *rbtreemultiset) Count(k int) int {
	z, cmp := q.findInsertionPoint(k)
	if cmp != 0 {
		return 0
	}
	return q.tree[z].count
}
func (q *rbtreemultiset) Len() int { return q.sz }
func (q *rbtreemultiset) MinKey() (k int) {
	if q.sz == 0 {
		panic("Called MinKey on an empty rbtreemultiset")
	}
	return q.tree[q.minidx].key
}
func (q *rbtreemultiset) MaxKey() (k int) {
	if q.sz == 0 {
		panic("Called MaxKey on an empty rbtreemultiset")
	}
	return q.tree[q.maxidx].key
}
func (q *rbtreemultiset) findLtIdx(k int) (int32, bool) {
	if q.sz == 0 || !q.lessthan(q.tree[q.minidx].key, k) {
		return 0, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos != 1 {
		idx = q.previdx(idx)
	}
	return idx, true
}
func (q *rbtreemultiset) findLeIdx(k int) (int32, bool) {
	if q.sz == 0 || q.lessthan(k, q.tree[q.minidx].key) {
		return 0, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos == -1 {
		idx = q.previdx(idx)
	}
	return idx, true
}
func (q *rbtreemultiset) findGtIdx(k int) (int32, bool) {
	if q.sz == 0 || !q.lessthan(k, q.tree[q.maxidx].key) {
		return 0, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos != -1 {
		idx = q.nextidx(idx)
	}
	return idx, true
}
func (q *rbtreemultiset) findGeIdx(k int) (int32, bool) {
	if q.sz == 0 || q.lessthan(q.tree[q.maxidx].key, k) {
		return 0, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos == 1 {
		idx = q.nextidx(idx)
	}
	return idx, true
}
func (q *rbtreemultiset) FindLt(k int) (int, bool) {
	var ans int
	idx, ok := q.findLtIdx(k)
	if ok {
		ans = q.tree[idx].key
	}
	return ans, ok
}
func (q *rbtreemultiset) FindLe(k int) (int, bool) {
	var ans int
	idx, ok := q.findLeIdx(k)
	if ok {
		ans = q.tree[idx].key
	}
	return ans, ok
}
func (q *rbtreemultiset) FindGt(k int) (int, bool) {
	var ans int
	idx, ok := q.findGtIdx(k)
	if ok {
		ans = q.tree[idx].key
	}
	return ans, ok
}
func (q *rbtreemultiset) FindGe(k int) (int, bool) {
	var ans int
	idx, ok := q.findGeIdx(k)
	if ok {
		ans = q.tree[idx].key
	}
	return ans, ok
}
func (q *rbtreemultiset) FindLtIter(k int) (rbtreemultisetIterator, bool) {
	var ans *rbtreemultisetiter
	idx, ok := q.findLtIdx(k)
	if ok {
		ans = &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}
	}
	return ans, ok
}
func (q *rbtreemultiset) FindLeIter(k int) (rbtreemultisetIterator, bool) {
	var ans *rbtreemultisetiter
	idx, ok := q.findLeIdx(k)
	if ok {
		ans = &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}
	}
	return ans, ok
}
func (q *rbtreemultiset) FindGtIter(k int) (rbtreemultisetIterator, bool) {
	var ans *rbtreemultisetiter
	idx, ok := q.findGtIdx(k)
	if ok {
		ans = &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}
	}
	return ans, ok
}
func (q *rbtreemultiset) FindGeIter(k int) (rbtreemultisetIterator, bool) {
	var ans *rbtreemultisetiter
	idx, ok := q.findGeIdx(k)
	if ok {
		ans = &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}
	}
	return ans, ok
}
func (q *rbtreemultiset) FindIter(k int) (rbtreemultisetIterator, bool) {
	if q.sz == 0 {
		return nil, false
	}
	idx, pos := q.findInsertionPoint(k)
	if pos != 0 {
		return nil, false
	}
	return &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *rbtreemultiset) MinIter() (rbtreemultisetIterator, bool) {
	if q.sz == 0 {
		return nil, false
	}
	idx := q.findminidx(q.root)
	return &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *rbtreemultiset) MaxIter() (rbtreemultisetIterator, bool) {
	if q.sz == 0 {
		return nil, false
	}
	idx := q.findmaxidx(q.root)
	return &rbtreemultisetiter{idx, q.tree[idx].key, q.tree[idx].count, q}, true
}
func (q *rbtreemultiset) rbTransplant(u, v int32) {
	tree := q.tree
	if tree[u].up == 0 {
		q.root = v
	} else {
		p := tree[u].up
		if u == tree[p].left {
			tree[p].left = v
		} else {
			tree[p].right = v
		}
	}
	tree[v].up = tree[u].up
}
func (q *rbtreemultiset) findInsertionPoint(k int) (int32, int8) {
	n, lt, tree := q.root, q.lessthan, q.tree
	for {
		nkey := tree[n].key
		if lt(nkey, k) {
			r := tree[n].right
			if r == 0 {
				return n, 1
			}
			n = r
		} else if lt(k, nkey) {
			l := tree[n].left
			if l == 0 {
				return n, -1
			}
			n = l
		} else {
			return n, 0
		}
	}
}
func (q *rbtreemultiset) findmaxidx(n1 int32) int32 {
	tree := q.tree
	for {
		xx := tree[n1].right
		if xx == 0 {
			break
		}
		n1 = xx
	}
	return n1
}
func (q *rbtreemultiset) findminidx(n1 int32) int32 {
	tree := q.tree
	for {
		xx := tree[n1].left
		if xx == 0 {
			break
		}
		n1 = xx
	}
	return n1
}
func (q *rbtreemultiset) nextidx(cur int32) int32 {
	last := int32(-2)
	tree := q.tree
	rr := tree[cur].right
	if rr > 0 {
		return q.findminidx(rr)
	}
	for {
		last, cur = cur, tree[cur].up
		if cur == 0 || tree[cur].left == last {
			break
		}
	}
	return cur
}
func (q *rbtreemultiset) previdx(cur int32) int32 {
	last := int32(0)
	tree := q.tree
	ll := tree[cur].left
	if ll > 0 {
		return q.findmaxidx(ll)
	}
	for {
		last, cur = cur, tree[cur].up
		if cur == 0 || tree[cur].right == last {
			break
		}
	}
	return cur
}
func (q *rbtreemultiset) rotleft(x int32) {
	tree := q.tree
	y := tree[x].right
	p := tree[x].up
	tree[x].right = tree[y].left
	if tree[y].left != 0 {
		tree[tree[y].left].up = x
	}
	tree[y].up = p
	if p == 0 {
		q.root = y
	} else if x == tree[p].left {
		tree[p].left = y
	} else {
		tree[p].right = y
	}
	tree[y].left = x
	tree[x].up = y
}
func (q *rbtreemultiset) rotright(x int32) {
	tree := q.tree
	y := tree[x].left
	p := tree[x].up
	tree[x].left = tree[y].right
	if tree[y].right != 0 {
		tree[tree[y].right].up = x
	}
	tree[y].up = p
	if p == 0 {
		q.root = y
	} else if x == tree[p].right {
		tree[p].right = y
	} else {
		tree[p].left = y
	}
	tree[y].right = x
	tree[x].up = y
}
func (q *rbtreemultiset) getNewNodenum() int32 {
	l := len(q.recycler)
	newnode := q.recycler[l-1]
	q.recycler = q.recycler[:l-1]
	if l == 1 {
		q.tree = append(q.tree, rbtreemultisetnode{})
		q.recycler = append(q.recycler, int32(len(q.tree)-1))
	}
	return newnode
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	Q := getI()
	s := Newrbtreemultiset(func(a, b int) bool { return a < b })
	for q := 0; q < Q; q++ {
		t := getI()
		x := getI()
		if t == 1 {
			s.Add(x)
		} else if t == 2 {
			k := getI()
			it, ok := s.FindLeIter(x)
			s := 0
			if ok {
				s = it.Count()
			}
			for ok && s < k {
				ok = it.Prev()
				if ok {
					s += it.Count()
				}
			}
			if s >= k {
				out(it.Key())
			} else {
				out("-1")
			}
		} else {
			k := getI()
			it, ok := s.FindGeIter(x)
			s := 0
			if ok {
				s = it.Count()
			}
			for ok && s < k {
				ok = it.Next()
				if ok {
					s += it.Count()
				}
			}
			if s >= k {
				out(it.Key())
			} else {
				out("-1")
			}
		}
	}
}
