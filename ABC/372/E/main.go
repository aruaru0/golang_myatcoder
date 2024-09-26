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

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	n            int
}

// newDsu :
func newDsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
	}
	return &d
}

// Merge :
func (d DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

// Same :
func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader :
func (d DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

// Size :
func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

// Groups : original implement
func (d DSU) Groups() [][]int {
	m := make(map[int][]int)
	for i := 0; i < d.n; i++ {
		x := d.Leader(i)
		if x < 0 {
			m[i] = append(m[i], i)
		} else {
			m[x] = append(m[x], i)
		}
	}
	ret := make([][]int, len(m))
	idx := 0
	for _, e := range m {
		ret[idx] = make([]int, len(e))
		copy(ret[idx], e)
		idx++
	}
	return ret
}

func merge(a, b []int) []int {
	ret := make([]int, 0)
	l, r := 0, 0
	for l < len(a) && r < len(b) {
		if a[l] < b[r] {
			ret = append(ret, a[l])
			l++
		} else {
			ret = append(ret, b[r])
			r++
		}
	}
	ret = append(ret, a[l:]...)
	ret = append(ret, b[r:]...)
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()

	uf := newDsu(N)
	p := make([][]int, N)
	for i := 0; i < N; i++ {
		p[i] = append(p[i], i)
	}

	for qi := 0; qi < Q; qi++ {
		t := getI()
		if t == 1 {
			u, v := getI()-1, getI()-1
			if !uf.Same(u, v) {
				x := p[uf.Leader(u)]
				y := p[uf.Leader(v)]
				// out(u, v, "xy", x, y)
				// z := merge(x, y)
				z := append(x, y...)
				sort.Ints(z)

				p[u] = []int{}
				p[v] = []int{}
				uf.Merge(u, v)
				l := uf.Leader(u)
				p[l] = z[max(0, len(z)-10):]
				// out(l, u, v, p)
			}
		} else {
			u, k := getI()-1, getI()
			l := uf.Leader(u)
			n := len(p[uf.Leader(u)])
			if n < k {
				out(-1)
			} else {
				// out("t==2 ", u, k, p)
				out(p[l][n-k] + 1)
			}
		}
	}

	// x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	// out(x[len(x)-10:])
}
