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

func check(cx, cy, r, x, y int) bool {
	dx := abs(x - cx)
	dy := abs(y - cy)
	// out(cx, cy, r, x, y, dx, dy)
	return r*r >= dx*dx+dy*dy
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	x1, y1, r := getInt(), getInt(), getInt()
	x2, y2, x3, y3 := getInt(), getInt(), getInt(), getInt()

	if x2 <= x1-r && x1+r <= x3 && y2 <= y1-r && y1+r <= y3 {
		out("NO")
	} else {
		out("YES")
	}

	if check(x1, y1, r, x2, y2) &&
		check(x1, y1, r, x2, y3) &&
		check(x1, y1, r, x3, y2) &&
		check(x1, y1, r, x3, y3) {
		out("NO")
	} else {
		out("YES")
	}
	// out(x1, y1, r, d, x1-d, x1+d, y1-d, y1+d)
	// out(x2, y2, x3, y3)
}
