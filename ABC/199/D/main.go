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

func dfs1(v, parent int) {
	if visited[v] {
		return
	}
	visited[v] = true
	nodes = append(nodes, v)
	parents[v] = parent
	for _, u := range node[v] {
		if u == parent {
			continue
		}
		dfs1(u, v)
	}
}

func dfs2(i int) int {
	if i == len(nodes) {
		for _, a := range nodes {
			for _, b := range node[a] {
				if c[a] == c[b] {
					return 0
				}
			}
		}
		return 1
	}
	ret := 0
	if i == 0 {
		for j := 0; j < 3; j++ {
			c[nodes[i]] = j
			ret += dfs2(i + 1)
		}
	} else {
		for j := 0; j < 3; j++ {
			if j != c[parents[nodes[i]]] {
				c[nodes[i]] = j
				ret += dfs2(i + 1)
			}
		}
	}
	return ret
}

var nodes []int
var parents []int
var c []int

func solve(v int) int {
	nodes = make([]int, 0)
	parents = make([]int, N)

	dfs1(v, v)

	if len(nodes) == 0 {
		return 1
	}
	if len(nodes) == 1 {
		return 3
	}

	c = make([]int, N)
	return dfs2(0)
}

var N, M int
var node [][]int
var visited []bool

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node = make([][]int, N)
	for i := 0; i < M; i++ {
		from, to := getI()-1, getI()-1
		node[from] = append(node[from], to)
		node[to] = append(node[to], from)
	}

	visited = make([]bool, N)

	ret := 1
	for i := 0; i < N; i++ {
		ret *= solve(i)
	}
	out(ret)
}
