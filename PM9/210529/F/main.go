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

var N, A, B int
var v []int

type pair struct {
	v, c int
}

type item struct {
	i, v, c int
}

var memo [][]*pair

func rec(n, m, a int) (int, int) {
	if m == a {
		return 0, 1
	}
	if n == N {
		return 0, 0
	}
	if memo[n][m] != nil {
		return memo[n][m].v, memo[n][m].c
	}
	ret1, cnt1 := rec(n+1, m, a)
	ret2, cnt2 := rec(n+1, m+1, a)
	ret2 += v[n]

	ret := 0
	cnt := 0
	if ret1 == ret2 {
		cnt = cnt1 + cnt2
		ret = ret1
	} else if ret1 > ret2 {
		cnt = cnt1
		ret = ret1
	} else {
		cnt = cnt2
		ret = ret2
	}

	memo[n][m] = &pair{ret, cnt}

	return ret, cnt
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, A, B = getI(), getI(), getI()
	v = getInts(N)

	t := make([]item, 0)
	for i := A; i <= B; i++ {
		memo = make([][]*pair, N)
		for j := 0; j < N; j++ {
			memo[j] = make([]*pair, N)
		}
		ret, cnt := rec(0, 0, i)
		t = append(t, item{i, ret, cnt})
		// out(i, ret, float64(ret)/float64(i), cnt)
	}

	sort.Slice(t, func(i, j int) bool {
		return t[i].v*t[j].i > t[j].v*t[i].i
	})

	value, num := t[0].v, t[0].i
	sum, items, cnt := 0, 0, 0
	for _, e := range t {
		if e.v*num == value*e.i {
			sum += e.v
			items += e.i
			cnt += e.c
		} else {
			break
		}
	}

	// out(sum, items, cnt)
	out(float64(sum) / float64(items))
	out(cnt)
}
