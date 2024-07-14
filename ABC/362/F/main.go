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
	n := getI()

	node := make([][]int, n)
	for i := 0; i < n-1; i++ {
		a, b := getI()-1, getI()-1
		node[a] = append(node[a], b)
		node[b] = append(node[b], a)
	}

	center := 0
	sz := make([]int, n)

	var dfs func(cur, prev int) int
	dfs = func(cur, prev int) int {
		mx := 0
		sz[cur] = 1
		for _, e := range node[cur] {
			if e == prev {
				continue
			}
			sz[cur] += dfs(e, cur)
			mx = max(mx, sz[e])
		}
		mx = max(mx, n-sz[cur])
		if mx*2 <= n {
			center = cur
		}
		return sz[cur]
	}

	dfs(0, -1)

	vs := make([]int, 0)
	var dfs2 func(cur, prev int)
	dfs2 = func(cur, prev int) {
		for _, e := range node[cur] {
			if e == prev {
				continue
			}
			dfs2(e, cur)
		}
		vs = append(vs, cur)
	}
	dfs2(center, -1)
	if n%2 == 1 {
		vs = vs[:len(vs)-1]
	}
	for i := 0; i < n/2; i++ {
		out(vs[i]+1, vs[i+n/2]+1)
	}
}
