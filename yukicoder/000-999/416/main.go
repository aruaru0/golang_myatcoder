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

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	child        [][]int
	n            int
}

// Dsu :
func Dsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	d.child = make([][]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
		d.child[i] = append(d.child[i], i)
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
	d.child[x] = append(d.child[x], d.child[y]...)
	d.child[y] = []int{}
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

// Lists :
func (d DSU) Lists(a int) []int {
	l := d.Leader(a)
	return d.child[l]
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

type edge struct {
	from, to int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, Q := getI(), getI(), getI()
	x := make(map[edge]bool)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		x[edge{a, b}] = true
	}
	y := make([]edge, Q)
	for i := 0; i < Q; i++ {
		c, d := getI()-1, getI()-1
		y[i] = edge{c, d}
		x[edge{c, d}] = false
	}

	uf := Dsu(N)
	for p, e := range x {
		if e == true {
			uf.Merge(p.from, p.to)
		}
	}

	ans := make([]int, N)
	for i := 0; i < N; i++ {
		if uf.Same(0, i) {
			ans[i] = -1
		}
	}

	for i := Q - 1; i >= 0; i-- {
		e := y[i]
		from := uf.Same(0, e.from)
		to := uf.Same(0, e.to)
		flist := uf.Lists(e.from)
		tlist := uf.Lists(e.to)
		uf.Merge(e.from, e.to)
		if from == false && uf.Same(0, e.from) {
			for _, l := range flist {
				ans[l] = i + 1
			}
		}
		if to == false && uf.Same(0, e.to) {
			for _, l := range tlist {
				ans[l] = i + 1
			}
		}
	}

	for i := 1; i < len(ans); i++ {
		out(ans[i])
	}
}
