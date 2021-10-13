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

type pqi struct{ cost, to int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

type edge struct {
	to, idx int
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, S := getI(), getI(), getI()
	S = min(2999, S)
	node := make([][]edge, N)
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		u, v := getI()-1, getI()-1
		A[i], B[i] = getI(), getI()
		node[u] = append(node[u], edge{v, i})
		node[v] = append(node[v], edge{u, i})
	}
	C := make([]int, N)
	D := make([]int, N)
	for i := 0; i < N; i++ {
		C[i], D[i] = getI(), getI()
	}

	dist := make([][]int, N)
	vis := make([][]bool, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, 3000)
		vis[i] = make([]bool, 3000)
		for j := 0; j < 3000; j++ {
			dist[i][j] = inf
		}
	}
	ans := make([]int, N)
	for i := 0; i < N; i++ {
		ans[i] = inf
	}

	pq := priorityQueue{}
	dist[0][S] = 0
	heap.Push(&pq, pqi{0, 0*5010 + S})
	for len(pq) != 0 {
		q := pq[0]
		heap.Pop(&pq)
		cst := q.cost
		cu := q.to / 5010
		s := q.to % 5010
		if vis[cu][s] {
			continue
		}
		vis[cu][s] = true
		chmin(&ans[cu], cst)
		for _, p := range node[cu] {
			if s < A[p.idx] {
				continue
			}
			cost := cst + B[p.idx]
			if dist[p.to][s-A[p.idx]] > cost {
				dist[p.to][s-A[p.idx]] = cost
				heap.Push(&pq, pqi{dist[p.to][s-A[p.idx]], p.to*5010 + s - A[p.idx]})
			}
		}
		cost := cst + D[cu]
		s2 := min(2999, s+C[cu])
		if dist[cu][s2] > cost {
			dist[cu][s2] = cost
			heap.Push(&pq, pqi{dist[cu][s2], cu*5010 + s2})

		}
	}

	for i := 1; i < N; i++ {
		out(ans[i])
	}
}
