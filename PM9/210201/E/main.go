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

func matmul(a, b [][]int) [][]int {
	n := len(a)
	m := len(b)
	l := len(b[0])
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, l)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < l; k++ {
				c[i][k] += a[i][j] * b[j][k]
			}
		}
	}
	return c
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	N := getI()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getI(), getI()
	}
	M := getI()
	a := [][][]int{{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}}
	for i := 0; i < M; i++ {
		op := getI()
		switch op {
		case 1:
			v := [][]int{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}}
			a = append(a, matmul(v, a[i]))
		case 2:
			v := [][]int{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}}
			a = append(a, matmul(v, a[i]))
		case 3:
			p := getI()
			v := [][]int{{-1, 0, 2 * p}, {0, 1, 0}, {0, 0, 1}}
			a = append(a, matmul(v, a[i]))
		case 4:
			p := getI()
			v := [][]int{{1, 0, 0}, {0, -1, 2 * p}, {0, 0, 1}}
			a = append(a, matmul(v, a[i]))
		}
	}

	Q := getI()
	for i := 0; i < Q; i++ {
		a0, b0 := getI(), getI()-1
		x := matmul(a[a0], [][]int{{x[b0]}, {y[b0]}, {1}})
		out(x[0][0], x[1][0])
	}
}
