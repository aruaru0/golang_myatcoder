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

var dist []int
var used []bool
var C map[int]bool
var node [][]int

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ a, n int }

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

func f(v int) {
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, v})
	dist[v] = 0
	for len(pq) != 0 {
		c := pq[0]
		heap.Pop(&pq)
		for _, e := range node[c.n] {
			if dist[e] > dist[c.n]+1 {
				dist[e] = dist[c.n] + 1
				heap.Push(&pq, pqi{dist[e], e})
			}
		}
	}
}

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, K := getI(), getI(), getI()
	h := getInts(N)
	C = make(map[int]bool)
	for i := 0; i < K; i++ {
		C[getI()-1] = true
	}
	node = make([][]int, N)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		if h[a] > h[b] {
			a, b = b, a
		}
		node[a] = append(node[a], b)
	}

	dist = make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	used = make([]bool, N)
	for i := 0; i < N; i++ {
		if C[i] == true {
			f(i)
		}
	}

	for _, e := range dist {
		if e == inf {
			out(-1)
		} else {
			out(e)
		}
	}
}
