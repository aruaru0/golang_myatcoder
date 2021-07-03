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

type pair struct {
	x, c int
}

const inf = int(1e18)

var cmin, cmax []int
var N int

var memo map[pair]*int

func rec(n, pos int) int {
	if n == N {
		return abs(pos)
	}
	for cmin[n] == inf {
		n++
		if n == N {
			return abs(pos)
		}
	}
	if memo[pair{n, pos}] != nil {
		return *memo[pair{n, pos}]
	}

	ret := 0
	d := cmax[n] - cmin[n]
	ret = rec(n+1, cmin[n]) + d + abs(pos-cmax[n])
	ret = min(ret, rec(n+1, cmax[n])+d+abs(pos-cmin[n]))

	memo[pair{n, pos}] = &ret
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	p := make([]pair, N)
	cmin = make([]int, N)
	cmax = make([]int, N)
	for i := 0; i < N; i++ {
		cmin[i] = inf
		cmax[i] = -inf
	}
	for i := 0; i < N; i++ {
		x, c := getI(), getI()-1
		p[i] = pair{x, c}
		chmin(&cmin[c], x)
		chmax(&cmax[c], x)
	}

	memo = make(map[pair]*int)
	out(rec(0, 0))
}
