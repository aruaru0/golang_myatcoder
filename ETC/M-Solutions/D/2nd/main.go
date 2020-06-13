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

var col []int
var ecol []int
var used int

func dfs(v, p int, node [][]int) {
	for _, e := range node[v] {
		if e == p {
			continue
		}
		dfs(e, v, node)
	}
	ecol[v] = col[used]
	used++
}

type pair struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	node := make([][]int, N)
	edge := make([]pair, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
		edge[i] = pair{f, t}
	}
	col = make([]int, N)
	ecol = make([]int, N)
	for i := 0; i < N; i++ {
		col[i] = getInt()
	}
	sort.Ints(col)
	dfs(0, -1, node)
	//out(ecol)
	ans := 0
	for i := 0; i < N-1; i++ {
		c0 := ecol[edge[i].x]
		c1 := ecol[edge[i].y]
		ans += min(c0, c1)
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, ans)
	for i := 0; i < N; i++ {
		fmt.Fprint(w, ecol[i], " ")
	}
	fmt.Fprintln(w)
}
