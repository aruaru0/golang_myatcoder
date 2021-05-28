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

const inf = int(1e10)

type pos struct {
	x, y int
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func bfs(sx, sy, s int) []pos {
	q := []pos{pos{sx, sy}}
	dist[sy][sx] = s
	ret := make([]pos, 0)
	for len(q) != 0 {
		cur := q[0]
		if dist[cur.y][cur.x] <= K {
			ret = append(ret, cur)
		}
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := cur.x + dx[i]
			py := cur.y + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if a[py][px] == '#' {
				continue
			}
			if dist[py][px] != inf {
				continue
			}
			if dist[cur.y][cur.x]+1 > K {
				continue
			}
			dist[py][px] = dist[cur.y][cur.x] + 1
			q = append(q, pos{px, py})
		}
	}
	return ret
}

var H, W, K int
var a [][]byte
var dist [][]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, K = getI(), getI(), getI()
	a = make([][]byte, H)
	for i := 0; i < H; i++ {
		a[i] = []byte(getS())
	}

	dist = make([][]int, H)
	sx, sy := 0, 0
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
			if a[i][j] == 'S' {
				sx, sy = j, i
			}
		}
	}

	if sx == 0 || sy == 0 || sx == W-1 || sy == H-1 {
		out(0)
		return
	}

	ret := bfs(sx, sy, 0)
	minL := inf
	for _, e := range ret {
		minL = nmin(minL,
			(e.x+K-1)/K,
			(e.y+K-1)/K,
			(W-1-e.x+K-1)/K,
			(H-1-e.y+K-1)/K)
	}
	out(minL + 1)
}
