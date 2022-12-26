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

var N int
var node [][]int

var route []int

func dfs(v, prev int, r []int) {
	if v == N-1 {
		route = r
	}
	for _, e := range node[v] {
		if e == prev {
			continue
		}
		dfs(e, v, append(r, e))
	}
}

var paint []int

func dfs2(v, prev, col int) {
	for _, e := range node[v] {
		if e == prev {
			continue
		}
		if paint[e] == col || paint[e] == 0 {
			paint[e] = col
			dfs2(e, v, col)
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		f, t := getI()-1, getI()-1
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
	}

	dfs(0, -1, []int{0})
	paint = make([]int, N)
	l := (len(route) + 1) / 2
	for i := 0; i < l; i++ {
		paint[route[i]] = 1
	}
	for i := l; i < len(route); i++ {
		paint[route[i]] = -1
	}
	// out(paint)
	dfs2(0, -1, 1)
	dfs2(N-1, -1, -1)
	// out(paint)
	cnt0 := 0
	cnt1 := 0
	for i := 0; i < N; i++ {
		if paint[i] == 1 {
			cnt0++
		} else {
			cnt1++
		}
	}
	if cnt0 > cnt1 {
		out("Fennec")
		return
	}
	out("Snuke")
}
