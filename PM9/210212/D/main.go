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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, C := getI(), getI()
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		a[i], b[i], c[i] = getI(), getI()+1, getI()
		m[a[i]] = 1
		m[b[i]] = 1
	}
	p := make([]int, 0)
	for e := range m {
		p = append(p, e)
	}
	sort.Ints(p)
	for i := 0; i < len(p); i++ {
		m[p[i]] = i + 1
	}
	// out(m, C)
	n := len(m)
	d := make([]int, n+1)
	for i := 0; i < N; i++ {
		d[m[a[i]]] += c[i]
		d[m[b[i]]] += -c[i]
	}
	// out(d)
	for i := 1; i <= n; i++ {
		d[i] += d[i-1]
	}

	// out(p, n)
	tot := 0
	for i := 1; i < n; i++ {
		from := p[i-1]
		to := p[i] - 1
		// out(from, to)
		cost := d[i]
		if d[i] > C {
			cost = C
		}
		tot += cost * (to - from + 1)
	}
	out(tot)
}