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

var node [][]int
var used []bool

func dfs(cur int) {
	used[cur] = true
	for _, e := range node[cur] {
		if used[e] {
			continue
		}
		used[e] = true
		dfs(e)
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	m := make(map[int]int)
	node = make([][]int, 1)
	m[1] = 0
	for i := 0; i < N; i++ {
		a, b := getI(), getI()
		v, ok := m[a]
		if ok != true {
			v = len(m)
			m[a] = v
			node = append(node, []int{})
		}
		u, ok := m[b]
		if ok != true {
			u = len(m)
			m[b] = u
			node = append(node, []int{})
		}
		node[v] = append(node[v], u)
		node[u] = append(node[u], v)
	}

	r := make(map[int]int)
	for e := range m {
		r[m[e]] = e
	}

	used = make([]bool, len(node))
	dfs(0)
	// out(m)
	// out(node)
	// out(used)
	ans := 0
	for i := 0; i < len(used); i++ {
		if used[i] == true {
			ans = max(ans, r[i])
		}
	}
	out(ans)
}
