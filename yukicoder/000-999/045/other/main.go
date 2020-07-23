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

var N int
var v []int
var memo [][2]int

func rec(x, p int) int {
	if x == N {
		return 0
	}
	if memo[x][p] != -1 {
		return memo[x][p]
	}
	ret := 0
	ret = max(ret, rec(x+1, 0))
	if p == 0 {
		ret = max(ret, rec(x+1, 1)+v[x])
	}
	memo[x][p] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	v = getInts(N)
	memo = make([][2]int, N)
	for i := 0; i < N; i++ {
		memo[i][0] = -1
		memo[i][1] = -1
	}
	out(rec(0, 0))
}
