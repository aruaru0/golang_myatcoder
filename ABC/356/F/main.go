package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
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

type BIT struct {
	v []int
}

func newBIT(n int) *BIT {
	b := new(BIT)
	b.v = make([]int, n)
	return b
}
func (b BIT) sum(a int) int {
	ret := 0
	for i := a + 1; i > 0; i -= i & -i {
		ret += b.v[i-1]
	}
	return ret
}
func (b BIT) rangeSum(x, y int) int {
	if y == 0 {
		return 0
	}
	y--
	if x == 0 {
		return b.sum(y)
	} else {
		return b.sum(y) - b.sum(x-1)
	}
}
func (b BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	q := getI()
	k := getI()
	mx := map[int]bool{}
	qq := make([][2]int, q)
	for i := 0; i < q; i++ {
		t := getI()
		x := getI()
		qq[i][0] = t
		qq[i][1] = x
		mx[x] = true
	}
	cx := make([]int, 0)
	for x := range mx {
		cx = append(cx, x)
	}
	const Inf = int(1e18)
	cx = append(cx, -Inf)
	cx = append(cx, Inf)
	sort.Ints(cx)
	n := len(cx)
	tox := map[int]int{}
	for i := 0; i < len(cx); i++ {
		tox[cx[i]] = i
	}
	f := newBIT(n)
	tree := redblacktree.NewWithIntComparator()
	treeFar := redblacktree.NewWithIntComparator()
	tree.Put(-Inf, true)
	tree.Put(math.MaxInt64, true)
	treeFar.Put(-Inf, true)
	for i := 0; i < q; i++ {
		x := qq[i][1]
		j := tox[qq[i][1]]
		if qq[i][0] == 1 {
			_, found := tree.Get(x)
			l, _ := tree.Floor(x - 1)
			lk := l.Key.(int)
			r, _ := tree.Ceiling(x + 1)
			rk := r.Key.(int)
			if found {
				lj := tox[lk]
				tree.Remove(lj)
				if x-lk > k {
					treeFar.Remove(lk)
				}
				if rk-x > k {
					treeFar.Remove(x)
				}
				if rk-lk > k {
					treeFar.Put(lk, true)
				}
				tree.Remove(x)
				f.add(j, -1)
			} else {
				if rk-lk > k {
					treeFar.Remove(lk)
				}
				if x-lk > k {
					treeFar.Put(lk, true)
				}
				if rk-x > k {
					treeFar.Put(x, true)
				}
				f.add(j, 1)
				tree.Put(x, true)
			}
			continue
		}
		l, _ := treeFar.Floor(x - 1)
		lk := l.Key.(int)
		r, found := treeFar.Ceiling(x)
		lnext, _ := tree.Ceiling(lk + 1)
		rnext := Inf
		if found {
			rnext = r.Key.(int)
		}
		out(f.rangeSum(tox[lnext.Key.(int)], tox[rnext]+1))
	}
}
