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

func f(s, n int, node [][]int) int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	q := []int{s}
	dist[s] = 0
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, e := range node[cur] {
			if dist[e] == -1 {
				dist[e] = dist[cur] + 1
				q = append(q, e)
			}
		}
	}

	ret := 0
	for i := 0; i < n; i++ {
		ret = max(ret, dist[i])
	}

	// out(dist)
	// out(ret)

	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N1, N2, M := getI(), getI(), getI()

	node1 := make([][]int, N1)
	node2 := make([][]int, N2)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		if a < N1 {
			node1[a] = append(node1[a], b)
			node1[b] = append(node1[b], a)
		} else {
			a -= N1
			b -= N1
			node2[a] = append(node2[a], b)
			node2[b] = append(node2[b], a)
		}
	}

	ret1 := f(0, N1, node1)
	ret2 := f(N2-1, N2, node2)

	out(ret1 + ret2 + 1)
}
