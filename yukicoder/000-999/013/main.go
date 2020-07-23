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
var m [][]int

type loc struct {
	x, y int
}

var dist [][]int

func dfs(sx, sy, n int) {
	dist[sy][sx] = n
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for k := 0; k < 4; k++ {
		x := sx + dx[k]
		y := sy + dy[k]
		if x < 0 || x >= W || y < 0 || y >= H {
			continue
		}
		if m[sy][sx] != m[y][x] {
			continue
		}
		if dist[y][x] != 0 {
			continue
		}
		dfs(x, y, n+1)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	W, H = getInt(), getInt()
	m = make([][]int, H)
	for i := 0; i < H; i++ {
		m[i] = getInts(W)
	}

	for sy := 0; sy < H; sy++ {
		for sx := 0; sx < W; sx++ {
			dist = make([][]int, H)
			for i := 0; i < H; i++ {
				dist[i] = make([]int, W)
			}
			dfs(sx, sy, 1)
			// for i := 0; i < H; i++ {
			// 	out(dist[i])
			// }
			dx := []int{-1, 1, 0, 0}
			dy := []int{0, 0, -1, 1}
			for k := 0; k < 4; k++ {
				x := sx + dx[k]
				y := sy + dy[k]
				if x < 0 || x >= W || y < 0 || y >= H {
					continue
				}
				if dist[y][x] > 2 {
					out("possible")
					return
				}
			}

		}
	}
	out("impossible")
}
