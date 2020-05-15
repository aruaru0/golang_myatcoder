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

type robot struct {
	x, d int
}

type robots []robot

func (p robots) Len() int {
	return len(p)
}

func (p robots) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p robots) Less(i, j int) bool {
	return p[i].x < p[j].x
}

const mod = 998244353

func dfs(n int, e [][]int) int {
	ret := 1
	for _, v := range e[n] {
		ret *= dfs(v, e)
		ret %= mod
	}
	ret = (ret + 1) % mod
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	r := make(robots, N)
	for i := 0; i < N; i++ {
		x, d := getInt(), getInt()
		r[i] = robot{x, d}
	}
	sort.Sort(r)
	// out(r)

	e := make([][]int, N)
	s := make([][2]int, 0)
	for i := N - 1; i >= 0; i-- {
		x := r[i].x
		d := r[i].d
		for len(s) > 0 && s[len(s)-1][0] < x+d {
			e[i] = append(e[i], s[len(s)-1][1])
			s = s[:len(s)-1]
		}
		s = append(s, [2]int{x, i})
	}
	// out(s)
	// out(e)
	ans := 1
	for _, v := range s {
		ans *= dfs(v[1], e)
		ans %= mod
	}
	out(ans)
}
