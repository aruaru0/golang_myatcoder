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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

func dfs(v int) int {
	cnt := 1
	used[v] = true
	for _, e := range node[v] {
		if used[e] {
			continue
		}
		cnt += dfs(e)
	}
	return cnt
}

var node [][]int
var used []bool

func main() {
	sc.Split(bufio.ScanWords)
	N, P := getInt(), getInt()

	node = make([][]int, N+1)
	n := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		if n[i] {
			continue
		}
		for j := i + i; j <= N; j += i {
			n[j] = true
			node[i] = append(node[i], j)
			node[j] = append(node[j], i)
		}
	}
	used = make([]bool, N+1)
	// out(node, P)
	out(dfs(P))
}
