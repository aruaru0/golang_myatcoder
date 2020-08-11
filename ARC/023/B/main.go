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

const inf = int(1e9)

type xy struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	R, C, D := getInt(), getInt(), getInt()
	a := make([][]int, R)
	// b := make([][]int, R)
	for i := 0; i < R; i++ {
		a[i] = getInts(C)
		// b[i] = make([]int, C)
		// for j := 0; j < C; j++ {
		// 	b[i][j] = inf
		// }
	}

	ans := 0
	for y := 0; y < R; y++ {
		for x := 0; x < C; x++ {
			if x+y > D {
				continue
			}
			if D%2 == 0 && (y+x)%2 == 0 {
				ans = max(ans, a[y][x])
			}
			if D%2 == 1 && (y+x)%2 == 1 {
				ans = max(ans, a[y][x])
			}
		}
	}
	out(ans)
}
