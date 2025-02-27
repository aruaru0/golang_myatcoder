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
	if len(s) == 0 {
		return
	}
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

type pair struct {
	s, t int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	c := make([]string, n)
	for i := 0; i < n; i++ {
		c[i] = getS()
	}

	const inf = int(1e9 + 10)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = inf
		}
	}

	q := make([]pair, 0)
	push := func(s, t, d int) {
		if dist[s][t] != inf {
			return
		}
		dist[s][t] = d
		q = append(q, pair{s, t})
	}

	for i := 0; i < n; i++ {
		push(i, i, 0)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if c[i][j] != '-' {
				push(i, j, 1)
			}
		}
	}

	for len(q) != 0 {
		s, t := q[0].s, q[0].t
		q = q[1:]
		for ns := 0; ns < n; ns++ {
			for nt := 0; nt < n; nt++ {
				if c[ns][s] == '-' {
					continue
				}
				if c[t][nt] == '-' {
					continue
				}
				if c[ns][s] != c[t][nt] {
					continue
				}
				push(ns, nt, dist[s][t]+2)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans := dist[i][j]
			if ans == inf {
				ans = -1
			}
			fmt.Fprint(wr, ans, " ")
		}
		out()
	}
}
