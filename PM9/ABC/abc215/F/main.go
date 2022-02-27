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
	x, y int
}

var N int
var p []pos
var miny, maxy []int
var rminy, rmaxy []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	p = make([]pos, N)
	for i := 0; i < N; i++ {
		x, y := getI(), getI()
		p[i] = pos{x, y}
	}
	// xでソートして格納
	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})

	ok := func(d int) bool {
		j := 0
		// 条件を満たす範囲のyの最大値と最小値を求める
		ly, ry := int(1e10), int(-1e10)
		for i := 0; i < N; i++ {
			x, y := p[i].x, p[i].y
			// |xi - xj|
			for j < i && x-p[j].x >= d { // |xi-xj| >= dの間
				ny := p[j].y
				ly = min(ly, ny)
				ry = max(ry, ny)
				j++
			}
			// |xi - xj| >= d and |yi - yj| >= dが条件

			// ly, ryは、|xi - xj| >= dとなる場合更新されるので、下記条件をみたすのはｘが条件を満たしたとき
			// ここまで見た中で最小のyとの絶対値差がdより大きい
			if y-ly >= d {
				return true
			}
			// ここまで見た中で最小のyとの絶対値差がdより大きい
			if ry-y >= d {
				return true
			}
		}
		return false
	}

	// 二分探索
	l, r := 0, int(1e10)
	for l+1 != r {
		m := (l + r) / 2
		if !ok(m) {
			r = m
		} else {
			l = m
		}
	}
	out(l)
}
