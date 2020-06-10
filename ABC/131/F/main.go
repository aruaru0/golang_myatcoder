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

var size = 100005
var node [][]int
var used []bool
var cnt []int

func dfs(v int) {
	if used[v] {
		return
	}
	used[v] = true
	cnt[v/size]++
	for _, e := range node[v] {
		dfs(e)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	var n int
	n = getInt()
	node = make([][]int, size*2)
	for i := 0; i < n; i++ {
		x, y := getInt(), getInt()
		y += size
		node[x] = append(node[x], y)
		node[y] = append(node[y], x)
	}
	used = make([]bool, size*2)
	ans := 0
	for i := 0; i < size*2; i++ {
		if used[i] {
			continue
		}
		cnt = make([]int, 2)
		dfs(i)
		// if cnt[0]*cnt[1] != 0 {
		// 	out(i, cnt)
		// }
		ans += cnt[0] * cnt[1]
	}
	ans -= n
	out(ans)
}
