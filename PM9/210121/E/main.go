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

var node [][]int
var color = []*int{}

func dfs(v, c int) bool {
	color[v] = &c
	for e, x := range node[v] {
		if x != 1 {
			continue
		}
		if color[e] == nil {
			r := dfs(e, c^1)
			if r == false {
				return false
			}
		}
		if *color[e] == c {
			return false
		}
	}
	return true
}

const inf = int(1e15)

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
		node[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			if s[j] == '1' {
				node[i][j] = 1
			} else {
				node[i][j] = inf
			}
		}
	}
	color = make([]*int, N)
	if dfs(0, 0) == false {
		out(-1)
		return
	}
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if node[i][j] > node[i][k]+node[k][j] {
					node[i][j] = node[i][k] + node[k][j]
				}
			}
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ans = max(ans, node[i][j])
		}
	}
	out(ans + 1)
}
