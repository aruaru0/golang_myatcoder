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

var n int
var l, s []int
var done []int
var used []int

func dfs(cur, tar int) (int, bool) {
	if done[s[cur]] == 1 {
		return -1, false
	}
	if cur == tar {
		return tar, true
	}
	if used[s[cur]] == 1 {
		return -1, false
	}
	used[cur] = 1
	x, ok := dfs(s[cur], tar)
	if !ok {
		return -1, false
	}
	if l[cur] < l[x] {
		return cur, true
	}
	return x, true
}

func main() {
	sc.Split(bufio.ScanWords)
	n = getInt()
	l = make([]int, n)
	s = make([]int, n)
	for i := 0; i < n; i++ {
		l[i], s[i] = getInt(), getInt()-1
		l[i] *= 2
	}

	done = make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		if done[i] == 0 {
			used = make([]int, n)
			x, ok := dfs(s[i], i)
			if ok {
				ans += l[x]
				done[x] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if done[s[j]] == 1 && done[j] == 0 {
				done[j] = 1
				ans += l[j] / 2
			}
		}
	}

	fmt.Printf("%1.1f\n", (float64(ans) * 0.5))
}
