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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, C := getI(), getI()
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	m := make(map[int]bool)
	for i := 0; i < N; i++ {
		x, y, z := getI(), getI(), getI()
		a[i] = x
		b[i] = y + 1
		c[i] = z
		m[x] = true
		m[y+1] = true
	}
	x := make([]int, 0)
	for e := range m {
		x = append(x, e)
	}
	sort.Ints(x)
	toIdx := make(map[int]int)
	toDay := make(map[int]int)
	for i, e := range x {
		toIdx[e] = i
		toDay[i] = e
	}
	// out(a, b, toIdx, toDay)

	d := make([]int, len(m))
	for i := 0; i < N; i++ {
		from, to, cost := a[i], b[i], c[i]
		d[toIdx[from]] += cost
		d[toIdx[to]] -= cost
	}

	for i := 1; i < len(d); i++ {
		d[i] += d[i-1]
	}

	from := 0
	cost := 0
	ans := 0
	for i := 0; i < len(d); i++ {
		to := toDay[i]
		if cost >= C {
			ans += (to - from) * C
		} else {
			ans += (to - from) * cost
		}
		cost = d[i]
		from = to
	}

	out(ans)
}
