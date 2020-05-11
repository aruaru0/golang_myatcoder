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

// 解説動画から。ダブリングを利用（覚える！！）
// BIT :
const BIT = 60

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	next := make([][]int, BIT)
	for i := 0; i < BIT; i++ {
		next[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		next[0][i] = getInt() - 1
	}

	for i := 0; i < BIT-1; i++ {
		for j := 0; j < N; j++ {
			next[i+1][j] = next[i][next[i][j]]
		}
	}

	pos := 0
	for i := BIT - 1; i >= 0; i-- {
		l := 1 << uint(i)
		if l <= K {
			// out(l, K, pos)
			K -= l
			pos = next[i][pos]
		}
	}
	out(pos + 1)
}
