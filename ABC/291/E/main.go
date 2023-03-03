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

var path [][]int
var used []bool
var node [][]int

func dfs(cur, pnum int) bool {
	out("dfs", cur, node[cur])
	used[cur] = true
	path[pnum] = append(path[pnum], cur)
	for _, e := range node[cur] {
		if used[e] == true {
			return false
		}
		ret := dfs(e, pnum)
		if ret == false {
			return false
		}
	}
	return true
}

type pair struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	node = make([][]int, N)
	num := make([]int, N)
	m := make(map[pair]bool)

	for i := 0; i < M; i++ {
		x, y := getI()-1, getI()-1
		if m[pair{x, y}] {
			continue
		}
		m[pair{x, y}] = true
		node[x] = append(node[x], y)
		num[y]++
	}

	q := []int{}
	for i := 0; i < N; i++ {
		if num[i] == 0 {
			q = append(q, i)
		}
	}

	ans := make([]int, N)
	cnt := 1

	for len(q) != 0 {
		if len(q) != 1 {
			out("No")
			return
		}
		cur := q[0]
		q = q[1:]
		ans[cur] = cnt
		cnt++
		for _, e := range node[cur] {
			num[e]--
			if num[e] == 0 {
				q = append(q, e)
			}
		}
	}

	out("Yes")
	for i := 0; i < N; i++ {
		fmt.Fprint(wr, ans[i], " ")
	}
	out("")
}
