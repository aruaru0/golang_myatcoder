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

var node [100100][]edge
var cost [100100]int
var used [100100]bool

func dfs(n, c int) bool {
	used[n] = true
	cost[n] = c
	for _, e := range node[n] {
		if used[e.to] == true {
			if c+e.cost != cost[e.to] {
				return false
			}
			continue
		}
		flg := dfs(e.to, c+e.cost)
		if flg == false {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()

	for i := 0; i < M; i++ {
		l, r, d := getInt()-1, getInt()-1, getInt()
		node[l] = append(node[l], edge{r, d})
		node[r] = append(node[r], edge{l, -d})
	}

	// out(node[:10], N, M)

	ans := true
	for i := 0; i < N; i++ {
		flg := true
		if used[i] == false {
			flg = dfs(i, 0)
			// out(flg)
		}
		if flg == false {
			ans = false
			break
		}
	}
	if ans {
		out("Yes")
	} else {
		out("No")
	}
}
