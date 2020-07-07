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

var H, W, K int
var s []string

func solve(h, w int) int {
	ret := 0
	// fmt.Printf("%b %b\n", h, w)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if h>>i%2 == 0 && w>>j%2 == 0 {
				if s[i][j] == '#' {
					ret++
				}
			}
		}
	}
	// out(ret)
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	H, W, K = getInt(), getInt(), getInt()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}

	h := 1 << H
	w := 1 << W

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			k := solve(i, j)
			if k == K {
				ans++
			}
		}
	}
	out(ans)
}
