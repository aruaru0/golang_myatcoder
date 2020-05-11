package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
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
	to, cost int
}

var node [][]edge

var res []int

func dfs(v, prev, cost int) {
	if prev != -1 {
		res[v] = res[prev] + cost
	}
	for _, e := range node[v] {
		if e.to == prev {
			continue
		}
		dfs(e.to, v, e.cost)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	node = make([][]edge, N)
	for i := 0; i < N-1; i++ {
		u, v, w := getInt()-1, getInt()-1, getInt()
		node[u] = append(node[u], edge{v, w})
		node[v] = append(node[v], edge{u, w})
	}

	res = make([]int, N)
	dfs(0, -1, 0)
	// out(res)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < N; i++ {
		fmt.Fprintln(w, res[i]%2)
	}
}
