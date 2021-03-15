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
	r1, c1 := getI(), getI()
	r2, c2 := getI(), getI()

	r2 -= r1
	c2 -= c1
	r1, c1 = 0, 0

	// out(r1, c1, r2, c2)

	if r1 == r2 && c1 == c2 {
		out(0)
		return
	}

	if abs(r2-r1)+abs(c2-c1) <= 3 {
		out(1)
		return
	}

	if r1+c1 == r2+c2 {
		out(1)
		return
	}

	if r1-c1 == r2-c2 {
		out(1)
		return
	}

	for r := -3; r <= 3; r++ {
		for c := -3; c <= 3; c++ {
			r3, c3 := r1+r, c1+c
			if abs(r1+r3)+abs(c1-c3) > 3 {
				continue
			}
			if abs(r2-r3)+abs(c2-c3) <= 3 {
				out(2)
				return
			}
			if r3+c3 == r2+c2 {
				out(2)
				return
			}

			if r3-c3 == r2-c2 {
				out(2)
				return
			}
		}
	}

	if (r1+c1)%2 == (r2+c2)%2 {
		out(2)
		return
	}

	out(3)
}
