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

var H, W int
var s []string
var used [][4]bool
var tot int

var dx []int = []int{-1, 1, 0, 0}
var dy []int = []int{0, 0, -1, 1}

type pos struct {
	x, y int
}

func rec(y, x, n int, v []pos) bool {
	if n == tot {
		out(len(v))
		for _, e := range v {
			out(e.y+1, e.x+1)
		}
		wr.Flush()
		os.Exit(0)
		return true
	}
	if used[y][x] == true {
		return false
	}
	used[y][x] = true
	ret := false
	for i := 0; i < 4; i++ {
		px := x + dx[i]
		py := y + dy[i]
		if px < 0 || px >= W || py < 0 || py >= H {
			continue
		}
		if s[py][px] == '.' {
			continue
		}
		if used[py][px] == true {
			continue
		}
		ret = ret || rec(py, px, n+1, append(v, pos{px, py}))
	}
	used[y][x] = false
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	s = make([]string, H)
	tot = 0
	for i := 0; i < H; i++ {
		s[i] = getS()
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				tot++
			}
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			used = make([][4]bool, H)
			if s[i][j] == '#' {
				rec(i, j, 1, []pos{{j, i}})
			}
		}
	}
	out("MISS")
}
