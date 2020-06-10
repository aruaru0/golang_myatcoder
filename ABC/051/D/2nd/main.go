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
				r[e.to].path = make([]int, len(r[v].path))
				copy(r[e.to].path, r[v].path)
				r[e.to].path = append(r[e.to].path, e.to)
				heap.Push(&pq, &Item{d[e.to], e.to, 0})
			}
		}
	}
	return d, r

}

type pair struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	path := make([]Path, N)
	e := make(map[pair]int)
	for i := 0; i < M; i++ {
		f, t, c := getInt()-1, getInt()-1, getInt()
		path[f].edges = append(path[f].edges, Edge{t, c})
		path[t].edges = append(path[t].edges, Edge{f, c})
		if f > t {
			e[pair{t, f}] = 0
		} else {
			e[pair{f, t}] = 0
		}
	}

	for i := 0; i < N; i++ {
		_, r := Dijkstra(N, i, path)
		for _, v := range r {
			if len(v.path) <= 1 {
				continue
			}
			p := v.path
			for j := 1; j < len(p); j++ {
				if p[j-1] > p[j] {
					e[pair{p[j], p[j-1]}]++
				} else {
					e[pair{p[j-1], p[j]}]++
				}
			}
		}
	}
	ans := 0
	for _, v := range e {
		if v == 0 {
			ans++
		}
	}
	out(ans)
}
