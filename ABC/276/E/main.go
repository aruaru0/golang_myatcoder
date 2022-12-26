package main

import (
	"bufio"
	"fmt"
	"math"
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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
var c []string

type pos struct {
	x, y int
}

const inf = int(1e18)

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func calc(sx, sy int) [][]int {
	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
		}
	}
	q := make([]pos, 0)
	q = append(q, pos{sx, sy})
	dist[sy][sx] = 0
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := cur.x + dx[i]
			py := cur.y + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if c[py][px] == '#' {
				continue
			}
			if c[py][px] == 'S' {
				continue
			}
			if dist[py][px] > dist[cur.y][cur.x]+1 {
				dist[py][px] = dist[cur.y][cur.x] + 1
				q = append(q, pos{px, py})
			}
		}
	}

	return dist
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	c = make([]string, H)
	sx, sy := 0, 0
	for i := 0; i < H; i++ {
		c[i] = getS()
		for j := 0; j < W; j++ {
			if c[i][j] == 'S' {
				sy, sx = i, j
			}
		}
	}

	for i := 0; i < 4; i++ {
		cx, cy := sx+dx[i], sy+dy[i]
		if cx < 0 || cx >= W || cy < 0 || cy >= H {
			continue
		}
		if c[cy][cx] == '#' {
			continue
		}
		ret := calc(cx, cy)
		// out("-------------------")
		// for h := 0; h < H; h++ {
		// 	out(ret[h])
		// }
		for j := 0; j < 4; j++ {
			px, py := sx+dx[j], sy+dy[j]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if c[py][px] == '#' {
				continue
			}
			if px == cx && py == cy {
				continue
			}
			// out("target", py, px)
			if ret[py][px] != inf {
				out("Yes")
				return
			}
		}
	}
	out("No")
}
