package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

type pos struct {
	y, x int
}

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	a := make([]string, H)
	w := make([][]pos, 26)
	var s, g pos
	for i := 0; i < H; i++ {
		a[i] = getS()
		for j := 0; j < W; j++ {
			if a[i][j] >= 'a' && a[i][j] <= 'z' {
				p := int(a[i][j] - 'a')
				w[p] = append(w[p], pos{i, j})
			}
			if a[i][j] == 'S' {
				s = pos{i, j}
			}
			if a[i][j] == 'G' {
				g = pos{i, j}
			}
		}
	}

	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
		}
	}
	q := make([]pos, 0)
	q = append(q, s)
	dist[s.y][s.x] = 0
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	used := make([]bool, 26)
	for len(q) != 0 {
		cur := q[0]
		cx := cur.x
		cy := cur.y
		q = q[1:]
		if a[cy][cx] >= 'a' && a[cy][cx] <= 'z' {
			p := int(a[cy][cx] - 'a')
			if used[p] != true {
				used[p] = true
				for _, e := range w[p] {
					px := e.x
					py := e.y
					if dist[py][px] > dist[cy][cx]+1 {
						dist[py][px] = dist[cy][cx] + 1
						q = append(q, pos{py, px})
					}
				}
			}
		}
		for i := 0; i < 4; i++ {
			px := cx + dx[i]
			py := cy + dy[i]
			if px < 0 || py < 0 || px >= W || py >= H {
				continue
			}
			if a[py][px] == '#' {
				continue
			}
			if dist[py][px] > dist[cy][cx]+1 {
				dist[py][px] = dist[cy][cx] + 1
				q = append(q, pos{py, px})
			}
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }

	if dist[g.y][g.x] == inf {
		out(-1)
		return
	}

	out(dist[g.y][g.x])
}
