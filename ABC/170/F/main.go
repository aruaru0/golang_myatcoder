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

const inf = -1

var H, W, K int
var c []string
var d [][]int

type pair struct {
	x, y int
}

func bfs(sx, sy int) {
	q := make([]pair, 0)
	d[sx][sy] = 0
	q = append(q, pair{sx, sy})
	dx := []int{1, 0, 0, -1}
	dy := []int{0, 1, -1, 0}
	for len(q) != 0 {
		px, py := q[0].x, q[0].y
		q = q[1:]
		for j := 0; j < 4; j++ {
			cx := px
			cy := py
			for k := 0; k < K; k++ {
				cx += dx[j]
				cy += dy[j]
				if cx < 0 || cx >= H || cy < 0 || cy >= W {
					break
				}
				if d[cx][cy] != inf && d[cx][cy] <= d[px][py] {
					break
				}
				if d[cx][cy] != inf {
					continue
				}
				if c[cx][cy] == '@' {
					break
				}
				d[cx][cy] = d[px][py] + 1
				q = append(q, pair{cx, cy})
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1100000)
	H, W, K = getInt(), getInt(), getInt()
	sx, sy, ex, ey := getInt()-1, getInt()-1,
		getInt()-1, getInt()-1

	c = make([]string, H)
	for i := 0; i < H; i++ {
		c[i] = getString()
	}
	d = make([][]int, H)
	for i := 0; i < H; i++ {
		d[i] = make([]int, W)
		for j := 0; j < W; j++ {
			d[i][j] = inf
		}
	}
	bfs(sx, sy)
	out(d[ex][ey])
}
