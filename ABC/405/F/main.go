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

type pair struct{ a, b int }

type query struct {
	c, d int
	idx  int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	SZ := 2 * N

	segs := make([]pair, M)
	endCnt := make([]int, SZ+1) // 1-based

	for i := 0; i < M; i++ {
		a, b := getI(), getI()
		if a > b {
			a, b = b, a
		}
		segs[i] = pair{a, b}
		endCnt[a]++
		endCnt[b]++
	}

	/* prefix sums of endpoint counts */
	pref := make([]int, SZ+1)
	for i := 1; i <= SZ; i++ {
		pref[i] = pref[i-1] + endCnt[i]
	}

	/* sort segments by right endpoint */
	sort.Slice(segs, func(i, j int) bool { return segs[i].b < segs[j].b })

	Q := getI()
	qs := make([]query, Q)
	for i := 0; i < Q; i++ {
		c, d := getI(), getI()
		if c > d {
			c, d = d, c
		}
		qs[i] = query{c: c, d: d, idx: i}
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i].d < qs[j].d })

	BIT := newBIT(SZ)
	ans := make([]int, Q)
	p := 0 // pointer over segs

	for _, q := range qs {
		/* sweep segments whose right end < q.d */
		for p < M && segs[p].b < q.d {
			BIT.add(segs[p].a, 1)
			out(segs[p].a, "+1")
			p++
		}
		insideEnd := pref[q.d-1] - pref[q.c]         // #端点 in (c,d)
		insideChord := BIT.sum(q.d-1) - BIT.sum(q.c) // #弦 C<A<B<D
		out(q, insideEnd, insideChord)
		ans[q.idx] = insideEnd - 2*insideChord // 端点が1つだけの弦
	}

	for _, v := range ans {
		out(v)
	}

}
