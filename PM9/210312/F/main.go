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

var N, A, B, C, D int
var memo map[int]*int

func rec(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return D
	}
	if memo[n] != nil {
		return *memo[n]
	}
	ret := 0
	if n%2 == 0 {
		ret = rec(n/2) + A
	} else {
		ret = min(rec((n+1)/2)+A+D, rec((n-1)/2)+A+D)
	}
	if n%3 == 0 {
		ret = min(ret, rec(n/3)+B)
	} else {
		diff := n % 3
		ret = nmin(ret, rec((n-diff)/3)+B+D*diff, rec((n+3-diff)/3)+B+D*(3-diff))
	}
	if n%5 == 0 {
		ret = min(ret, rec(n/5)+C)
	} else {
		diff := n % 5
		ret = nmin(ret, rec((n-diff)/5)+C+D*diff, rec((n+5-diff)/5)+C+D*(5-diff))
	}
	if math.MaxInt64/n > D {
		ret = min(ret, D*abs(n))
	}
	memo[n] = &ret
	return ret
}

func solve() {
	N, A, B, C, D = getI(), getI(), getI(), getI(), getI()
	memo = make(map[int]*int)
	out(rec(N))
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		solve()
	}
}
