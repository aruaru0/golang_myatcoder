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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	S = getS()
	T = getS()
	out(f(0, N) / 3)
}

var N int
var S, T string
var memo [101][101]*int

func f(l, r int) int {
	if memo[l][r] != nil {
		return *memo[l][r]
	}
	if r-l < 3 {
		return 0
	}
	if r-l == 3 {
		if S[l] == T[0] && S[l+1] == T[1] && S[l+2] == T[2] {
			return 3
		}
		return 0
	}
	ret := 0
	ret = nmax(ret, f(l+1, r), f(l, r-1))
	for m := l + 1; m < r; m++ {
		if S[l] == T[0] && S[m] == T[1] && S[r-1] == T[2] && f(l+1, m) == m-(l+1) && f(m+1, r-1) == (r-1)-(m+1) {
			ret = (r - l)
			break
		}
		ret = nmax(ret, f(l, m)+f(m, r))
	}
	memo[l][r] = &ret
	return ret
}
