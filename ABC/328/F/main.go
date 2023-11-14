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

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	weight       []int
	n            int
}

// newDsu :
func newDsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	d.weight = make([]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = i
	}
	return &d
}

// Merge :
func (d DSU) Merge(a, b, w int) bool {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		// out("a, b, w", x, y, d.weight)
		return d.weight[b]-d.weight[a] == w
	}
	d.parentOrSize[x] = y
	d.weight[x] = d.weight[b] - d.weight[a] - w
	return true
}

// Same :
func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader :
func (d DSU) Leader(i int) int {
	if d.parentOrSize[i] == i {
		return i
	}
	r := d.Leader(d.parentOrSize[i])
	d.weight[i] = d.weight[d.parentOrSize[i]] + d.weight[i]
	d.parentOrSize[i] = r
	return r
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()

	uf := newDsu(N)
	for qi := 0; qi < Q; qi++ {
		a, b, d := getI()-1, getI()-1, getI()
		// out(qi, a, b, d, uf.parentOrSize, uf.weight)
		if uf.Merge(a, b, d) == true {
			fmt.Fprint(wr, qi+1, " ")
		}
	}
	out()
}
