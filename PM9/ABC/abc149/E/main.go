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
	N, M := getI(), getI()
	a := getInts(N)
	sort.Ints(a)

	b := make([]int, N+1)
	for i := 1; i <= N; i++ {
		b[i] = b[i-1] + a[i-1]
	}

	l, r := 0, int(1e16)
	for l+1 != r {
		m := (l + r) / 2
		tot := 0
		// out("m", m)
		for i := 0; i < N; i++ {
			pos := lowerBound(a, m-a[i])
			// out(pos, m-a[i], a, N-pos)
			tot += N - pos
		}
		// out("tot", tot, "m", m, l, r)
		if tot < M {
			r = m
		} else {
			l = m
		}
	}

	// out(l)
	// out(b, l)
	// out(a)
	ans := 0
	m := 0
	for i := 0; i < N; i++ {
		pos := lowerBound(a, l-a[i])
		d := b[len(b)-1] - b[pos]
		// out(i, d, pos, l-a[i])
		ans += d + a[i]*(len(b)-1-pos)
		m += len(b) - 1 - pos
	}
	// out(m)
	out(ans - l*(m-M))
}
