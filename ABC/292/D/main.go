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

var node [][]int
var group []int

func dfs(cur, idx int) {
	group[cur] = idx
	for _, e := range node[cur] {
		if group[e] != -1 {
			continue
		}
		dfs(e, idx)
	}
}

type edge struct {
	u, v int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	node = make([][]int, N)
	ed := make([]edge, 0)
	for i := 0; i < M; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
		ed = append(ed, edge{u, v})
	}

	group = make([]int, N)
	for i := 0; i < N; i++ {
		group[i] = -1
	}

	idx := 0
	for i := 0; i < N; i++ {
		if group[i] == -1 {
			dfs(i, idx)
			idx++
		}
	}

	nodes := make([]int, idx)
	for _, e := range group {
		nodes[e]++
	}

	edges := make([]int, idx)
	for _, e := range ed {
		edges[group[e.u]]++
	}

	ans := true
	for i := 0; i < idx; i++ {
		if nodes[i] != edges[i] {
			ans = false
		}
	}
	// out(ed)
	// out(group)
	// out(nodes)
	// out(edges)

	if ans {
		out("Yes")
	} else {
		out("No")
	}
}
