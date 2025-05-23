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
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	n := getI()

	// Build the tree
	type edge struct {
		a, b int
	}
	es := make([]edge, n-1)
	to := make([][]int, n)

	for i := 0; i < n-1; i++ {
		a, b := getI()-1, getI()-1
		es[i] = edge{a, b}
		to[a] = append(to[a], b)
		to[b] = append(to[b], a)
	}

	in_ := make([]int, n)
	out_ := make([]int, n)
	vid := 0

	// DFS to assign in and out times
	var dfs func(int, int)
	dfs = func(v, p int) {
		in_[v] = vid
		vid++
		for _, u := range to[v] {
			if u != p {
				dfs(u, v)
			}
		}
		out_[v] = vid
	}

	dfs(0, -1)

	// Initialize BIT
	bit := newBIT(n)
	for i := 0; i < n; i++ {
		bit.add(i, 1)
	}

	q := getI()

	for qi := 0; qi < q; qi++ {
		type_ := getI()
		if type_ == 1 {
			v, w := getI()-1, getI()
			bit.add(in_[v], w)
		} else {
			ei := getI() - 1
			a, b := es[ei].a, es[ei].b
			if in_[a] < in_[b] {
				a, b = b, a
			}
			as := bit.rangeSum(in_[a], out_[a])
			bs := bit.rangeSum(0, n) - as
			out(abs(as - bs))
			// fmt.Println(abs(as - bs))
		}
	}
}
