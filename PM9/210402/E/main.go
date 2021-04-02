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

var N, K int
var x, y []int

func f3(i, j, k, l int) (int, int) {
	x0 := nmin(x[i], x[j], x[k], x[l])
	y0 := nmin(y[i], y[j], y[k], y[l])
	x1 := nmax(x[i], x[j], x[k], x[l])
	y1 := nmax(y[i], y[j], y[k], y[l])

	cnt := 0
	for i := 0; i < N; i++ {
		if x0 <= x[i] && x[i] <= x1 && y0 <= y[i] && y[i] <= y1 {
			cnt++
		}
	}
	return cnt, (x1 - x0) * (y1 - y0)
}

func f2(i, j, k int) (int, int) {
	x0 := nmin(x[i], x[j], x[k])
	y0 := nmin(y[i], y[j], y[k])
	x1 := nmax(x[i], x[j], x[k])
	y1 := nmax(y[i], y[j], y[k])

	cnt := 0
	for i := 0; i < N; i++ {
		if x0 <= x[i] && x[i] <= x1 && y0 <= y[i] && y[i] <= y1 {
			cnt++
		}
	}
	return cnt, (x1 - x0) * (y1 - y0)
}

func f(n, m int) (int, int) {
	x0, y0 := x[n], y[n]
	x1, y1 := x[m], y[m]
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}

	cnt := 0
	for i := 0; i < N; i++ {
		if x0 <= x[i] && x[i] <= x1 && y0 <= y[i] && y[i] <= y1 {
			cnt++
		}
	}
	return cnt, (x1 - x0) * (y1 - y0)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K = getI(), getI()
	x = make([]int, N)
	y = make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getI(), getI()
	}

	ans := math.MaxInt64
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			n, s := f(i, j)
			if n >= K {
				ans = min(ans, s)
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				n, s := f2(i, j, k)
				if n >= K {
					ans = min(ans, s)
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				for l := k + 1; l < N; l++ {
					n, s := f3(i, j, k, l)
					if n >= K {
						ans = min(ans, s)
					}
				}
			}
		}
	}

	out(ans)
}
