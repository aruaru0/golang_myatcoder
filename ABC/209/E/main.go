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

var dist []int
var used []bool
var node [][]int

func dfs(c int) int {
	used[c] = true
	ret := 0
	loop := false
	for _, e := range node[c] {
		if used[e] {
			if dist[e] == -1 {
				loop = true
			}
			continue
		}
		r := dfs(e)
		if r != -1 {
			ret |= dfs(e)
		} else {
			loop = true
		}
	}

	if len(node[c]) == 0 {
		ret = 1
	} else if loop == true && ret == 0 {
		dist[c] = -1
		ret = -1
	}

	dist[c] = ret
	if ret < 0 {
		return ret
	}
	return ret ^ 1
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	front := make(map[string][]int)
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
		front[s[i][:3]] = append(front[s[i][:3]], i)
	}
	node = make([][]int, N)
	for i := 0; i < N; i++ {
		t := s[i][len(s[i])-3:]
		for _, e := range front[t] {
			node[i] = append(node[i], e)
		}
	}

	dist = make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = -1
	}
	used = make([]bool, N)
	for i := 0; i < N; i++ {
		if used[i] {
			continue
		}
		dfs(i)
	}

	// out("dist", dist)
	// out(s)
	// out(front)
	// out(node)
	for i := 0; i < N; i++ {
		if dist[i] < 0 {
			out("Draw")
		} else if dist[i] == 1 {
			out("Takahashi")
		} else {
			out("Aoki")
		}
	}
}
