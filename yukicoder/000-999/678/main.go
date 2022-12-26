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
	x0, y0, x1, y1, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, L, R := getI(), getI(), getI()
	enemy := make([]int, 1280+1)
	e := make([]pos, N)
	for i := 0; i < N; i++ {
		x0, y0, x1, y1 := getI(), getI(), getI(), getI()
		e[i] = pos{x0, y0, x1, y1, i + 1}
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].y1 > e[j].y1
	})

	for i := 0; i < N; i++ {
		l, r := e[i].x0, e[i].x1
		l = max(0, l)
		r = min(r, 1280)
		for j := l; j <= r; j++ {
			if enemy[j] == 0 {
				enemy[j] = e[i].idx
			}
		}
	}

	hit := make([]int, N+1)
	for i := L; i <= R; i++ {
		hit[enemy[i]] = 1
	}
	for i := 1; i < N+1; i++ {
		out(hit[i])
	}
}
