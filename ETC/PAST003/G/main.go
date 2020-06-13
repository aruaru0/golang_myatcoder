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

// テキスト "#"が壁のやつを幅優先探索(0,0)の距離
type queue struct {
	x, y int
}

// sx,syからの距離を探索 結果は距離の配列
func bfs(sx, sy, H, W int, s [][]byte) [][]int {
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

	dx := []int{+1, 0, -1, +1, -1, 0}
	dy := []int{+1, +1, +1, 0, 0, -1}

	dist[sx][sy] = 0
	for len(q) != 0 {
		x := q[0].x
		y := q[0].y
		q = q[1:]
		for k := 0; k < 6; k++ {
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

	N, X, Y := getInt(), getInt(), getInt()
	s := make([][]byte, 500)
	for i := 0; i < 500; i++ {
		s[i] = make([]byte, 500)
	}
	for i := 0; i < N; i++ {
		x, y := getInt()+250, getInt()+250
		s[y][x] = '#'
	}
	dist := bfs(250, 250, 500, 500, s)

	// for i := 200; i < 204; i++ {
	// 	out(dist[i][200:204])
	// }
	out(dist[Y+250][X+250])
}
