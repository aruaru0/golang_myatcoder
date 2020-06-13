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

func main() {
	sc.Split(bufio.ScanWords)
	H, W, _ := getInt(), getInt(), getInt()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}
	t := make([][]int, H)
	h := make([]int, H)
	cnt := 1
	for i := 0; i < H; i++ {
		t[i] = make([]int, W)
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				t[i][j] = cnt
				cnt++
				h[i]++
			}
		}
	}

	for i := 0; i < H; i++ {
		for j := 1; j < W; j++ {
			if t[i][j] == 0 {
				t[i][j] =
					t[i][j-1]
			}
		}
		for j := W - 2; j >= 0; j-- {
			if t[i][j] == 0 {
				t[i][j] = t[i][j+1]
			}
		}
	}
	for i := 1; i < H; i++ {
		if h[i] == 0 {
			for j := 0; j < W; j++ {
				t[i][j] = t[i-1][j]
			}
		}
	}

	for i := H - 2; i >= 0; i-- {
		if h[i] == 0 {
			for j := 0; j < W; j++ {
				t[i][j] = t[i+1][j]
			}
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Print(t[i][j], " ")
		}
		fmt.Println()
	}
}
