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

func solve(n, x, m int) {
	tot := 0
	for i := 0; i < n; i++ {
		tot += x
		x = x * x % m
	}
	out("solve", tot)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, X, M := getI(), getI(), getI()
	m := make(map[int]int)
	a := make([]int, 0)

	tot := 0
	f := X
	rest := N
	from, to := 0, 0
	for i := 0; i < N; i++ {
		rest--
		_, ok := m[f]
		if ok {
			tot += f
			a = append(a, tot)
			from = m[f]
			to = i
			break
		}
		m[f] = i
		tot += f
		a = append(a, tot)
		f = f * f % M
	}
	// solve(N, X, M)

	if rest == 0 {
		out(a[len(a)-1])
		return
	}

	// out("rest", rest, from, to)

	n := rest / (to - from)
	r := rest % (to - from)
	// out(n, r)
	sum := a[to] - a[from]
	ans := a[len(a)-1] + sum*n
	ans += a[from+r] - a[from]
	// out(n, r, sum)
	out(ans)
}
