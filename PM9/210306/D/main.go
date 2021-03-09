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
	x, y int
}

const inf = int(1e15)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	if s[0][0] == '#' || s[H-1][W-1] == '#' {
		out(-1)
		return
	}

	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
		}
	}
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	q := make([]pos, 0)
	q = append(q, pos{0, 0})
	dist[0][0] = 0
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
			if dist[py][px] != inf {
				continue
			}
			dist[py][px] = dist[c.y][c.x] + 1
			q = append(q, pos{px, py})
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }

	if dist[H-1][W-1] == inf {
		out(-1)
		return
	}

	cnt := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				cnt++
			}
		}
	}
	ans := H*W - (dist[H-1][W-1] + 1) - cnt
	out(ans)
}
