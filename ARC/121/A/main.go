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

type pt struct {
	p, i int
}

type pt2 struct {
	p, i, j int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	p := make([]pos, N)

	x := make([]pt, N)
	y := make([]pt, N)
	for i := 0; i < N; i++ {
		xx, yy := getI(), getI()
		p[i] = pos{xx, yy}
		x[i] = pt{xx, i}
		y[i] = pt{yy, i}
	}

	// if N < 1000 {
	// 	a := make([]int, 0)
	// 	for i := 0; i < N; i++ {
	// 		for j := i + 1; j < N; j++ {
	// 			d := max(abs(x[i].p-x[j].p), abs(y[i].p-y[j].p))
	// 			a = append(a, d)
	// 		}
	// 	}
	// 	sort.Slice(a, func(i, j int) bool {
	// 		return a[i] > a[j]
	// 	})
	// 	out(a[1])
	// 	return
	// }

	sort.Slice(x, func(i, j int) bool {
		return x[i].p < x[j].p
	})
	sort.Slice(y, func(i, j int) bool {
		return y[i].p < y[j].p
	})

	t := make([]pt2, 0)
	for i := 1; i < N-1; i++ {
		t = append(t, pt2{abs(x[0].p - x[i].p), x[0].i, x[i].i})
		t = append(t, pt2{abs(y[0].p - y[i].p), y[0].i, y[i].i})
	}
	for i := N - 2; i >= 0; i-- {
		t = append(t, pt2{abs(x[N-1].p - x[i].p), x[i].i, x[N-1].i})
		t = append(t, pt2{abs(y[N-1].p - y[i].p), y[i].i, y[N-1].i})
	}

	sort.Slice(t, func(i, j int) bool {
		return t[i].p > t[j].p
	})

	m := make(map[pt]bool)
	cnt := 0
	for k := 0; k < len(t); k++ {
		i, j := t[k].i, t[k].j
		if i > j {
			i, j = j, i
		}
		if m[pt{i, j}] {
			continue
		}
		m[pt{i, j}] = true
		if cnt == 1 {
			out(t[k].p)
			return
		}
		cnt++
	}
}
