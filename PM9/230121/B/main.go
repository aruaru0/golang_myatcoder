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

func dist(i, j int) float64 {
	dx, dy := x[i]-x[j], y[i]-y[j]
	return math.Hypot(dx, dy)
}

var x, y []float64

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x = make([]float64, N)
	y = make([]float64, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getF(), getF()
	}

	prev := 0
	used := make([]bool, N)
	ans := []int{0}
	used[0] = true
	D := 0.0
	for i := 0; i < N-1; i++ {
		min_dis, nxt := math.MaxFloat64, -1
		for j := 0; j < N; j++ {
			if used[j] {
				continue
			}
			dis := dist(prev, j)
			if min_dis > dis {
				min_dis = dis
				nxt = j
			}
		}
		// out("-->", nxt, ans)
		used[nxt] = true
		ans = append(ans, nxt)
		prev = nxt
		D += min_dis
	}
	D += dist(prev, 0)
	ans = append(ans, 0)

	for _, e := range ans {
		out(e + 1)
	}
}
