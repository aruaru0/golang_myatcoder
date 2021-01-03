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

var node [][]int
var used []bool
var A, B []int
var cost []int

func dfs(v, c, prev int) {
	cost[v] = c + A[v] + B[v]
	c = cost[v]
	for _, e := range node[v] {
		if e == prev {
			continue
		}
		dfs(e, c, v)
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node = make([][]int, N)
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N-1; i++ {
		f, t := getI()-1, getI()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
		a[i] = f
		b[i] = t
	}

	depth := make([]int, N)
	q := make([][2]int, 0)
	q = append(q, [2]int{0, 0})
	used := make([]bool, N)
	for len(q) != 0 {
		cur := q[0][0]
		d := q[0][1]
		q = q[1:]
		used[cur] = true
		depth[cur] = d
		for _, e := range node[cur] {
			if used[e] {
				continue
			}
			q = append(q, [2]int{e, d + 1})
		}
	}

	A = make([]int, N)
	B = make([]int, N)
	Q := getI()
	for i := 0; i < Q; i++ {
		t, e, x := getI(), getI()-1, getI()
		if t == 1 {
			if depth[a[e]] < depth[b[e]] {
				A[0] += x
				B[b[e]] -= x
			} else {
				A[a[e]] += x
			}
		} else {
			if depth[a[e]] > depth[b[e]] {
				A[0] += x
				B[a[e]] -= x
			} else {
				A[b[e]] += x
			}
		}
	}
	cost = make([]int, N)
	// out(A, B)
	dfs(0, 0, 0)
	for _, e := range cost {
		out(e)
	}
}
