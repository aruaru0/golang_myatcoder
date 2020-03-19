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

type queue struct {
	x, y int
}

const inf = 100000

func bfs(sx, sy, H, W int, s []string) ([][]int, int) {
	if s[sy][sx] == '#' {
		return nil, 0
	}

	dist := make([][]int, H)
	for y := 0; y < H; y++ {
		dist[y] = make([]int, W)
		for x := 0; x < W; x++ {
			dist[y][x] = -1
		}
	}

	q := make([]queue, 0)
	q = append(q, queue{sx, sy})

	dx := []int{0, -1, 1, 0}
	dy := []int{-1, 0, 0, 1}

	ma := -1
	dist[sy][sx] = 0
	for len(q) != 0 {
		p := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			xx := p.x + dx[i]
			yy := p.y + dy[i]
			if xx < 0 || yy < 0 || xx >= W || yy >= H {
				continue
			}
			if s[yy][xx] == '.' && dist[yy][xx] == -1 {
				dist[yy][xx] = dist[p.y][p.x] + 1
				ma = max(ma, dist[yy][xx])
				q = append(q, queue{xx, yy})
			}
		}
	}
	return dist, ma
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W := getInt(), getInt()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}

	ans := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			_, ma := bfs(x, y, H, W, s)
			ans = max(ans, ma)
		}
	}
	out(ans)

}
