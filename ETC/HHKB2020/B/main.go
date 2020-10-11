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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	H, W := getInt(), getInt()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}
	dx := []int{1, 0}
	dy := []int{0, 1}

	cnt := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] == '#' {
				continue
			}
			for i := 0; i < 2; i++ {
				px := x + dx[i]
				py := y + dy[i]
				if px < 0 || py < 0 || px >= W || py >= H {
					continue
				}
				if s[py][px] == '.' {
					cnt++
				}
			}
		}
	}
	out(cnt)
}
