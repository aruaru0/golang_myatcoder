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

var L []int
var R []int
var N int

type pos struct {
	l0, r0, l1, r1, flg int
}

var memo map[pos]int

func rec(n, l0, r0, l1, r1, flg int, pat []int) int {
	if n == N {
		if flg != 3 {
			return 0
		}
		d0, d1 := 0, 0
		if l0 <= r0 {
			d0 = r0 - l0 + 1
		}
		if l1 <= r1 {
			d1 = r1 - l1 + 1
		}
		// out(pat, l0, r0, l1, r1, d0+d1)
		return d0 + d1
	}

	v, ok := memo[pos{l0, r0, l1, r1, flg}]
	if ok {
		return v
	}

	l := L[n]
	r := R[n]
	l = max(l, l0)
	r = min(r, r0)
	// out(l0, r0, "-", L[n], R[n], "->", l, r)
	ret := rec(n+1, l, r, l1, r1, flg|1, append(pat, 1))
	l = L[n]
	r = R[n]
	l = max(l, l1)
	r = min(r, r1)
	ret = max(ret, rec(n+1, l0, r0, l, r, flg|2, append(pat, 2)))

	memo[pos{l0, r0, l1, r1, flg}] = ret
	return ret
}

type pair struct {
	a, b int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		l, r := getI(), getI()
		a[i] = [2]int{l, r + 1}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	miR := make([]int, n+1)
	miR[n] = 1 << 60
	maL := a[n-1][0]
	for i := n - 1; i >= 0; i-- {
		miR[i] = min(miR[i+1], a[i][1])
	}

	ans := 0
	r := 1 << 60
	for i := 0; i < n-1; i++ {
		l1, r1 := a[i][0], a[i][1]
		l2, r2 := maL, miR[i+1]
		s1 := r1 - l1 + max(0, min(r2, r)-l2)
		s2 := max(0, min(r1, r)-l1) + max(0, r2-l2)
		ans = max(ans, max(s1, s2))
		r = min(r, r1)
	}
	out(ans)
}
