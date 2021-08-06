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

type pair struct {
	x, y int
}
type pos struct {
	n, x, y int
}

var memo map[pos]*pair

func rec(n, x, y int) (int, int) {
	if n == N {
		return 0, 0
	}
	if memo[pos{n, x, y}] != nil {
		ret := *memo[pos{n, x, y}]
		return ret.x, ret.y
	}
	x0, y0 := rec(n+1, x, y)
	x1, y1 := rec(n+1, x+X[n], y+Y[n])
	x1 += X[n]
	y1 += Y[n]

	xx0 := x0 + x
	yy0 := y0 + y
	xx1 := x1 + x
	yy1 := y1 + y

	if xx0*xx0+yy0*yy0 > xx1*xx1+yy1*yy1 {
		memo[pos{n, x, y}] = &pair{x0, y0}
		return x0, y0
	}
	memo[pos{n, x, y}] = &pair{x1, y1}
	return x1, y1
}

var N int
var X, Y []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	X = make([]int, N)
	Y = make([]int, N)
	for i := 0; i < N; i++ {
		X[i], Y[i] = getI(), getI()
	}

	memo = make(map[pos]*pair)
	x, y := rec(0, 0, 0)

	out(math.Sqrt(float64(x*x + y*y)))
}
