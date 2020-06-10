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

type circle struct {
	x, y, r int
}

func calc(a, b circle) float64 {
	dx := abs(a.x - b.x)
	dy := abs(a.y - b.y)
	r := abs(a.r + b.r)
	d := dx*dx + dy*dy
	if d < a.r || d < b.r {
		return 0.0
	}

	ret := math.Max(math.Sqrt(float64(dx*dx+dy*dy))-float64(r), 0.0)
	return ret
}

// Priority Queue
type Item struct {
	priority     float64
	value, index int
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

func Dijkstra(N, S int, path [][]float64) []float64 {
	pq := make(PQ, 0)
	heap.Init(&pq)
	d := make([]float64, N)
	// init
	for i := 0; i < N; i++ {
		d[i] = math.MaxFloat64
	}
	d[S] = 0
	heap.Push(&pq, &Item{0, S, 0})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		v := item.value
		if d[v] < item.priority {
			continue
		}
		for i := 0; i < len(path); i++ {
			if d[i] > d[v]+path[i][v] {
				d[i] = d[v] + path[i][v]
				heap.Push(&pq, &Item{d[i], i, 0})
			}
		}
	}
	return d

}

func main() {
	sc.Split(bufio.ScanWords)
	xs, ys, xt, yt := getInt(), getInt(), getInt(), getInt()
	N := getInt()
	c := make([]circle, N+2)
	c[0] = circle{xs, ys, 0}
	for i := 1; i <= N; i++ {
		c[i] = circle{getInt(), getInt(), getInt()}
	}
	c[N+1] = circle{xt, yt, 0}

	// out(xs, ys, xt, yt)
	// out(N, c)
	N = N + 2

	t := make([][]float64, N)
	for i := 0; i < N; i++ {
		t[i] = make([]float64, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			t[i][j] = calc(c[i], c[j])
			t[j][i] = t[i][j]
			// out(i, j, t[i][j])
			// out(j, i, t[j][i])
		}
	}

	// for i := 0; i < N; i++ {
	// 	out(t[i])
	// }
	// out("----")
	cost := Dijkstra(N, 0, t)
	out(cost[N-1])

}
