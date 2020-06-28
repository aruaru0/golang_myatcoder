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

func check(Y, X, h, w int, s []string) int {
	ret := 0
	dx := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	for i := 0; i < 9; i++ {
		x := X + dx[i]
		y := Y + dy[i]
		if x < 0 || x >= w || y < 0 || y >= h {
			continue
		}
		if s[y][x] == '#' {
			ret++
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	H, W := getInt(), getInt()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				fmt.Print(string('#'))
			} else {
				n := check(i, j, H, W, s)
				fmt.Print(string(n + '0'))
			}
		}
		out()
	}
}
