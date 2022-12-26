package main

import (
	"bufio"
	"container/heap"
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

type pqi struct{ c, t int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].c < pq[j].c }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

func pos2idx(r, c, C int) int {
	return r*C + c
}

type edge struct {
	t, c int
}

var node [][]edge
var R, C int

const inf = int(1e18)

func dijkstra(s int) []int {
	dist := make([]int, R*C*2)
	for i := 0; i < R*C*2; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	dist[s] = 0
	heap.Push(&pq, pqi{0, s})
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if dist[cur.t] < cur.c {
			continue
		}
		// out(cur, node[cur.t])
		for _, e := range node[cur.t] {
			if dist[e.t] > dist[cur.t]+e.c {
				dist[e.t] = dist[cur.t] + e.c
				heap.Push(&pq, pqi{dist[e.t], e.t})

			}
		}
	}
	return dist
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	R, C = getI(), getI()
	a := make([][]int, R)
	for i := 0; i < R; i++ {
		a[i] = getInts(C - 1)
	}
	b := make([][]int, R-1)
	for i := 0; i < R-1; i++ {
		b[i] = getInts(C)
	}

	// mirror nodes
	node = make([][]edge, R*C*2)
	off := R * C
	for i := 0; i < C; i++ {
		for j := 1; j < R; j++ {
			from, to := pos2idx(j, i, C), pos2idx(j-1, i, C)
			node[from+off] = append(node[from+off], edge{to + off, 1})
			node[from] = append(node[from], edge{to + off, 1})
		}
		for j := 0; j < R; j++ {
			from := pos2idx(j, i, C)
			node[from+off] = append(node[from+off], edge{from, 1})
		}
	}
	for i := 0; i < R; i++ {
		// ->
		for j := 0; j < C-1; j++ {
			from, to := pos2idx(i, j, C), pos2idx(i, j+1, C)
			node[from] = append(node[from], edge{to, a[i][j]})
		}
		// <-
		for j := 1; j < C; j++ {
			from, to := pos2idx(i, j, C), pos2idx(i, j-1, C)
			node[from] = append(node[from], edge{to, a[i][j-1]})
		}
	}
	for i := 0; i < C; i++ {
		for j := 0; j < R-1; j++ {
			from, to := pos2idx(j, i, C), pos2idx(j+1, i, C)
			node[from] = append(node[from], edge{to, b[j][i]})
		}
	}

	dist := dijkstra(0)
	// out(dist)
	out(dist[R*C-1])
}
