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

func f(s, e int, b []int) (int, int) {
	// out(b[s : e+1])
	l := s + 1
	r := e
	for l+1 != r {
		mid := (l + r) / 2
		L := b[mid] - b[s]
		R := b[e] - b[mid]
		if L > R {
			r = mid
		} else {
			l = mid
		}
	}

	L := b[l] - b[s]
	R := b[e] - b[l]
	diff := abs(L - R)
	for i := l - 1; i <= l+1; i++ {
		if i <= s {
			continue
		}
		if i == e {
			break
		}
		LL := b[i] - b[s]
		RR := b[e] - b[i]
		diff2 := abs(LL - RR)
		// out(i, diff2, l)
		if diff > diff2 {
			L = LL
			R = RR
		}
	}

	return L, R
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	b := make([]int, N+1)
	for i := 1; i <= N; i++ {
		b[i] = b[i-1] + a[i-1]
	}
	// out("b", b)
	ans := int(1e15)
	for i := 2; i <= N-2; i++ {
		p, q := f(0, i, b)
		r, s := f(i, N, b)
		// out(p, q, r, s)
		tot := nmax(p, q, r, s) - nmin(p, q, r, s)
		ans = min(ans, tot)
	}
	out(ans)
}
