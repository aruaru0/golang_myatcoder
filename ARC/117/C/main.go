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

var bf, bg [1 << 19]int

func inittbl() {
	for i := 1; i <= 400000; i++ {
		pos := i
		for pos%3 == 0 {
			pos /= 3
			bf[i]++
		}
		bg[i] = pos % 3
	}
	bg[0] = 1
	for i := 1; i <= 400000; i++ {
		bf[i] += bf[i-1]
	}
	for i := 1; i <= 400000; i++ {
		bg[i] = (bg[i] * bg[i-1]) % 3
	}
}

func ncr_mod_3(n, r int) int {
	if bf[n] != bf[r]+bf[n-r] {
		return 0
	}

	if bg[n] == 1 && bg[r]*bg[n-r] == 1 {
		return 1
	}
	if bg[n] == 1 && bg[r]*bg[n-r] == 2 {
		return 2
	}
	if bg[n] == 1 && bg[r]*bg[n-r] == 4 {
		return 1
	}
	if bg[n] == 2 && bg[r]*bg[n-r] == 1 {
		return 2
	}
	if bg[n] == 2 && bg[r]*bg[n-r] == 2 {
		return 1
	}
	if bg[n] == 2 && bg[r]*bg[n-r] == 4 {
		return 2
	}
	return -1
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	inittbl()
	N := getI()
	S := getS()
	ret := 0
	for i := 0; i < N; i++ {
		p1, p2 := 0, ncr_mod_3(N-1, i)
		if S[i] == 'B' {
			p1 = 0
		}
		if S[i] == 'W' {
			p1 = 1
		}
		if S[i] == 'R' {
			p1 = 2
		}
		ret += p1 * p2
		ret %= 3
	}
	if N%2 == 0 {
		ret = (3 - ret) % 3
	}
	if ret == 0 {
		out("B")
	}
	if ret == 1 {
		out("W")
	}
	if ret == 2 {
		out("R")
	}
}
