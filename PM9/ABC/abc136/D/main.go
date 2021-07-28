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
	s := getS()
	s = "R" + s + "L"
	n := len(s)
	x := make([]int, n)
	for i := 1; i < n-1; i++ {
		x[i] = 1
	}
	y := make([]int, n)
	l, r := 0, 0
	for i := 0; i < n-1; i++ {
		if s[i] == 'L' && s[i+1] == 'R' {
			r = i
			// out(s[l:r+1], x[l:r+1])
			L, R := 0, 0
			pos := 0
			flg := false
			for j := l; j <= r; j++ {
				if flg == false && s[j] == 'L' {
					flg = true
					pos = j
				}
				if s[j] == 'L' {
					L += x[j]
				} else {
					R += x[j]
				}
			}
			// out(L, R, pos, y)
			if (L+R)%2 == 0 {
				L, R = (L+R)/2, (L+R)/2
			} else {
				if L%2 == 1 {
					L, R = (L+R)/2+1, (L+R)/2
				} else {
					L, R = (L+R)/2, (L+R)/2+1
				}
			}

			y[pos-1] = R
			y[pos] = L
			l = r + 1
		}
	}

	L, R := 0, 0
	pos := 0
	flg := false
	for j := l; j < n; j++ {
		if flg == false && s[j] == 'L' {
			flg = true
			pos = j
		}
		if s[j] == 'L' {
			L += x[j]
		} else {
			R += x[j]
		}
	}
	if (L+R)%2 == 0 {
		L, R = (L+R)/2, (L+R)/2
	} else {
		if L%2 == 1 {
			L, R = (L+R)/2+1, (L+R)/2
		} else {
			L, R = (L+R)/2, (L+R)/2+1
		}
	}
	y[pos-1] = R
	y[pos] = L
	// out(s[l:], x[l:], y)
	for i := 1; i < n-1; i++ {
		fmt.Fprint(wr, y[i], " ")
	}
	out()
}
