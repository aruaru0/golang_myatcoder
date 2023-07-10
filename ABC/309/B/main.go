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
	N := getI()
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		t := getS()
		a[i] = make([]int, N)
		for j, e := range t {
			if e == '0' {
				a[i][j] = 0
			} else {
				a[i][j] = 1
			}
		}
	}

	x := make([]int, 0)
	for i := 0; i < N; i++ {
		x = append(x, a[0][i])
	}
	for i := 1; i < N-1; i++ {
		x = append(x, a[i][N-1])
	}
	for i := N - 1; i >= 0; i-- {
		x = append(x, a[N-1][i])
	}
	for i := N - 2; i >= 1; i-- {
		x = append(x, a[i][0])
	}

	idx := len(x) - 1
	for i := 0; i < N; i++ {
		a[0][i] = x[idx]
		idx++
		idx %= len(x)
	}
	for i := 1; i < N-1; i++ {
		a[i][N-1] = x[idx]
		idx++
	}
	for i := N - 1; i >= 0; i-- {
		a[N-1][i] = x[idx]
		idx++
	}
	for i := N - 2; i >= 1; i-- {
		a[i][0] = x[idx]
		idx++
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fprint(wr, a[i][j])
		}
		out()
	}
}
