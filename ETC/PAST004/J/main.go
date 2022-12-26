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

type edge struct {
	to, cost int
}

var N, M int
var node [][]edge
var dist []int
var used []int

const inf = math.MaxInt64

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

func dijkstra(n int) {
	dist[n] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, n})
	for len(pq) != 0 {
		v := pq[0].to
		heap.Pop(&pq)
		if used[v] == 1 {
			continue
		}
		used[v] = 1
		for _, e := range node[v] {
			if dist[e.to] > dist[v]+e.cost {
				dist[e.to] = dist[v] + e.cost
				heap.Push(&pq, pqi{dist[e.to], e.to})
			}
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	Xab, Xac, Xbc := getI(), getI(), getI()
	s := getS()
	node = make([][]edge, N+6)
	for i := 0; i < M; i++ {
		f, t, c := getI()-1, getI()-1, getI()
		node[f] = append(node[f], edge{t, c})
		node[t] = append(node[t], edge{f, c})
	}

	ab, ba, ac, ca, bc, cb := N, N+1, N+2, N+3, N+4, N+5

	XabI := Xab/2 + Xab%2
	XacI := Xac/2 + Xac%2
	XbcI := Xbc/2 + Xbc%2
	XabO := Xab / 2
	XacO := Xac / 2
	XbcO := Xbc / 2

	for i := 0; i < N; i++ {
		switch s[i] {
		case 'A':
			node[i] = append(node[i], edge{ab, XabI})
			node[i] = append(node[i], edge{ac, XacI})
			node[ba] = append(node[ba], edge{i, XabO})
			node[ca] = append(node[ca], edge{i, XacO})
		case 'B':
			node[i] = append(node[i], edge{ba, XabI})
			node[i] = append(node[i], edge{bc, XbcI})
			node[ab] = append(node[ab], edge{i, XabO})
			node[cb] = append(node[cb], edge{i, XbcO})
		case 'C':
			node[i] = append(node[i], edge{ca, XacI})
			node[i] = append(node[i], edge{cb, XbcI})
			node[ac] = append(node[ac], edge{i, XacO})
			node[bc] = append(node[bc], edge{i, XbcO})
		}
	}

	used = make([]int, N+6)
	dist = make([]int, N+6)
	for i := 0; i < N+6; i++ {
		dist[i] = inf
	}
	dijkstra(0)
	// out(dist)
	out(dist[N-1])
}
