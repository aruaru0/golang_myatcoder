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

func f(a []int, m int) ([]int, bool) {
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		if a[i] != m {
			a[i]++
			if i != n-1 {
				for j := i + 1; j < n; j++ {
					a[j] = a[i]
				}
			}
			return a, false
		}
	}
	return a, true
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, Q := getI(), getI(), getI()
	A := make([]int, Q)
	B := make([]int, Q)
	C := make([]int, Q)
	D := make([]int, Q)
	for i := 0; i < Q; i++ {
		A[i], B[i], C[i], D[i] = getI()-1, getI()-1, getI(), getI()
	}

	x := []int{}
	for i := 0; i < N; i++ {
		x = append(x, 1)
	}
	ans := 0
	for {
		tot := 0
		for i := 0; i < Q; i++ {
			a, b, c, d := A[i], B[i], C[i], D[i]
			if x[b]-x[a] == c {
				tot += d
			}
		}
		ans = max(ans, tot)
		t, ok := f(x, M)
		x = t
		if ok {
			break
		}
	}
	out(ans)
}
