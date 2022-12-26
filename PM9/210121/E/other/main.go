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

func bfs(v int) {
	q := make([]int, 0)
	used[v] = true
	depth[v] = 0
	q = append(q, v)
	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		for _, e := range node[c] {
			if used[e] {
				continue
			}
			used[e] = true
			depth[e] = depth[c] + 1
			q = append(q, e)
		}
	}
}

var node [][]int
var used []bool
var depth []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node = make([][]int, N)
	for i := 0; i < N; i++ {
		s := getS()
		for j := 0; j < N; j++ {
			if s[j] == '1' {
				node[i] = append(node[i], j)
			}
		}
	}

	ans := -1
	for i := 0; i < N; i++ {
		used = make([]bool, N)
		depth = make([]int, N)
		bfs(i)
		ok := true
		m := -1
		for i := 0; i < N; i++ {
			for _, e := range node[i] {
				if abs(depth[i]-depth[e]) != 1 {
					ok = false
				}
			}
			m = max(m, depth[i]+1)
		}
		if ok {
			ans = max(ans, m)
		}
	}
	out(ans)
}
