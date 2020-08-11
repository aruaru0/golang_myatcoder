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

var node [][]int
var used []bool
var N, M int

func dfs(cur, prev int) int {
	used[cur] = true
	ret := 0
	for _, e := range node[cur] {
		if e == prev {
			continue
		}
		if used[e] == true {
			return -1
		}
		ret = min(ret, dfs(e, cur))
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M = getInt(), getInt()
	node = make([][]int, N)
	for i := 0; i < M; i++ {
		f, t := getInt()-1, getInt()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
	}
	used = make([]bool, N)
	ans := 0
	for i := 0; i < N; i++ {
		if used[i] != true {
			ret := dfs(i, -1)
			if ret == 0 {
				ans++
			}
		}
	}
	out(ans)
}
