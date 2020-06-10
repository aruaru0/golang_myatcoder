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

type pos struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([][]int, 2)
	d := make([][]int, 2)
	for i := 0; i < 2; i++ {
		a[i] = make([]int, N)
		d[i] = make([]int, N)
		for j := 0; j < N; j++ {
			a[i][j] = getInt()
		}
	}

	s := make([]pos, 0)
	s = append(s, pos{0, 0})
	for len(s) != 0 {
		x, y := s[0].x, s[0].y
		s = s[1:]
		if x == 0 && y == 0 {
			d[y][x] = a[y][x]
		} else if y == 0 {
			d[y][x] = a[y][x] + d[y][x-1]
		} else if x == 0 {
			d[y][x] = a[y][x] + d[y-1][x]
		} else {
			d[y][x] = a[y][x] + max(d[y][x-1], d[y-1][x])
		}
		if x+1 != N {
			s = append(s, pos{x + 1, y})
		}
		if y+1 != 2 {
			s = append(s, pos{x, y + 1})
		}
	}
	out(d[1][N-1])
}
