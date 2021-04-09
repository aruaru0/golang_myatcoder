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

func toIdx(x, y int) int {
	return y*W + x
}

func toXY(idx int) (int, int) {
	return idx % W, idx / W
}

func rec(idx, a, b int, d [][]int) int {
	// out(idx, a, b, d)
	if a == A && b == B {
		return 1
	}
	if H*W == idx {
		return 0
	}

	x, y := toXY(idx)
	ret := 0

	// out(x, y)
	if d[y][x] != 0 {
		ret += rec(idx+1, a, b, d)
	} else {
		if b < B {
			d[y][x] = -(b + 1)
			ret += rec(idx+1, a, b+1, d)
			d[y][x] = 0
		}

		if a < A {
			if x < W-1 && d[y][x+1] == 0 {
				d[y][x] = (a + 1)
				d[y][x+1] = (a + 1)
				ret += rec(idx+1, a+1, b, d)
				d[y][x+1] = 0
				d[y][x] = 0
			}
			if y < H-1 && d[y+1][x] == 0 {
				d[y][x] = (a + 1)
				d[y+1][x] = (a + 1)
				ret += rec(idx+1, a+1, b, d)
				d[y+1][x] = 0
				d[y][x] = 0
			}
		}
	}
	return ret
}

var H, W int
var A, B int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	A, B = getI(), getI()

	d := make([][]int, H)
	for i := 0; i < H; i++ {
		d[i] = make([]int, W)
	}
	out(rec(0, 0, 0, d))
}
