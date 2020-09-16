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

type pos struct {
	x, y int
}

const inf = int(1e15)

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	p := make([]pos, N)
	for i := 0; i < N; i++ {
		x, y := getInt(), getInt()
		p[i] = pos{x - y, x + y}
	}
	ma0 := -inf
	mi0 := inf
	ma1 := -inf
	mi1 := inf
	for i := 0; i < N; i++ {
		ma0 = max(ma0, p[i].x)
		ma1 = max(ma1, p[i].y)
		mi0 = min(mi0, p[i].x)
		mi1 = min(mi1, p[i].y)
	}
	ans := max(ma0-mi0, ma1-mi1)
	out(ans)
}
