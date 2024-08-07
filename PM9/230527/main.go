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

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func dfs(x, y int) {
	for i := 0; i < 4; i++ {
		px := x + dx[i]
		py := y + dy[i]
		if px < 0 || px >= W || py < 0 || py >= H {
			continue
		}
		if c[py][px] == '#' {
			continue
		}
		if !dist[py][px] {
			dist[py][px] = true
			dfs(px, py)
		}
	}
	return
}

var H, W int
var c []string
var dist [][]bool

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	c = make([]string, H)

	sx, sy := 0, 0
	gx, gy := 0, 0
	for i := 0; i < H; i++ {
		c[i] = getS()
		for j := 0; j < W; j++ {
			if c[i][j] == 's' {
				sx, sy = j, i
			}
			if c[i][j] == 'g' {
				gx, gy = j, i
			}
		}
	}

	dist = make([][]bool, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]bool, W)
	}
	dist[sy][sx] = true
	dfs(sx, sy)

	if dist[gy][gx] {
		out("Yes")
	} else {
		out("No")
	}
}
