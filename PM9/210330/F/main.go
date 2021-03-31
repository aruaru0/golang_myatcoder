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
	from, to, cost int
}

const inf = int(1e18)

// ベルマンフォード法
// SからEで、Eが関連するネガティブループを見つける
// N : num of node
// S : start
// E : end
// edges : edge struct
// return
// distance [] int
// loop flag bool
func bellmanFord(N, S, E int, edges []edge) ([]int, bool) {
	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = inf
	}
	d[S] = 0
	for i := 0; i < N-1; i++ {
		for _, e := range edges {
			if d[e.from] != inf && d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
			}
		}
	}

	negative := make([]bool, N)
	for i := 0; i < N; i++ {
		for _, e := range edges {
			if d[e.from] != inf && d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
				negative[e.to] = true
			}
			negative[e.to] = negative[e.to] || negative[e.from]
		}
	}

	return d, negative[E]
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	edges := make([]edge, M)
	for i := 0; i < M; i++ {
		f, t, c := getI()-1, getI()-1, getI()
		edges[i] = edge{f, t, -c}
	}

	d, ok := bellmanFord(N, 0, N-1, edges)

	if ok == true {
		out("inf")
		return
	}
	out(-d[N-1])
}
