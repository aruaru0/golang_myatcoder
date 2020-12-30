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

type warp struct {
	x, y, c int
	f       int
}

var gx, gy, N, F int
var w []warp

var memo [51][101][101]int

func rec(n, x, y int) int {
	if n == len(w) {
		return F * (abs(gx-x) + abs(gy-y))
	}
	if memo[n][x][y] != 0 {
		return memo[n][x][y]
	}
	ret := rec(n+1, x, y)
	X := x + w[n].x
	Y := y + w[n].y
	C := w[n].c
	if X <= gx && Y <= gy {
		ret = min(ret, rec(n+1, X, Y)+C)
	}
	memo[n][x][y] = ret
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	gx, gy, N, F = getI(), getI(), getI(), getI()
	w = make([]warp, 0)
	for i := 0; i < N; i++ {
		x, y, c := getI(), getI(), getI()
		if (x+y)*F <= c {
			continue
		}
		w = append(w, warp{x, y, c, (x + y) * F})
	}
	// sort.Slice(w, func(i, j int) bool {
	// 	return w[i].f*w[j].c > w[j].f*w[i].c
	// })

	// out(w)
	out(rec(0, 0, 0))
	// out(gx, gy)
}
