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

const inf = int(1e8)

type pos struct {
	x, y int
}

func f(sx, sy, ex, ey int) int {
	if s[sy][sx] == '#' || s[ey][ex] == '#' {
		return -1
	}
	d := make([][]int, H)
	for i := 0; i < H; i++ {
		d[i] = make([]int, W)
		for j := 0; j < W; j++ {
			d[i][j] = inf
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	q := make([]pos, 0)
	d[sy][sx] = 0
	q = append(q, pos{sx, sy})
	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := c.x + dx[i]
			py := c.y + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if s[py][px] == '#' {
				continue
			}
			if d[py][px] > d[c.y][c.x]+1 {
				d[py][px] = d[c.y][c.x] + 1
				q = append(q, pos{px, py})
			}
		}
	}

	if d[ey][ex] == inf {
		return -1
	}
	return d[ey][ex]
}

var H, W int
var s []string

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	ans := 0
	for sx := 0; sx < W; sx++ {
		for sy := 0; sy < H; sy++ {
			for ex := 0; ex < W; ex++ {
				for ey := 0; ey < H; ey++ {
					ans = max(ans, f(sx, sy, ex, ey))
				}
			}
		}
	}
	out(ans)
}
