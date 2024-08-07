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

// 最大流を求めるプログラム　Ford-Fulkerson法
type edge struct {
	to, cap, rev int
}

type node struct {
	to []edge
}

// G :
var G []node
var used []bool
var xN int

// initG : グラフの初期化
func initG(N int) {
	G = make([]node, N)
	xN = N
}

func addEdge(from, to, cap int) {
	G[from].to = append(G[from].to, edge{to, cap, len(G[to].to)})
	G[to].to = append(G[to].to, edge{from, 0, len(G[from].to) - 1})
}

func dfs(v, t, f int) int {
	if v == t {
		return f
	}
	used[v] = true
	for i, e := range G[v].to {
		if used[e.to] || e.cap <= 0 {
			continue
		}
		d := dfs(e.to, t, min(f, e.cap))
		if d > 0 {
			G[v].to[i].cap -= d
			G[e.to].to[e.rev].cap += d
			return d
		}
	}
	return 0
}

func fordFulkerson(s, t int) int {
	N := xN
	flow := 0
	for {
		used = make([]bool, N)
		f := dfs(s, t, inf)
		if f == 0 {
			break
		}
		flow += f
	}
	return flow
}

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, N := getI(), getI(), getI()
	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N)
	D := make([]int, N)

	for i := 0; i < N; i++ {
		A[i], B[i], C[i], D[i] = getI()-1, getI()-1, getI()-1, getI()-1
	}

	n := H + W + N*2 + 2
	initG(n)
	start, end := 0, n-1
	h, w := 1, 1+H
	x, y := 1+H+W, 1+H+W+N

	for i := 0; i < H; i++ {
		addEdge(start, h+i, 1)
	}
	for i := 0; i < W; i++ {
		addEdge(w+i, end, 1)
	}
	for i := 0; i < N; i++ {
		for j := A[i]; j <= C[i]; j++ {
			addEdge(h+j, x+i, 1)
		}
		addEdge(x+i, y+i, 1)
		for j := B[i]; j <= D[i]; j++ {
			addEdge(y+i, w+j, 1)
		}
	}

	res := fordFulkerson(start, end)
	out(res)
}
