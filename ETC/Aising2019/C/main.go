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
var u [][]int

type pos struct {
	x, y int
	c    bool
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func bsf(sx, sy, c int) (int, int) {
	q := make([]pos, 0)
	q = append(q, pos{sx, sy, true})
	u[sy][sx] = c
	b := 1
	w := 0
	for len(q) != 0 {
		cx := q[0].x
		cy := q[0].y
		black := q[0].c
		q = q[1:]
		for i := 0; i < 4; i++ {
			x := cx + dx[i]
			y := cy + dy[i]
			if x < 0 || x >= W || y < 0 || y >= H {
				continue
			}
			if black == true && s[y][x] == '#' {
				continue
			}
			if black == false && s[y][x] == '.' {
				continue
			}
			if u[y][x] != 0 {
				continue
			}
			u[y][x] = c
			if s[y][x] == '#' {
				b++
			} else {
				w++
			}
			q = append(q, pos{x, y, !black})
		}
	}
	return b, w
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	H, W = getInt(), getInt()
	s = make([]string, H)
	u = make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
		u[i] = make([]int, W)
	}

	cnt := 1
	ans := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] == '#' && u[y][x] == 0 {
				b, w := bsf(x, y, cnt)
				ans += b * w
				cnt++
			}
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(u[i])
	// }
	out(ans)
}
