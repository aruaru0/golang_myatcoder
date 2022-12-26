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

func f(a, b, c, bit int) bool {
	n := 1 << bit
	for j := 0; j < n; j++ {
		x := make([]int, 0, bit)
		for k := 0; k < bit; k++ {
			x = append(x, (j>>k)&1)
		}
		and, or, xor := x[0], x[0], x[0]
		for k := 1; k < len(x); k++ {
			and &= x[k]
			or |= x[k]
			xor ^= x[k]
		}
		// out(x, and, or, xor)
		if and == b && or == a && xor == c {
			return true
		}
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		A, B, C := getI(), getI(), getI()
		ok := false
		for bit := 1; bit < 5; bit++ {
			flg := true
			for j := 0; j < 60; j++ {
				a := (A >> j) & 1
				b := (B >> j) & 1
				c := (C >> j) & 1
				// out(f(a, b, c, bit), a, b, c)
				flg = flg && f(a, b, c, bit)
			}
			ok = ok || flg
		}
		if ok {
			out("Yes")
		} else {
			out("No")
		}
	}
}
