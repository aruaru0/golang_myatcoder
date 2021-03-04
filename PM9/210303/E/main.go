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

var N int
var A, B, c []int
var node [][]int
var depth []int

func dfs(v, p, n int) {
	depth[v] = n
	for _, e := range node[v] {
		if e == p {
			continue
		}
		dfs(e, v, n+1)
	}
}

func dfs2(v, p, x int) {
	c[v] += x
	for _, e := range node[v] {
		if e == p {
			continue
		}
		dfs2(e, v, c[v])
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	A = make([]int, N-1)
	B = make([]int, N-1)
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		A[i], B[i] = getI()-1, getI()-1
		node[A[i]] = append(node[A[i]], B[i])
		node[B[i]] = append(node[B[i]], A[i])
	}

	depth = make([]int, N)
	dfs(0, -1, 0)

	c = make([]int, N)
	Q := getI()
	for i := 0; i < Q; i++ {
		t, e, x := getI(), getI()-1, getI()
		a, b := A[e], B[e]
		if t == 2 {
			a, b = b, a
		}
		if depth[a] < depth[b] {
			c[0] += +x
			c[b] += -x
		} else {
			c[a] += +x
		}
	}

	dfs2(0, -1, 0)

	for _, e := range c {
		out(e)
	}
}
