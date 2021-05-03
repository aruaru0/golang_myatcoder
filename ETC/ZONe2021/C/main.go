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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	d := make([]int, N)
	e := make([]int, N)
	amax, bmax, cmax, dmax, emax := 0, 0, 0, 0, 0
	for i := 0; i < N; i++ {
		a[i], b[i], c[i], d[i], e[i] = getI(), getI(), getI(), getI(), getI()
		chmax(&amax, a[i])
		chmax(&bmax, b[i])
		chmax(&cmax, c[i])
		chmax(&dmax, d[i])
		chmax(&emax, e[i])
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			A := max(a[i], a[j])
			B := max(b[i], b[j])
			C := max(c[i], c[j])
			D := max(d[i], d[j])
			E := max(e[i], e[j])
			tot0 := nmin(amax, B, C, D, E)
			tot1 := nmin(A, bmax, C, D, E)
			tot2 := nmin(A, B, cmax, D, E)
			tot3 := nmin(A, B, C, dmax, E)
			tot4 := nmin(A, B, C, D, emax)
			ans = nmax(ans, tot0, tot1, tot2, tot3, tot4)
		}
	}
	out(ans)
}
