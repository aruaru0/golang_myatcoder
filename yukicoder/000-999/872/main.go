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

type edge struct {
	to, w int
	idx   int
}

var n int
var node [][]edge
var dist [][]int

func dfs(v, p int) int {
	ret := 1
	dist[v] = make([]int, len(node[v]))
	for i, e := range node[v] {
		if e.to == p {
			continue
		}
		x := dfs(e.to, v)
		ret += x
		dist[v][i] = x
	}
	return ret
}

func dfs2(v, p, cnt int) {
	pidx := -1
	for i, e := range node[v] {
		if e.to == p {
			pidx = i
			continue
		}
		dfs2(e.to, v, n-dist[v][i])
	}
	if pidx != -1 {
		dist[v][pidx] = cnt
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n = getI()
	node = make([][]edge, n)
	c := make([]int, n)
	for i := 0; i < n-1; i++ {
		from, to, cost := getI()-1, getI()-1, getI()
		node[from] = append(node[from], edge{to, cost, i})
		node[to] = append(node[to], edge{from, cost, i})
		c[i] = cost
	}

	dist = make([][]int, n)
	dfs(0, -1)
	dfs2(0, -1, 0)
	// out(dist)

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		for j, e := range node[i] {
			// out(e, dist[i][j])
			if m[e.idx] == 0 {
				m[e.idx] = dist[i][j]
			} else {
				m[e.idx] *= dist[i][j]
			}
		}
	}
	// out(m)
	ans := 0
	for i, e := range m {
		ans += e * c[i]
	}
	out(ans * 2)
}
