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
	N, M, K, Q := getI(), getI(), getI(), getI()
	p0 := make([]int, 0)
	p1 := make([]int, 0)
	for i := 0; i < N; i++ {
		p, t := getI(), getI()
		if t == 0 {
			p0 = append(p0, p)
		} else {
			p1 = append(p1, p)
		}
	}
	sort.Ints(p0)
	sort.Ints(p1)
	b0 := make([]int, len(p0)+1)
	for i := 0; i < len(b0)-1; i++ {
		b0[i+1] = b0[i] + p0[i]
	}
	b1 := make([]int, len(p1)+1)
	for i := 0; i < len(b1)-1; i++ {
		b1[i+1] = b1[i] + p1[i]
	}
	// out(b0, b1)

	ans := int(1e18)
	for i := 0; i < len(b0); i++ {
		rest := M - i
		if rest < 0 {
			continue
		}
		if rest >= len(b1) {
			continue
		}
		key := (rest + K - 1) / K
		cost := key*Q + b0[i] + b1[rest]
		// out(i, rest, key, cost, "B", b0[i], b1[rest])
		ans = min(ans, cost)
	}
	out(ans)
}
