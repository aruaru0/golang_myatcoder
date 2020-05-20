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

const inf = 1001001001001001

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	node := make([][]int, N)
	for i := 0; i < M; i++ {
		f, t := getInt()-1, getInt()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
	}

	ans := make([]int, N)
	// cnt := make([]int, N)
	used := make([]bool, N)
	// for i := 0; i < N; i++ {
	// 	cnt[i] = inf
	// }
	s := make([]int, 0)
	s = append(s, 0)
	// cnt[0] = 0
	used[0] = true
	for len(s) > 0 {
		x := s[0]
		s = s[1:]
		for _, v := range node[x] {
			// if cnt[v] > cnt[x]+1 {
			// 	cnt[v] = cnt[x] + 1
			// }
			if used[v] == true {
				continue
			}
			used[v] = true
			ans[v] = x
			s = append(s, v)
		}
	}

	out("Yes")
	for i := 1; i < N; i++ {
		out(ans[i] + 1)
	}
}
