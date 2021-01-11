package main

import (
	"bufio"
	"fmt"
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

type days struct {
	a, b, c int
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

func f(d []days) {
	a := make([]int, 100)
	for _, e := range d {
		a[e.a] = e.c
		a[e.b+1] = -e.c
	}
	for i := 1; i < 100; i++ {
		a[i] += a[i-1]
	}
	ans := 0
	for i := 0; i < 31; i++ {
		fmt.Fprint(wr, "{", i, a[i], "}")
		ans += a[i]
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, C := getI(), getI()

	d := make([]days, N)
	x := make(map[int]int, 0)
	for i := 0; i < N; i++ {
		a, b, c := getI(), getI(), getI()
		d[i] = days{a, b, c}
		x[a] = 1
		x[b+1] = 1
	}

	t := make([]int, 0, len(x))
	for e := range x {
		t = append(t, e)
	}

	sort.Ints(t)
	for i := 0; i < len(t); i++ {
		x[t[i]] = i
	}
	bit := newBIT(len(t) + 100)

	for i := 0; i < N; i++ {
		l := x[d[i].a]
		r := x[d[i].b+1]
		bit.add(l, d[i].c)
		bit.add(r, -d[i].c)
	}
	ans := 0
	for i := 0; i < len(t)-1; i++ {
		v := bit.sum(i)
		ll := t[i+1] - t[i]
		ans += min(v, C) * ll
		// out(v, ll, t[i], t[i+1], ans)
	}
	out(ans)
}
