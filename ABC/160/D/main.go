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

// Priority Queue
type Item struct {
	priority, value, index int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// End Priority Queue

// Path
type Edge struct {
	to, cost int
}

type Path struct {
	edges []Edge
}

// Dijkstra
type Route struct {
	path []int
}

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
	N, X, Y := getInt(), getInt()-1, getInt()-1

	path := make([]Path, N)
	for i := 0; i < N-1; i++ {
		f := i
		t := i + 1
		path[f].edges = append(path[f].edges, Edge{t, 1})
		path[t].edges = append(path[t].edges, Edge{f, 1})
	}
	path[X].edges = append(path[X].edges, Edge{Y, 1})
	path[Y].edges = append(path[Y].edges, Edge{X, 1})

	//out(path)
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		d, _ := Dijkstra(N, i, path)
		//out(d, r)
		for j := 0; j < N; j++ {
			m[d[j]]++
		}
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 1; i < N; i++ {
		out(m[i] / 2)
	}
}
