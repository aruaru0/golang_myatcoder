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

type edge struct {
	to, x, y int
}

var node [][]edge
var pos [][2]int
var used []bool

func dfs(cur int) {
	used[cur] = true
	px, py := pos[cur][0], pos[cur][1]
	for _, e := range node[cur] {
		if used[e.to] == true {
			continue
		}
		pos[e.to][0] = e.x + px
		pos[e.to][1] = e.y + py
		dfs(e.to)
	}
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	node = make([][]edge, N)
	for i := 0; i < M; i++ {
		a, b, x, y := getI()-1, getI()-1, getI(), getI()
		node[a] = append(node[a], edge{b, x, y})
		node[b] = append(node[b], edge{a, -x, -y})
	}

	pos = make([][2]int, N)
	for i := 0; i < N; i++ {
		pos[i][0] = inf
		pos[i][1] = inf
	}

	pos[0][0] = 0
	pos[0][1] = 0
	used = make([]bool, N)

	dfs(0)

	for i := 0; i < N; i++ {
		if used[i] == false {
			out("undecidable")
		} else {
			out(pos[i][0], pos[i][1])
		}
	}
}
