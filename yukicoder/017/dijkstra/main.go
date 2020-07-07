package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ a int }

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

const inf = 1001001001001

type edge struct {
	to, cost int
}

var node [][]edge
var s []int
var dist []int
var N int

func dijkstra(from int) {
	dist = make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	heap.Push(&pq, pqi{from})
	dist[from] = 0
	for len(pq) > 0 {
		x := heap.Pop(&pq).(pqi)
		for _, e := range node[x.a] {
			if dist[e.to] < dist[x.a]+e.cost {
				continue
			}
			dist[e.to] = dist[x.a] + e.cost
			heap.Push(&pq, pqi{e.to})
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	s = getInts(N)
	M := getInt()
	node = make([][]edge, N)

	for i := 0; i < M; i++ {
		f, t, c := getInt(), getInt(), getInt()
		node[f] = append(node[f], edge{t, c})
		node[t] = append(node[t], edge{f, c})
	}

	ans := inf
	for i := 1; i < N-1; i++ {
		dijkstra(0)
		cost0i := dist[i] + s[i]
		dijkstra(i)
		back := make([]int, N)
		copy(back, dist)
		for j := 1; j < N-1; j++ {
			if i == j {
				continue
			}
			dijkstra(j)
			tot := cost0i + back[j] + s[j] + dist[N-1]
			// out(tot)
			ans = min(ans, tot)
		}
	}
	out(ans)
}
