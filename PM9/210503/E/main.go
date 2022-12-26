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

type pair struct {
	v, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		a := getI()
		m[a]++
	}

	a := make([]pair, 0)
	for i, e := range m {
		a = append(a, pair{i, e})
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].v > a[j].v
	})

	a2 := make([]int, 0)
	for x := 2; x <= 1e9*2; x *= 2 {
		a2 = append(a2, x)
	}

	// out(a)
	ans := 0
	for _, aa := range a {
		x := aa.v
		e := aa.c
		if e == 0 {
			continue
		}
		pos := upperBound(a2, x)
		v := a2[pos] - x

		if x == v {
			d := m[x] / 2
			// out("x, v", x, v, m[x], m[v], d)
			ans += d
			m[x] %= 2
		} else {
			d := min(m[x], m[v])
			// out("x, v", x, v, m[x], m[v], d)
			ans += d
			m[x] -= d
			m[v] -= d
		}
	}
	out(ans)
}
