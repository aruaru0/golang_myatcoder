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

// テキスト "#"が壁のやつを幅優先探索(0,0)の距離
type queue struct {
	x, y int
}

func bfs(sx, sy, H, W int, s []string) [][]int {
	dist := make([][]int, H)
	// 初期化
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = -1
		}
	}
	q := make([]queue, 0)
	q = append(q, queue{sx, sy})

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	dist[sx][sy] = 0
	for len(q) != 0 {
		x := q[0].x
		y := q[0].y
		q = q[1:]
		for k := 0; k < 4; k++ {
			yy := y + dy[k]
			xx := x + dx[k]
			if yy < 0 || xx < 0 || yy >= H || xx >= W {
				continue
			}
			if s[yy][xx] == '#' {
				continue
			}
			if dist[yy][xx] != -1 {
				continue
			}
			dist[yy][xx] = dist[y][x] + 1
			q = append(q, queue{xx, yy})
		}

	}

	return dist
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	R, C := getInt(), getInt()
	sx, sy := getInt()-1, getInt()-1
	ex, ey := getInt()-1, getInt()-1
	s := make([]string, R)
	for i := 0; i < R; i++ {
		s[i] = getString()
	}

	d := bfs(sx, sy, R, C, s)
	out(d[ex][ey])
}
