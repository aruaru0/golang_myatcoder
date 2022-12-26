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
var w []int
var flg bool
var one map[int]bool

func dfs(v int) int {
	if used[v] {
		return 0
	}
	if one[v] {
		flg = true
	}
	used[v] = true
	ret := 0
	if w[v] == 0 {
		ret++
	}
	for _, e := range node[v] {
		ret += dfs(e)
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	d := getInts(N)
	w = getInts(N)
	node = make([][]int, N)
	one = make(map[int]bool)
	for i := 0; i < N; i++ {
		f := (i + d[i]) % N
		t := (i - d[i]) % N
		if t < 0 {
			t += N
		}
		// out(i, d[i], f, t)
		if f == t {
			one[f] = true
			continue
		}
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
	}
	// out(d)
	// out(w)
	// out(node)
	used = make([]bool, N)
	for i := 0; i < N; i++ {
		flg = false
		ret := dfs(i)
		// out(ret, flg, used)
		if ret%2 != 0 && flg == false {
			out("No")
			return
		}
	}
	out("Yes")
}
