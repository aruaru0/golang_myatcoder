package main

import (
	"bufio"
	"container/heap"
	"fmt"
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
	to, cost int
}

var N, M, P, Q, T int
var node [][]edge

const inf = int(1e18)

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ a, pos int }

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

func dijkstra(s int) []int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	dist[s] = 0
	heap.Push(&pq, pqi{0, s})
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		for _, e := range node[cur.pos] {
			if dist[e.to] > dist[cur.pos]+e.cost {
				dist[e.to] = dist[cur.pos] + e.cost
				heap.Push(&pq, pqi{dist[e.to], e.to})
			}
		}
	}
	return dist
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, P, Q, T = getI(), getI(), getI()-1, getI()-1, getI()
	node = make([][]edge, N)
	for i := 0; i < M; i++ {
		from, to, cost := getI()-1, getI()-1, getI()
		node[from] = append(node[from], edge{to, cost})
		node[to] = append(node[to], edge{from, cost})
	}

	s := dijkstra(0)
	p := dijkstra(P)
	q := dijkstra(Q)

	ans := -1

	d := s[P] + s[Q] + p[Q]
	if d <= T {
		ans = T
	}

	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			d0 := s[x] + p[x] + p[y] + s[y]
			d1 := s[x] + q[x] + q[y] + s[y]
			// out(x, d0, d1, s[x])
			if max(d0, d1) <= T {
				ans = max(ans, s[x]+s[y]+T-max(d0, d1))
			}

		}
	}
	// out(s, p, q)
	out(ans)
}
