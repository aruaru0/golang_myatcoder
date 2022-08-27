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

var H, W, A, B int
var N int
var cnt int

func set_pos(s, x, y int) int {
	pos := y*W + x
	return s | (1 << pos)
}

func get_pos(s, x, y int) int {
	pos := y*W + x
	return (s >> pos) % 2
}

func rec(x, y, a, b, s int) {
	if s == 1<<(H*W)-1 {
		if a == A && b == B {
			cnt++
		}
		return
	}

	if get_pos(s, x, y) == 1 {
		x++
		if x >= W {
			x = 0
			y++
		}
		rec(x, y, a, b, s)
		return
	}

	s = set_pos(s, x, y)
	// Bを配置
	nx, ny := x+1, y
	if nx >= W {
		nx = 0
		ny++
	}
	rec(nx, ny, a, b+1, s)

	// 横にAを配置
	if x+1 < W {
		rec(nx, ny, a+1, b, set_pos(s, x+1, y))
	}
	// 縦にAを配置
	if y+1 < H {
		rec(nx, ny, a+1, b, set_pos(s, x, y+1))
	}

}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, A, B = getI(), getI(), getI(), getI()
	N = 1 << (H * W)

	cnt = 0
	rec(0, 0, 0, 0, 0)
	out(cnt)
}
