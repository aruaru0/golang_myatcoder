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

// N:
var N int
var node [][]int
var dist1 []int
var dist2 []int
var u, v int

func dfs1(x, p, cnt int) {
	if len(node[x]) == 1 {
		dist1[x] = cnt
	}
	for _, e := range node[x] {
		if e == p {
			continue
		}
		dfs1(e, x, cnt+1)
	}
}

func dfs2(x, p, cnt int) {
	if len(node[x]) == 1 {
		dist2[x] = cnt
	}
	for _, e := range node[x] {
		if e == p {
			continue
		}
		dfs2(e, x, cnt+1)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N, u, v = getInt(), getInt()-1, getInt()-1
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		f, t := getInt()-1, getInt()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
	}

	dist1 = make([]int, N)
	dist2 = make([]int, N)
	dfs1(v, -1, 0)
	dfs2(u, -1, 0)
	// out(dist1)
	// out(dist2)

	ans := 0
	for i := 0; i < N; i++ {
		if dist1[i] > dist2[i] {
			diff := dist1[i] - dist2[i]
			// out(dist1[i], dist2[i], diff)
			ans = max(ans, dist2[i]+diff-1)

		}
	}
	out(ans)
}
