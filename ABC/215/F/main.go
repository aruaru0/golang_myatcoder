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

func ok(k int) bool {
	l, r := 0, 0
	// out("k=", k, p)
	for r < N && l < N {
		if r < N && p[r].x-p[l].x < k {
			r++
		}
		// out(l, r)
		if r == N {
			break
		}
		// out(l, r, "l", miny[l], maxy[l], "r", rminy[r], rmaxy[r])
		if p[r].x-p[l].x >= k && (abs(miny[l]-rmaxy[r]) >= k || abs(maxy[l]-rminy[r]) >= k) {
			// out("pass")
			return true
		}
		// out(l, r, p[r], miny, maxy)
		if l < N && p[r].x-p[l].x >= k {
			l++
		}
	}
	return false
}

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
	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})

	miny = make([]int, N)
	maxy = make([]int, N)
	miny[0] = p[0].y
	maxy[0] = p[0].y
	for i := 1; i < N; i++ {
		miny[i] = min(miny[i-1], p[i].y)
		maxy[i] = max(maxy[i-1], p[i].y)
	}

	rminy = make([]int, N)
	rmaxy = make([]int, N)
	rminy[N-1] = p[N-1].y
	rmaxy[N-1] = p[N-1].y
	for i := N - 2; i >= 0; i-- {
		rminy[i] = min(rminy[i+1], p[i].y)
		rmaxy[i] = max(rmaxy[i+1], p[i].y)
	}

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
