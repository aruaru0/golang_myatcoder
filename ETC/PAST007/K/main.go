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

type pqi struct{ to, cost, val int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int      { return len(pq) }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].cost == pq[j].cost {
		return pq[i].val > pq[j].val
	}
	return pq[i].cost < pq[j].cost
}
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

type edge struct {
	to, cost int
}

const inf = 1e18

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	a := getInts(N)
	node := make([][]edge, N)
	for i := 0; i < M; i++ {
		u, v, t := getI()-1, getI()-1, getI()
		node[u] = append(node[u], edge{v, t})
		node[v] = append(node[v], edge{u, t})
	}

	dist := make([]int, N)
	sat := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	dist[0] = 0
	sat[0] = a[0]
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, 0, a[0]})
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		for _, e := range node[cur.to] {
			if dist[e.to] == dist[cur.to]+e.cost {
				sat[e.to] = max(sat[e.to], sat[cur.to]+a[e.to])
			}
			if dist[e.to] > dist[cur.to]+e.cost {
				dist[e.to] = dist[cur.to] + e.cost
				sat[e.to] = sat[cur.to] + a[e.to]
				heap.Push(&pq, pqi{e.to, dist[e.to], sat[e.to]})
			}
		}
	}
	out(sat[N-1])
}
