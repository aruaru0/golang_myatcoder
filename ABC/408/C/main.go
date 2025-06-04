package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type nums interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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
func max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// min for n entry
func nmin[T Ordered](a ...T) T {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax[T Ordered](a ...T) T {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin[T Ordered](a *T, b T) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax[T Ordered](a *T, b T) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
}

func asub[T nums](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}

func abs[T nums](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound[T nums](a []T, x T) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound[T nums](a []T, x T) int {
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
	N, M := getI(), getI()

	a := make([]int, N+1)
	for i := 0; i < M; i++ {
		l, r := getI()-1, getI()
		a[l]++
		a[r]--
	}

	for i := 1; i < N+1; i++ {
		a[i] += a[i-1]
	}

	minM := M
	for i := 0; i < N; i++ {
		minM = min(minM, a[i])
	}
	// out(a)
	out(minM)
}
