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

var uf []int

func ufinit(n int) {
	uf = make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = -1
	}
}

func root(x int) int {
	if uf[x] < 0 {
		return x
	}
	return root(uf[x])
}

func unite(x, y int) {
	x = root(x)
	y = root(y)
	if x == y {
		return
	}
	if uf[x] > uf[y] {
		x, y = y, x
	}
	uf[x] += uf[y]
	uf[y] = x
}

func same(x, y int) bool {
	return root(x) == root(y)
}

func size(x int) int {
	return -uf[root(x)]
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	node := make([][]int, N)

	ufinit(N)
	m := make(map[int]int, 0)
	s := 0

	for i := 0; i < M; i++ {
		f, t := getI(), getI()
		node[f] = append(node[f], t)
		node[t] = append(node[t], f)
		unite(f, t)
		m[f]++
		m[t]++
		s = f
	}

	if size(s) != len(m) {
		out("NO")
		return
	}

	even, odd := 0, 0
	for i := 0; i < N; i++ {
		if len(node[i])%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if odd == 2 || odd == 0 {
		out("YES")
		return
	}
	out("NO")
	return
}
