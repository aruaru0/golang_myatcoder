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

type town struct {
	a, b int
}

func solve(p []town) {
	N := len(p)
	n := 1 << N
	ans := 1000000
	for i := 0; i < n; i++ {
		cnt := 0
		a := 0
		b := 0
		for j := 0; j < N; j++ {
			if (i>>j)&1 == 1 {
				b += p[j].a + p[j].b
				cnt++
			} else {
				a += p[j].a
			}
		}
		// out(a, b, cnt)
		if a < b {
			ans = min(ans, cnt)
		}
	}
	out(ans)
}

var N int
var p []town
var A int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	p = make([]town, N)
	A = 0
	for i := 0; i < N; i++ {
		a, b := getI(), getI()
		A += a
		p[i] = town{a, b}
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].a+p[i].b-p[j].a > p[j].a+p[j].b-p[i].a
	})

	// out(A, p)
	tot := 0
	idx := 0
	for tot <= A {
		tot += p[idx].a + p[idx].b
		A -= p[idx].a
		// out(tot, A)
		idx++
	}
	out(idx)
	// solve(p)
}
