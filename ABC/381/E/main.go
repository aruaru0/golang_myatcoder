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

func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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
	N, Q := getI(), getI()
	s := getS()

	one, two, slash := make([]int, 0), make([]int, 0), make([]int, 0)
	for i := 0; i < N; i++ {
		if s[i] == '1' {
			one = append(one, i)
		} else if s[i] == '2' {
			two = append(two, i)
		} else {
			slash = append(slash, i)
		}
	}

	// out(one, two, slash)

	for qi := 0; qi < Q; qi++ {
		l, r := getI()-1, getI()

		f := func(tot int) bool {
			if tot == 0 {
				j := lowerBound(slash, l)
				if j >= len(slash) {
					return false
				}
				return slash[j] < r
			}
			i := lowerBound(one, l)
			if i+tot-1 >= len(one) {
				return false
			}
			j := lowerBound(slash, one[i+tot-1])
			if j >= len(slash) {
				return false
			}
			k := lowerBound(two, slash[j])
			if k+tot-1 >= len(two) {
				return false
			}
			return two[k+tot-1] < r
		}

		ok, ng := -1, N+1
		for ok+1 < ng {
			m := (ok + ng) / 2
			if f(m) {
				ok = m
			} else {
				ng = m
			}
		}
		if ok == -1 {
			out(0)
		} else {
			out(ok*2 + 1)
		}
	}

}
