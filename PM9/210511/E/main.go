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

var N int
var t string

func f(a0, a1 int) (bool, []int) {
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = -1
	}
	s[0] = a0
	s[1] = a1
	last := 0
	if t[0] == 'o' {
		if a0 == 0 {
			last = a1
		} else {
			last = 1 - a1
		}
	} else {
		if a0 == 0 {
			last = 1 - a1
		} else {
			last = a1
		}
	}
	for i := 1; i < N-1; i++ {
		if s[i] == 0 {
			if t[i] == 'o' {
				s[i+1] = s[i-1]
			} else {
				s[i+1] = 1 - s[i-1]
			}
		} else {
			if t[i] == 'o' {
				s[i+1] = 1 - s[i-1]
			} else {
				s[i+1] = s[i-1]
			}
		}
	}

	first := 0
	if s[N-1] == 0 {
		if t[N-1] == 'o' {
			first = s[N-2]
		} else {
			first = 1 - s[N-2]
		}
	} else {
		if t[N-1] == 'o' {
			first = 1 - s[N-2]
		} else {
			first = s[N-2]
		}
	}

	// out(t)
	// out(s, last)

	if s[N-1] == last && s[0] == first {
		return true, s
	}
	return false, s
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	t = getS()

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			ok, s := f(i, j)
			if ok {
				ans := ""
				for _, e := range s {
					if e == 0 {
						ans += "S"
					} else {
						ans += "W"
					}
				}
				out(ans)
				return
			}
		}
	}
	out(-1)
}
