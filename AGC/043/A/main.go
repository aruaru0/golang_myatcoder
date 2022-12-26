package main

import (
	"bufio"
	"fmt"
	"os"
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

// テキスト "#"が壁のやつを幅優先探索(0,0)の距離
type queue struct {
	x, y int
}

const inf = 10000

func bfs(H, W int, s []string) [][]int {
	dist := make([][]int, H)
	// 初期化
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = -1
		}
	}

	if s[0][0] == '#' {
		dist[0][0] = 1
	} else {
		dist[0][0] = 0
	}

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if y == 0 && x == 0 {
				continue
			} else if y == 0 {
				dist[y][x] = dist[y][x-1]
				if s[y][x-1] == '.' && s[y][x] == '#' {
					dist[y][x]++
				}
			} else if x == 0 {
				dist[y][x] = dist[y-1][x]
				if s[y-1][x] == '.' && s[y][x] == '#' {
					dist[y][x]++
				}
			} else {
				a0 := dist[y][x-1]
				if s[y][x-1] == '.' && s[y][x] == '#' {
					a0++
				}

				a1 := dist[y-1][x]
				if s[y-1][x] == '.' && s[y][x] == '#' {
					a1++
				}
				dist[y][x] = min(a0, a1)
			}
		}
	}

	return dist
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W := getInt(), getInt()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}
	ret := bfs(H, W, s)

	out(ret[H-1][W-1])
}
