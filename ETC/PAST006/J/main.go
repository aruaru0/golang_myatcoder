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

const inf = 1e18
const delta = 1e-10

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, C := getI(), getF()
	x := make([]float64, N)
	y := make([]float64, N)
	l := inf
	r := -inf
	for i := 0; i < N; i++ {
		x[i], y[i] = getF(), getF()
		l = math.Min(l, x[i])
		r = math.Max(r, x[i])
	}
	for i := 0; i < 1000; i++ {
		diff := (r - l) / 100
		mid := (l + r) / 2
		mid2 := mid + diff
		// out(l, r, mid, mid2)
		var sum float64
		for j := 0; j < N; j++ {
			sum += (x[j]-mid)*(x[j]-mid) + (y[j]-C)*(y[j]-C)
		}
		var sum2 float64
		for j := 0; j < N; j++ {
			sum2 += (x[j]-mid2)*(x[j]-mid2) + (y[j]-C)*(y[j]-C)
		}

		// out(mid, mid2)
		if sum > sum2 {
			l = mid
		} else {
			r = mid2
		}
	}

	// out(l, r)
	mid := l
	ans := 0.0
	for j := 0; j < N; j++ {
		ans += (x[j]-mid)*(x[j]-mid) + (y[j]-C)*(y[j]-C)
	}
	out(ans)
}
