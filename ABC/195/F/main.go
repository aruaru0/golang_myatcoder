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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

var a []int

type pair struct {
	n   int
	sel string
}

var memo map[pair]*int
var tbl [][]bool

func rec(n int, sel string) int {
	if n == len(a) {
		return 1
	}
	if memo[pair{n, sel}] != nil {
		return *memo[pair{n, sel}]
	}
	ok := true
	for i, e := range sel {
		if e == 'o' {
			if tbl[i][n] == false {
				ok = false
				break
			}
		}
	}
	ret := 0
	if ok {
		ret += rec(n+1, sel+"o")
	}
	ret += rec(n+1, sel+"-")
	memo[pair{n, sel}] = &ret
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	A, B := getI(), getI()

	a = make([]int, 0)
	for i := A; i < B+1; i++ {
		a = append(a, i)
	}

	n := len(a)
	tbl = make([][]bool, n)
	for i := 0; i < n; i++ {
		tbl[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if gcd(a[i], a[j]) == 1 {
				tbl[i][j] = true
			}
		}
	}

	memo = make(map[pair]*int)
	out(rec(0, ""))
}
