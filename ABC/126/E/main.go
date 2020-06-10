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

var used []bool

func dfs(v, n int, node [][]int) {
	used[v] = true
	for _, e := range node[v] {
		if used[e] {
			continue
		}
		dfs(e, n, node)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	node := make([][]int, N)
	for i := 0; i < M; i++ {
		x, y, _ := getInt()-1, getInt()-1, getInt()
		node[x] = append(node[x], y)
		node[y] = append(node[y], x)
	}

	used = make([]bool, N)
	ans := 0
	for i := 0; i < N; i++ {
		if used[i] {
			continue
		}
		ans++
		// out("i", i)
		dfs(i, i, node)
	}
	out(ans)
}
