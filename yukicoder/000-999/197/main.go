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

type dice [3]byte

func solve(s0, s1 string, n int) {
	a := make(map[dice]int)
	a[[3]byte{s0[0], s0[1], s0[2]}]++
	if n > 1 {
		if n%2 == 0 {
			n = 2
		} else {
			n = 3
		}
	}
	for i := 0; i < n; i++ {
		b := make(map[dice]int)
		for e := range a {
			b[[3]byte{e[1], e[0], e[2]}]++
			b[[3]byte{e[0], e[2], e[1]}]++
		}
		a = b
	}
	if a[[3]byte{s1[0], s1[1], s1[2]}] == 0 {
		out("SUCCESS")
		return
	}
	out("FAILURE")
}

func main() {
	sc.Split(bufio.ScanWords)
	s0, N, s1 := getString(), getInt(), getString()
	solve(s0, s1, N)
}
