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

func check(x, y []int, a, b, c int) (int, int) {
	ax := x[b] - x[a]
	ay := y[b] - y[a]
	bx := x[c] - x[a]
	by := y[c] - y[a]
	// out(ax, ay, bx, by)
	if ax*ax+ay*ay == bx*bx+by*by &&
		ax*bx+ay*by == 0 {
		return x[c] + ax, y[c] + ay
	}
	return -1, -1
}

func main() {
	sc.Split(bufio.ScanWords)
	x := make([]int, 3)
	y := make([]int, 3)
	for i := 0; i < 3; i++ {
		x[i] = getInt()
		y[i] = getInt()
	}
	rx, ry := check(x, y, 0, 1, 2)
	if rx != -1 {
		out(rx, ry)
		return
	}
	rx, ry = check(x, y, 1, 0, 2)
	if rx != -1 {
		out(rx, ry)
		return
	}
	rx, ry = check(x, y, 2, 0, 1)
	if rx != -1 {
		out(rx, ry)
		return
	}
	out(-1)
}
