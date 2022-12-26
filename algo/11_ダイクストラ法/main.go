package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
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
func Dijkstra(N, S int, path []Path) ([]int, []Route) {
	pq := make(PQ, 0)
	heap.Init(&pq)
	d := make([]int, N+1)
	r := make([]Route, N+1)
	// init
	for i := 0; i <= N; i++ {
		d[i] = math.MaxInt32
	}
	d[S] = 0
	r[S].path = []int{S}
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
				r[e.to].path = append(r[v].path, e.to)
				heap.Push(&pq, &Item{d[e.to], e.to, 0})
			}
		}
	}
	return d, r

}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	path := make([]Path, N+1)

	for i := 0; i < M; i++ {
		f, t, c := getInt(), getInt(), getInt()
		path[f].edges = append(path[f].edges, Edge{t, c})
		path[t].edges = append(path[t].edges, Edge{f, c})
	}

	S, D := getInt(), getInt()

	d, r := Dijkstra(N, S, path)

	out("cost = ", d[D])
	for i, v := range r[D].path {
		fmt.Print(v)
		if i != len(r[D].path)-1 {
			fmt.Print("->")
		}
	}
	fmt.Print("\n")
}
