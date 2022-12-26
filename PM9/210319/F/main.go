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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

type pair struct {
	x, y int
}

type item struct {
	w, v int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	ga, sa, ba := getI(), getI(), getI()
	gb, sb, bb := getI(), getI(), getI()

	x := []item{{ga, gb - ga}, {sa, sb - sa}, {ba, bb - ba}}
	dp0 := make([]int, N+1)
	// out(x)
	d := 0
	for i := 0; i < N+1; i++ {
		for j := 0; j < 3; j++ {
			if i+x[j].w < N+1 {
				dp0[i+x[j].w] = max(dp0[i+x[j].w], dp0[i]+x[j].v)
				d = max(d, dp0[i+x[j].w])
			}
		}
	}
	N += d
	y := []item{{gb, ga - gb}, {sb, sa - sb}, {bb, ba - bb}}
	dp1 := make([]int, N+1)
	d = 0
	for i := 0; i < N+1; i++ {
		for j := 0; j < 3; j++ {
			if i+y[j].w < N+1 {
				dp1[i+y[j].w] = max(dp1[i+y[j].w], dp1[i]+y[j].v)
				d = max(d, dp1[i+y[j].w])
			}
		}
	}
	N += d
	out(N)
}
