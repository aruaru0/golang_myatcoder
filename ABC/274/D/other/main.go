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
	e, n int
}

var memo map[pos]*bool

func rec(cur, n, target int, x []int) bool {
	if n == len(x) {
		if cur == target {
			return true
		}
		return false
	}

	if memo[pos{cur, n}] != nil {
		return *memo[pos{cur, n}]
	}

	var ret bool
	ret = rec(cur+x[n], n+1, target, x)
	ret = ret || rec(cur-x[n], n+1, target, x)

	memo[pos{cur, n}] = &ret
	return ret
}

var N, x, y int
var a []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, x, y = getI(), getI(), getI()
	a = getInts(N)

	b0 := make([]int, 0)
	b1 := make([]int, 0)
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			b0 = append(b0, a[i])
		} else {
			b1 = append(b1, a[i])
		}
	}

	memo = make(map[pos]*bool)
	ret_x := rec(b0[0], 0, x, b0[1:])
	memo = make(map[pos]*bool)
	ret_y := rec(0, 0, y, b1)

	if ret_x && ret_y {
		out("Yes")
	} else {
		out("No")
	}
}
