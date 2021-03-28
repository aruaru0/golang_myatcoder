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

func rec(b []int) {
	n := len(b)
	out(b)
	cur := b[0]
	xor := 0
	or := 0
	for i := 0; i < n; i++ {
		if cur == b[i] {
			or |= a[i]
		} else {
			xor |= or
			or = a[i]
			cur = b[i]
		}
		// out("i", i, cur, b[i], a[i], "or", or)
	}
	xor ^= or
	ans = min(ans, xor)
	// out(b, xor)

	pos := 0
	for i := 0; i < n; i++ {
		if b[i] != i {
			break
		}
		pos = i
	}
	if pos == n-1 {
		return
	}
	for i := pos + 1; i < n; i++ {
		if b[i] != b[i-1] {
			break
		}
		pos = i
	}
	b[pos] = b[pos-1] + 1
	rec(b)
}

func f(k int) int {
	s := "00000000000000000000" + strconv.FormatInt(int64(k), 2)
	s = s[len(s)-N+1:]

	b := make([]int, N)
	cnt := 0
	for i, e := range s {
		if e == '1' {
			cnt++
		}
		b[i+1] = cnt
	}
	// out(s, b)
	cur := b[0]
	xor := 0
	or := 0
	for i := 0; i < N; i++ {
		if cur == b[i] {
			or |= a[i]
		} else {
			xor ^= or
			or = a[i]
			cur = b[i]
		}
		// out("i", i, cur, b[i], a[i], "or", or)
	}
	xor ^= or
	return xor
}

var a []int
var ans int
var N int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	a = getInts(N)

	n := 1 << (N - 1)
	ans := int(1e18)
	for k := 0; k < n; k++ {
		ret := f(k)
		ans = min(ans, ret)
	}
	out(ans)
}
