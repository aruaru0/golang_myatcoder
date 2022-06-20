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

// 3角形の面積
func calc_triangle(x1, y1, x2, y2, x3, y3 int) int {
	return abs(x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = getI(), getI()
	}

	f := func(i, j, k int) int {
		i %= n
		j %= n
		k %= n
		// 3角形の8倍した値を返す
		return calc_triangle(x[i], y[i], x[j], y[j], x[k], y[k]) * 4
	}

	// 多角形の面積を求める
	tot := 0
	for j := 2; j < n; j++ {
		tot += f(0, j-1, j)
	}
	// 1/4の面積の８倍を計算
	tot /= 4

	s, ans := 0, math.MaxInt64
	j := 1
	for i := 0; i < n; i++ {
		for s < tot {
			s += f(i, j, j+1)
			ans = min(ans, abs(tot-s))
			j++
		}
		s -= f(i, i+1, j)
		ans = min(ans, abs(tot-s))
	}
	out(ans)
}
