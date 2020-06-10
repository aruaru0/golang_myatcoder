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

func main() {
	sc.Split(bufio.ScanWords)
	N, C := getInt(), getInt()
	s := make([]int, N)
	t := make([]int, N)
	c := make([]int, N)
	for i := 0; i < N; i++ {
		s[i], t[i], c[i] = getInt(), getInt(), getInt()-1
	}

	p := make([][]int, C)
	for i := 0; i < C; i++ {
		p[i] = make([]int, 100005)
	}
	for i := 0; i < N; i++ {
		// out(s[i], t[i], c[i])
		for j := s[i] - 1; j < t[i]; j++ {
			p[c[i]][j] = 1
		}
	}

	ans := 0
	for i := 0; i <= 100000; i++ {
		cnt := 0
		for c := 0; c < C; c++ {
			if p[c][i] == 1 {
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	out(ans)
}
