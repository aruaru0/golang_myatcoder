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
	sc.Buffer([]byte{}, 1000000)
	W, H, N := getInt(), getInt(), getInt()
	X := make([]int, W)
	Y := make([]int, H)
	for i := 0; i < N; i++ {
		x, y, a := getInt(), getInt(), getInt()
		switch a {
		case 1:
			for j := 0; j < x; j++ {
				X[j] = 1
			}
		case 2:
			for j := x; j < W; j++ {
				X[j] = 1
			}
		case 3:
			for j := 0; j < y; j++ {
				Y[j] = 1
			}
		case 4:
			for j := y; j < H; j++ {
				Y[j] = 1
			}
		}
	}
	ans := 0

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if X[j] == 0 && Y[i] == 0 {
				ans++
			}
		}
	}
	out(ans)
}
