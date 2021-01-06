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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K, _ := getI(), getI(), getI()
	a := make([]int, K)
	b := make([]int, K)
	for i := 0; i < K; i++ {
		aa, bb := getS(), getS()
		if aa == "?" {
			a[i], b[i] = -1, -1
		} else {
			a[i], _ = strconv.Atoi(aa)
			b[i], _ = strconv.Atoi(bb)
			a[i]--
			b[i]--
		}
	}
	// out(a, b)
	cN := getInts(N)
	c0 := make([]int, N)
	for i := 0; i < N; i++ {
		c0[i] = i + 1
	}
	for i := 0; i < K; i++ {
		p0, p1 := a[i], b[i]
		if p0 == -1 {
			break
		}
		c0[p0], c0[p1] = c0[p1], c0[p0]
	}
	for i := K - 1; i >= 0; i-- {
		p0, p1 := a[i], b[i]
		if p0 == -1 {
			break
		}
		cN[p0], cN[p1] = cN[p1], cN[p0]
	}

	// out(c0)
	// out(cN)
	for i := 0; i < N; i++ {
		if c0[i] != cN[i] {
			fmt.Fprint(wr, i+1, " ")
		}
	}
	out()
}
