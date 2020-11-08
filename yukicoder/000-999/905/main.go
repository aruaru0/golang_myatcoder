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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)

	aL := make([]int, N)
	aG := make([]int, N)

	cnt := 0
	for i := 1; i < N; i++ {
		if a[i-1] <= a[i] {
			cnt++
		} else {
			cnt = 0
		}
		aL[i] = cnt
	}
	cnt = 0
	for i := 1; i < N; i++ {
		if a[i-1] >= a[i] {
			cnt++
		} else {
			cnt = 0
		}
		aG[i] = cnt
	}

	// out(aL)
	// out(aG)

	Q := getI()
	for i := 0; i < Q; i++ {
		l, r := getI(), getI()
		dL := aL[r] - aL[l]
		dG := aG[r] - aG[l]
		flgL := 0
		// out("dL, dG", dL, dG, r-l)
		if dL == r-l {
			flgL = 1
		}
		flgG := 0
		if dG == r-l {
			flgG = 1
		}
		out(flgL, flgG)
	}
}
