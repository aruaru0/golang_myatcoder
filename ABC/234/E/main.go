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

func calc_f(s, t string) int {
	cnt := 0
	pos := 0
	for i := 0; i < len(s) && pos < len(t); i++ {
		if s[i] == t[pos] {
			pos++
			cnt++
		}
	}
	return cnt
}

func calc_t(s, t string) int {
	cnt := 0
	pos := len(t) - 1
	for i := len(s) - 1; i >= 0 && pos >= 0; i-- {
		if s[i] == t[pos] {
			pos--
			cnt++
		}
	}
	return cnt
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, t := getI(), getS()

	front := make([]int, N)
	back := make([]int, N)
	for i := 0; i < N; i++ {
		s := getS()
		f := calc_f(s, t)
		r := calc_t(s, t)
		front[i] = f
		back[i] = r
	}

	sort.Ints(front)
	sort.Ints(back)

	ans := 0
	for i := 0; i < N; i++ {
		cnt := len(t) - front[i]
		pos := lowerBound(back, cnt)
		// out(cnt, pos, N-pos)
		ans += N - pos
	}
	out(ans)
}
