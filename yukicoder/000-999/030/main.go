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

type pair struct {
	q, p int
}

var N, M int
var t map[int][]pair
var n []int
var memo [101]int
var calced [101]bool

func dfs(s int) int {
	if s == N {
		return 1
	}
	if calced[s] {
		return memo[s]
	}
	ret := 0
	for _, e := range t[s] {
		ret += dfs(e.p) * e.q
	}
	calced[s] = true
	memo[s] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M = getInt(), getInt()
	t = make(map[int][]pair)
	it := make(map[int][]pair)
	n = make([]int, N+1)

	for i := 0; i < M; i++ {
		p, q, r := getInt(), getInt(), getInt()
		t[p] = append(t[p], pair{q, r})
		it[r] = append(it[r], pair{q, p})
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for i := 1; i < N; i++ {
		if len(it[i]) == 0 {
			ret := dfs(i)
			fmt.Fprintln(wr, ret)
		} else {
			fmt.Fprintln(wr, 0)
		}
	}
}
