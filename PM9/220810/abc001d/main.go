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

func f(s string, isStart bool) (int, int) {
	hour, _ := strconv.Atoi(s[:2])
	min, _ := strconv.Atoi(s[2:])
	if isStart {
		min = 5 * (min / 5)
	} else {
		if min%5 != 0 {
			min += 5 - min%5
		}
		if min == 60 {
			hour++
			min = 0
		}
	}
	return hour, min
}

func g(n int) {
	n *= 5
	hour := n / 60
	min := n % 60
	fmt.Fprintf(wr, "%2.2d%2.2d", hour, min)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := make([]bool, 24*12)
	for i := 0; i < N; i++ {
		t := getS()
		// out(t)
		h, m := f(t[:4], true)
		s := h*12 + m/5
		// out("start", h, m)
		h, m = f(t[5:], false)
		e := h*12 + m/5
		// out("end", h, m)
		// out(s, e)
		for j := s; j < e; j++ {
			a[j] = true
		}
	}

	ok := false
	s, e := 0, 0
	for i, v := range a {
		if v == true {
			if ok == false {
				s = i
				ok = true
			}
		} else {
			if ok == true {
				e = i
				ok = false
				g(s)
				fmt.Fprintf(wr, "-")
				g(e)
				out()
			}
		}
	}
	if ok == true {
		g(s)
		fmt.Fprintf(wr, "-")
		out("2400")
	}
}
