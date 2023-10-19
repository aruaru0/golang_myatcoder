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

type edge struct {
	to, b, c int
}

type pair struct {
	b, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	// sum(b1)/sum(c1) > sum(b2)/sum(c2)
	// sum(b1)*sum(c2) > sum(b2)*sum(c1)
	node := make([][]edge, N)
	for i := 0; i < M; i++ {
		u, v, b, c := getI()-1, getI()-1, getI(), getI()
		node[u] = append(node[u], edge{v, b, c})
	}

	const inf = 1e18
	f := func(x float64) bool {
		dp := make([]float64, N)
		for i := 0; i < N; i++ {
			dp[i] = -inf
		}
		dp[0] = 0.0
		for i := 0; i < N; i++ {
			for _, e := range node[i] {
				b := float64(e.b)
				c := float64(e.c)
				y := dp[i] + b - c*x
				if dp[e.to] < y {
					dp[e.to] = y
				}
			}
		}
		return dp[N-1] >= 0
	}

	ac, wa := 0.0, 1e4
	for i := 0; i < 100; i++ {
		wj := (ac + wa) / 2
		if f(wj) {
			ac = wj
		} else {
			wa = wj
		}
	}

	out(ac)
}
