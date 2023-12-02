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

type pos struct {
	ax, ay, sx, sy int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H)

	sx, sy := 0, 0
	gx, gy := 0, 0
	ax, ay := 0, 0

	for i := 0; i < H; i++ {
		s[i] = getS()
		for j := 0; j < W; j++ {
			if s[i][j] == 's' {
				sx, sy = i, j
			}
			if s[i][j] == 'g' {
				gx, gy = i, j
			}
			if s[i][j] == 'a' {
				ax, ay = i, j
			}
		}
	}

	var dist [51][51][51][51]int
	const inf = int(1e18)
	for i := 0; i < 51; i++ {
		for j := 0; j < 51; j++ {
			for k := 0; k < 51; k++ {
				for l := 0; l < 51; l++ {
					dist[i][j][k][l] = inf
				}
			}
		}
	}

	dist[ax][ay][sx][sy] = 0
	q := make([]pos, 0)
	q = append(q, pos{ax, ay, sx, sy})

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		ax, ay, cx, cy := cur.ax, cur.ay, cur.sx, cur.sy
		for d := 0; d < 4; d++ {
			nx := cx + dx[d]
			ny := cy + dy[d]
			if nx < 0 || nx >= H || ny < 0 || ny >= W {
				continue
			}
			if s[nx][ny] == '#' {
				continue
			}
			if nx == ax && ny == ay {
				bx := nx + dx[d]
				by := ny + dy[d]
				if bx < 0 || bx >= H || by < 0 || by >= W {
					continue
				}
				if s[bx][by] == '#' || dist[bx][by][nx][ny] < inf {
					continue
				}
				dist[bx][by][nx][ny] = dist[ax][ay][cx][cy] + 1
				if bx == gx && by == gy {
					break
				}
				q = append(q, pos{bx, by, nx, ny})
			} else {
				if dist[ax][ay][nx][ny] < inf {
					continue
				}
				dist[ax][ay][nx][ny] = dist[ax][ay][cx][cy] + 1
				q = append(q, pos{ax, ay, nx, ny})
			}
		}
	}

	ans := inf
	for x := 0; x < H; x++ {
		for y := 0; y < W; y++ {
			ans = min(ans, dist[gx][gy][x][y])
		}
	}
	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}
}
