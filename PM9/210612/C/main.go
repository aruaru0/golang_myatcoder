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

type pair struct {
	v, i int
}

type dist struct {
	d, i, j int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x := make([]pair, N)
	y := make([]pair, N)
	for i := 0; i < N; i++ {
		xx, yy := getI(), getI()
		x[i] = pair{xx, i}
		y[i] = pair{yy, i}
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i].v < x[j].v
	})
	sort.Slice(y, func(i, j int) bool {
		return y[i].v < y[j].v
	})

	a := make([]dist, 0)
	i, j := 0, N-1
	a = append(a, dist{x[j].v - x[i].v, x[i].i, x[j].i})
	a = append(a, dist{y[j].v - y[i].v, y[i].i, y[j].i})
	i++
	a = append(a, dist{x[j].v - x[i].v, x[i].i, x[j].i})
	a = append(a, dist{y[j].v - y[i].v, y[i].i, y[j].i})
	i--
	j--
	a = append(a, dist{x[j].v - x[i].v, x[i].i, x[j].i})
	a = append(a, dist{y[j].v - y[i].v, y[i].i, y[j].i})
	sort.Slice(a, func(i, j int) bool {
		return a[i].d > a[j].d
	})

	if a[0].i == a[1].i && a[0].j == a[1].j {
		out(a[2].d)
	} else {
		out(a[1].d)
	}
}
