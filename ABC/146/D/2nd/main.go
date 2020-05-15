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
	to, idx int
}

var ma int
var node [][]edge
var col []int

func dfs(e, p, c int) {
	color := 1
	for _, v := range node[e] {
		if v.to == p {
			continue
		}
		if color == c {
			color++
		}
		col[v.idx] = color
		ma = max(ma, color)
		dfs(v.to, e, color)
		color++
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	node = make([][]edge, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		node[f] = append(node[f], edge{t, i})
		node[t] = append(node[t], edge{f, i})
	}
	col = make([]int, N-1)
	dfs(0, -1, 0)

	out(ma)
	for _, v := range col {
		out(v)
	}
}
