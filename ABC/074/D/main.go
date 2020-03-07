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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Item :
type Item struct {
	priority, value, index int
}

// PQ :
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
func Dijkstra(N, S int, path [][]int) ([]int, []Route) {
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
		for i, c := range path[v] {
			e := Edge{i, c}
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

	N := getInt()
	b := make([][]int, N)
	n := make([][]int, N)
	for i := 0; i < N; i++ {
		b[i] = make([]int, N)
		n[i] = make([]int, N)
		for j := 0; j < N; j++ {
			b[i][j] = getInt()
			n[i][j] = b[i][j]
		}
	}

	ans := 0
LOOP:
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if b[i][j] > b[i][k]+b[k][j] {
					b[i][j] = b[i][k] + b[k][j]
					ans = -1
					break LOOP
				} else if b[i][j] == b[i][k]+b[k][j] {
					// 経由しても距離が同じなら、ルートを消去
					if i != k && j != k {
						n[i][j] = -1
					}
				}
			}
		}
	}

	if ans == 0 {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if n[i][j] != -1 {
					ans += n[i][j]
				}
			}
		}
		out(ans / 2)
	} else {
		out(-1)
	}
}
