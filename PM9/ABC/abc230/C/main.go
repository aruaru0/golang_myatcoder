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

func check1(x, y int) bool {
	kmin := max(1-A, 1-B)
	kmax := min(N-A, N-B)

	dx := x - A
	dy := y - B

	if dx != dy {
		return false
	}
	if kmin <= dx && dx <= kmax {
		return true
	}
	return false
}

func check2(x, y int) bool {
	kmin := max(1-A, B-N)
	kmax := min(N-A, B-1)

	dx := x - A
	dy := B - y
	// out(x, y, dx, dy)

	if dx != dy {
		return false
	}
	if kmin <= dx && dx <= kmax {
		return true
	}
	return false
}

var N, A, B int
var P, Q, R, S int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, A, B = getI(), getI(), getI()
	P, Q, R, S = getI(), getI(), getI(), getI()

	h := Q - P + 1
	w := S - R + 1
	m := make([]byte, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if check1(P+i, R+j) || check2(P+i, R+j) {
				m[j] = '#'
			} else {
				m[j] = '.'
			}
		}
		out(string(m))
	}
}
