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

func next(x, y, dx, dy int) (int, int, int, int) {

	if x == 0 && dx == -1 {
		dx = 1
	}
	if x == 8 && dx == 1 {
		dx = -1
	}
	if y == 0 && dy == -1 {
		dy = 1
	}
	if y == 8 && dy == 1 {
		dy = -1
	}

	x += dx
	y += dy

	// out(x, y, dx, dy)
	return x, y, dx, dy
}

func main() {
	sc.Split(bufio.ScanWords)
	x, y, W := getInt()-1, getInt()-1, getString()
	c := make([]string, 9)
	for i := 0; i < 9; i++ {
		c[i] = getString()
	}

	var dx, dy int
	switch W {
	case "U":
		dx, dy = 0, -1
	case "D":
		dx, dy = 0, 1
	case "L":
		dx, dy = -1, 0
	case "R":
		dx, dy = 1, 0
	case "RU":
		dx, dy = 1, -1
	case "RD":
		dx, dy = 1, 1
	case "LU":
		dx, dy = -1, -1
	case "LD":
		dx, dy = -1, 1
	}

	pw := make([]byte, 4)
	pw[0] = c[y][x]
	for i := 1; i < 4; i++ {
		x, y, dx, dy = next(x, y, dx, dy)
		pw[i] = c[y][x]
	}
	out(string(pw))
}
