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

const inf = 1e9

func calc(t float64) (float64, float64) {
	xmin, xmax := inf, -inf
	ymin, ymax := inf, -inf
	for i := 0; i < N; i++ {
		xx, yy := x[i], y[i]
		switch d[i] {
		case 'R':
			xx += t
		case 'L':
			xx -= t
		case 'U':
			yy += t
		case 'D':
			yy -= t
		}
		xmin = math.Min(xmin, xx)
		xmax = math.Max(xmax, xx)
		ymin = math.Min(ymin, yy)
		ymax = math.Max(ymax, yy)
	}
	return xmax - xmin, ymax - ymin
}

var N int
var x, y []float64
var d []byte

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	x = make([]float64, N)
	y = make([]float64, N)
	d = make([]byte, N)
	for i := 0; i < N; i++ {
		x[i], y[i], d[i] = getF(), getF(), ([]byte(getS()))[0]
	}

	l, r := 0.0, inf
	for k := 0; k < 300; k++ {
		m0 := l + (r-l)/3
		m1 := l + 2*(r-l)/3
		x0, y0 := calc(m0)
		x1, y1 := calc(m1)
		// out(l, r, "m", m0, m1, "xy", x0*y0, x1*y1)
		if x0*y0 < x1*y1 {
			r = m1
		} else {
			l = m0
		}
	}
	// out(l, r)
	a, b := calc((l + r) / 2)
	out(a * b)
}
