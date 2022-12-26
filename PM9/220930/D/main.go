package main

import (
	"bufio"
	"fmt"
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

var P, B []int
var N, X int

func rec(n, x, p int) {
	if n == 0 {
		fmt.Println(p + 1)
		os.Exit(0)
	}

	// B
	x++
	if x == X {
		fmt.Println(p)
		os.Exit(0)
	}

	// n-1バーガー
	if x+B[n-1] >= X {
		rec(n-1, x, p)
	}
	x += B[n-1]
	p += P[n-1]

	// P
	p++
	x++
	if x == X {
		fmt.Println(p)
		os.Exit(0)
	}

	// n-1バーガー
	if x+B[n-1] >= X {
		rec(n-1, x, p)
	}
	x += B[n-1]
	p += P[n-1]

	// B
	x++
	fmt.Println(p)
	os.Exit(0)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	B = make([]int, 51)
	P = make([]int, 51)
	B[0] = 1 // レベルiバーガーの層数
	P[0] = 1 // レベルiバーガーのパティの数
	for i := 0; i < 50; i++ {
		B[i+1] = 2*B[i] + 3
		P[i+1] = 2*P[i] + 1
	}

	// out(B)
	// out(P)
	// wr.Flush()
	N, X = getI(), getI()
	rec(N, 0, 0)
}
