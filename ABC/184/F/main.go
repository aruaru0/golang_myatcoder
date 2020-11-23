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

func solve(a []int) []int {
	n := 1
	for i := 0; i < len(a); i++ {
		n *= 2
	}
	ret := make([]int, 0)
	for i := 0; i < n; i++ {
		tot := 0
		for j := 0; j < len(a); j++ {
			if (i>>j)&1 == 1 {
				tot += a[j]
			}
		}
		ret = append(ret, tot)
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, T := getI(), getI()
	a := getInts(N)

	a0 := a[:N/2]
	a1 := a[N/2:]

	ret0 := solve(a0)
	ret1 := solve(a1)

	ans := 0
	sort.Ints(ret1)
	for i := 0; i < len(ret0); i++ {
		l := 0
		r := len(ret1)
		for l+1 != r {
			m := (l + r) / 2
			if ret0[i]+ret1[m] > T {
				r = m
			} else {
				l = m
			}
		}
		tot := ret0[i] + ret1[l]
		if tot <= T {
			ans = max(ans, tot)
		}
	}
	out(ans)
}
