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
	to int
	c  byte
}

var N, M int
var node [][]edge
var dist [][]int

const inf = int(1e15)

func bfs(l, r int) int {
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = inf
		}
	}

	dist[0][N-1] = 0
	q := [][2]int{{0, N - 1}}

	ret := inf
	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		x, y := c[0], c[1]
		if x == y {
			ret = min(ret, dist[x][y]*2)
		}
		for _, e := range node[x] {
			for _, v := range node[y] {
				if e.to == y && v.to == x {
					ret = min(ret, dist[x][y]*2+1)
				}
				if dist[e.to][v.to] != inf {
					continue
				}
				if e.c != v.c {
					continue
				}
				dist[e.to][v.to] = dist[x][y] + 1
				q = append(q, [2]int{e.to, v.to})
			}
		}
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node = make([][]edge, N)
	for i := 0; i < M; i++ {
		f, t, c := getI()-1, getI()-1, []byte(getS())
		node[f] = append(node[f], edge{t, c[0]})
		node[t] = append(node[t], edge{f, c[0]})
	}

	ret := bfs(0, N-1)

	if ret == inf {
		out(-1)
		return
	}
	out(ret)
}
