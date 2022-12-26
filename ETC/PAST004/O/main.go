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

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ a, to int }

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

type edge struct {
	to, cost int
}

var node [][]edge
var dist []int
var used []bool

func dijkstra(v int) {
	pq := priorityQueue{}
	dist[v] = 0
	heap.Push(&pq, pqi{0, v})
	for len(pq) != 0 {
		cur := pq[0].to
		heap.Pop(&pq)
		if used[cur] {
			continue
		}
		used[cur] = true
		for _, e := range node[cur] {
			if dist[e.to] > dist[cur]+e.cost {
				dist[e.to] = dist[cur] + e.cost
				heap.Push(&pq, pqi{dist[e.to], e.to})
			}
		}
	}
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	a := getInts(N)
	tot := 0
	for i := 0; i < N; i++ {
		tot += a[i]
	}
	node = make([][]edge, N+1)
	for i := 0; i < N; i++ {
		to, from := i+1, i
		node[from] = append(node[from], edge{to, a[i]})
		node[to] = append(node[to], edge{from, 0})
	}
	for i := 0; i < M; i++ {
		l, r, c := getI()-1, getI(), getI()
		node[l] = append(node[l], edge{r, c})
	}
	dist = make([]int, N+1)
	for i := 0; i <= N; i++ {
		dist[i] = inf
	}
	used = make([]bool, N+1)
	dijkstra(0)
	out(tot - dist[N])
}
