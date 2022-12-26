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

const inf int = 1e18

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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()[]
	N, M, K := getI(), getI(), getI()

	node := make([][]int, N*2)
	for i := 0; i < M; i++ {
		u, v, a := getI()-1, getI()-1, getI()
		if a == 0 {
			u += N
			v += N
		}
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}
	for i := 0; i < K; i++ {
		s := getI() - 1
		node[s] = append(node[s], s+N)
		node[s+N] = append(node[s+N], s)
	}

	dist := make([]int, N*2)
	for i := 0; i < N*2; i++ {
		dist[i] = inf
	}
	dist[0] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0})
	for len(pq) != 0 {
		cur := pq[0].a
		heap.Pop(&pq)
		for _, e := range node[cur] {
			if e%N == cur%N {
				if dist[e] > dist[cur] {
					dist[e] = dist[cur]
					heap.Push(&pq, pqi{e})
				}
			} else {
				if dist[e] > dist[cur]+1 {
					dist[e] = dist[cur] + 1
					heap.Push(&pq, pqi{e})
				}
			}
		}
	}
	ans := min(dist[N-1], dist[2*N-1])
	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}
}
