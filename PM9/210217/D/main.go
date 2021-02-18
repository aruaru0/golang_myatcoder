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
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	b, c := getI(), getI()

	s := make([]pair, 0)
	// --
	s = append(s, pair{b, b - c/2})
	// -o
	s = append(s, pair{-b, -(b - (c-1)/2)})
	// ox
	s = append(s, pair{-b, -b - (c-1)/2})
	// oo
	if c >= 2 {
		s = append(s, pair{b, -(-b - (c-2)/2)})
	}

	xx := make([]pair, 0)
	for _, e := range s {
		x, y := e.x, e.y
		if x > y {
			x, y = y, x
		}
		xx = append(xx, pair{x, +1})
		xx = append(xx, pair{y + 1, -1})
	}

	sort.Slice(xx, func(i, j int) bool {
		return xx[i].x < xx[j].x
	})

	pre := -int(1e18 + 1)
	cnt := 0
	ans := 0
	for _, e := range xx {
		if cnt > 0 {
			ans += e.x - pre
		}
		cnt += e.y
		pre = e.x
	}
	out(ans)
}
