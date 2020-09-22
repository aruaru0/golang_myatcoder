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

const inf = int(1e10)

// メモ化再帰
var memo [1010][1010]int

func rec(s, t string) int {
	if len(s) == 0 {
		return len(t)
	}
	if len(t) == 0 {
		return len(s)
	}
	if memo[len(s)][len(t)] != -1 {
		return memo[len(s)][len(t)]
	}

	cost := 1
	if s[0] == t[0] {
		cost = 0
	}
	l0 := rec(s[1:], t[1:]) + cost // change
	l1 := rec(s, t[1:]) + 1        // insert
	l2 := rec(s[1:], t) + 1        // delete
	l := min(l0, min(l1, l2))
	memo[len(s)][len(t)] = l
	return l
}

var n, m int

func main() {
	sc.Split(bufio.ScanWords)
	n, m = getInt(), getInt()
	s, t := getString(), getString()

	for i := 0; i < 1010; i++ {
		for j := 0; j < 1010; j++ {
			memo[i][j] = -1
		}

	}

	out(rec(s, t))
}
