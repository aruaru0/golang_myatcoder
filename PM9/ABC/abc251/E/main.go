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

type pqi struct{ a, cost int }

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
var N int

const inf = int(1e18)

func dijkstra(s int) []int {
	dist := make([]int, 2*N)
	for i := 0; i < 2*N; i++ {
		dist[i] = inf
	}

	pq := priorityQueue{}
	heap.Push(&pq, pqi{s, 0})
	dist[s] = 0
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if dist[cur.a] < cur.cost {
			continue
		}
		for _, e := range node[cur.a] {
			if dist[e.to] > dist[cur.a]+e.cost {
				dist[e.to] = dist[cur.a] + e.cost
				heap.Push(&pq, pqi{e.to, dist[e.to]})

			}
		}
	}
	return dist

}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	a := getInts(N)

	node = make([][]edge, 2*N)
	for i := 0; i < N; i++ {
		node[i] = append(node[i], edge{N + (i+1)%N, a[i]})
		node[N+i] = append(node[N+i], edge{i, 0})
		node[N+i] = append(node[N+i], edge{(i + 1) % N, 0})
	}
	dist := dijkstra(0)
	ans := dist[2*N-1]
	dist = dijkstra(1)
	ans = min(ans, dist[N])
	dist = dijkstra(2)
	ans = min(ans, dist[N+1])
	out(ans)
}
