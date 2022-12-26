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

var rt []int

func dfs(s, n int, x []int) {
	if used[n] == true {
		if s == n {
			if len(rt) == 0 {
				rt = make([]int, len(x))
				copy(rt, x)
			} else if len(rt) > len(x) {
				rt = make([]int, len(x))
				copy(rt, x)
			}
		}
		return
	}
	used[n] = true
	for _, e := range node[n] {
		dfs(s, e, append(x, e))
	}
}

var N, M int
var node [][]int
var inode [][]int
var used []bool

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node = make([][]int, N)
	inode = make([][]int, N)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		node[a] = append(node[a], b)
		inode[b] = append(inode[b], a)
	}

	for i := 0; i < N; i++ {
		used = make([]bool, N)
		rt = []int{}
		dfs(i, i, []int{i})
		// out(rt)
		if len(rt) == 0 {
			continue
		}
		if rt[0] != rt[len(rt)-1] {
			continue
		}
		m := make(map[int]bool)
		for _, e := range rt {
			m[e] = true
		}
		ok := true
		for _, e := range rt {
			cnt := 0
			for _, v := range inode[e] {
				if m[v] {
					cnt++
				}
			}
			if cnt > 1 {
				ok = false
				break
			}
		}
		if ok {
			// out(node)
			// out(inode)
			out(len(rt) - 1)
			for i := 0; i < len(rt)-1; i++ {
				out(rt[i] + 1)
			}
			return
		}
	}

	out(-1)
}
