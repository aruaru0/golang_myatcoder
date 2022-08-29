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
var idx []int
var n []int

func dfs(cur, prev, num int) {
	idx[cur] = num
	for _, e := range node[cur] {
		if n[e] != 0 {
			continue
		}
		if e == prev {
			continue
		}
		dfs(e, cur, num)
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node = make([][]int, N)
	for i := 0; i < N; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	n = make([]int, N)
	q := []int{}
	for i := 0; i < N; i++ {
		n[i] = len(node[i])
		if n[i] == 1 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		cur := q[0]
		n[cur]--
		q = q[1:]
		for _, e := range node[cur] {
			if n[e] != 0 {
				n[e]--
			}
			if n[e] == 1 {
				q = append(q, e)
			}
		}
	}

	idx = make([]int, N)
	for i := 0; i < N; i++ {
		if n[i] == 0 {
			continue
		}
		dfs(i, -1, i)
	}
	Q := getI()
	for i := 0; i < Q; i++ {
		u, v := getI()-1, getI()-1
		if idx[u] == idx[v] {
			out("Yes")
		} else {
			out("No")
		}
	}
}
