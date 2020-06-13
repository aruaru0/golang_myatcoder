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

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
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

// Priority Queue
// Item :
type Item struct {
	priority, value, index int
}

// PQ :
type PQ []*Item

// Len :
func (pq PQ) Len() int {
	return len(pq)
}

// Less :
func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap :
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push :
func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop :
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// End Priority Queue

// Edge :
type Edge struct {
	to, cost int
}

// Path :
type Path struct {
	edges []Edge
}

// Route :
type Route struct {
	path []int
}

// Dijkstra :
func Dijkstra(N, S int, path []Path) []int {
	pq := make(PQ, 0)
	heap.Init(&pq)
	d := make([]int, N)
	// init
	for i := 0; i < N; i++ {
		d[i] = math.MaxInt32
	}
	d[S] = 0
	heap.Push(&pq, &Item{0, S, 0})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		v := item.value
		if d[v] < item.priority {
			continue
		}
		for _, e := range path[v].edges {
			if d[e.to] > d[v]+e.cost {
				d[e.to] = d[v] + e.cost
				heap.Push(&pq, &Item{d[e.to], e.to, 0})
			}
		}
	}
	return d

}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	path := make([]Path, N)

	for i := 0; i < M; i++ {
		f, t := getInt()-1, getInt()-1
		path[f].edges = append(path[f].edges, Edge{t, 1})
		path[t].edges = append(path[t].edges, Edge{f, 1})
	}

	s := getInt()
	K := getInt()
	t := make([]int, K+1)
	t[0] = s
	for i := 1; i <= K; i++ {
		t[i] = getInt()
	}

	for i := 0; i < K; i++ {
		d := Dijkstra(N, i, path)
		out(d)
	}

	n = 1 << uint(K)
	for i := 0; i < n; i++ {
		for k := 0; k < K; k++ {

		}
	}

}
