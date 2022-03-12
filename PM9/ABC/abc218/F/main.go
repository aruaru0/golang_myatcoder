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

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ a, idx int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].a < pq[j].a }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

const inf = int(1e18)

func dijkstra(s int, node [][]edge) (int, []int) {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	rt := make([][]int, N)
	pq := priorityQueue{}
	dist[s] = 0
	heap.Push(&pq, pqi{0, s})
	for len(pq) != 0 {
		cur := pq[0].idx
		heap.Pop(&pq)
		for _, e := range node[cur] {
			if dist[e.t] > dist[cur]+1 {
				rt[e.t] = append(rt[e.t], rt[cur]...)
				rt[e.t] = append(rt[e.t], e.idx)
				dist[e.t] = dist[cur] + 1
				heap.Push(&pq, pqi{dist[e.t], e.t})
			}
		}
	}
	return dist[N-1], rt[N-1]
}

func dijkstra2(s int, node [][]int) int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	dist[s] = 0
	heap.Push(&pq, pqi{0, s})
	for len(pq) != 0 {
		cur := pq[0].idx
		heap.Pop(&pq)
		for _, e := range node[cur] {
			if dist[e] > dist[cur]+1 {
				dist[e] = dist[cur] + 1
				heap.Push(&pq, pqi{dist[e], e})
			}
		}
	}
	return dist[N-1]
}

type pair struct {
	s, t int
}

type edge struct {
	t, idx int
}

var N, M int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node := make([][]edge, N)
	e := make([]pair, M)
	for i := 0; i < M; i++ {
		s, t := getI()-1, getI()-1
		node[s] = append(node[s], edge{t, i})
		e[i] = pair{s, t}
	}

	// 0 -->
	dist, rt := dijkstra(0, node)

	route := make(map[int]bool)
	for _, e := range rt {
		route[e] = true
	}

	for i := 0; i < M; i++ {
		if route[i] {
			node := make([][]int, N)
			for j := 0; j < M; j++ {
				if i == j {
					continue
				}
				s, t := e[j].s, e[j].t
				node[s] = append(node[s], t)
			}
			d := dijkstra2(0, node)
			if d == inf {
				out(-1)
			} else {
				out(d)
			}
		} else {
			if dist == inf {
				out(-1)
			} else {
				out(dist)
			}
		}
	}
}
