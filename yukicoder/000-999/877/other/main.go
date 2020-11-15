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

type bit struct {
	v []int
}

func newBIT(n int) *bit {
	b := new(bit)
	b.v = make([]int, n)
	return b
}
func (b bit) sum(a int) int {
	ret := 0
	for i := a + 1; i > 0; i -= i & -i {
		ret += b.v[i-1]
	}
	return ret
}
func (b bit) rangeSum(x, y int) int {
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
func (b bit) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
	}
}

type data struct {
	t       int
	i       int
	x, l, r int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()
	d := make([]data, 0, N+Q)
	for i := 0; i < N; i++ {
		d = append(d, data{0, i, getI(), 0, 0})
	}
	for i := 0; i < Q; i++ {
		_, l, r, x := getI(), getI(), getI(), getI()
		d = append(d, data{1, i, x, l - 1, r})
	}

	sort.Slice(d, func(i, j int) bool {
		// if d[i].x == d[j].x {
		// 	return d[i].t > d[j].t
		// }
		return d[i].x > d[j].x
	})

	b0 := newBIT(N)
	b1 := newBIT(N)
	ans := make([]int, Q)
	for _, e := range d {
		if e.t == 0 {
			b0.add(e.i, e.x)
			b1.add(e.i, 1)
		} else {
			ans[e.i] = b0.rangeSum(e.l, e.r) - b1.rangeSum(e.l, e.r)*e.x
		}
	}
	for i := 0; i < Q; i++ {
		out(ans[i])
	}
}
