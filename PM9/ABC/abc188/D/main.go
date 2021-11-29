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
	day := make(map[int]int)
	for i := 0; i < N; i++ {
		a[i], b[i], c[i] = getI(), getI(), getI()
		day[a[i]] = 1
		day[b[i]+1] = 1
	}

	x := []int{}
	for e := range day {
		x = append(x, e)
	}
	sort.Ints(x)
	for i := 0; i < len(x); i++ {
		day[x[i]] = i
	}

	d := make([]int, len(x))
	for i := 0; i < N; i++ {
		d[day[a[i]]] += c[i]
		d[day[b[i]+1]] -= c[i]
	}
	for i := 1; i < len(x); i++ {
		d[i] += d[i-1]
	}

	ans := 0
	// out(x)
	// out(d)
	for i := 0; i < len(x)-1; i++ {
		cnt := x[i+1] - x[i]
		// out(cnt, "cost", d[i])
		if d[i] > C {
			ans += C * cnt
		} else {
			ans += d[i] * cnt
		}
	}
	out(ans)
}
