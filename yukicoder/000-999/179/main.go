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

var H, W int
var s []string
var tot int

func match(dx, dy int) bool {
	u := make([][]int, H)
	for i := 0; i < H; i++ {
		u[i] = make([]int, W)
	}
	r, b := 0, 0
	// out("------")
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			rx, ry := x, y
			bx, by := x+dx, y+dy
			if bx < 0 || bx >= W || by < 0 || by >= H {
				continue
			}
			if u[ry][rx] == 0 && s[ry][rx] == '#' {
				b++
				if s[by][bx] == '#' {
					// out("x", by, bx)
					u[by][bx] = 1
					r++
				}
			}
		}
	}
	// for y := 0; y < H; y++ {
	// 	// for x := 0; x < W; x++ {
	// 	out(u[y])
	// 	// }
	// }
	// out(r, b, tot)
	if r == b && r+b == tot {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	H, W = getInt(), getInt()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] == '#' {
				tot++
			}
		}
	}
	if tot == 0 {
		out("NO")
		return
	}

	for y := -H; y < H; y++ {
		for x := -W; x < W; x++ {
			if x == 0 && y == 0 {
				continue
			}
			// out("[x,y]------:", x, y)
			ret := match(x, y)
			// out(x, y, ret)
			if ret == true {
				out("YES")
				return
			}
		}
	}
	out("NO")
}
