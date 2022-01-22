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

type edge struct {
	to, c, d int
}

var N, M int
var node [][]edge

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

const inf = int(1e18)

func dijkstra(c int) {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, c})
	dist[c] = 0
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if dist[cur.to] < cur.a {
			continue
		}
		for _, e := range node[cur.to] {
			cost := e.c + e.d/(cur.a+1) // 到着してすぐのコスト
			// 最適な時間を求める
			d := int(math.Sqrt(float64(e.d)))
			l, r := max(1, d-2), d+2
			for i := l; i <= r; i++ { // [l,r]の範囲で最小コストを探索する
				x := e.d / i
				if x > cur.a {
					// x-cur.a --> 待ち時間
					cost = min(cost, x-cur.a+e.c+e.d/(x+1))
				}
			}
			// out(cur, e, cost, x, x-cur.a)
			if dist[e.to] > dist[cur.to]+cost {
				dist[e.to] = dist[cur.to] + cost
				heap.Push(&pq, pqi{dist[e.to], e.to})
			}
		}
	}
	// out(dist)
	if dist[N-1] == inf {
		out(-1)
		return
	}
	out(dist[N-1])
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node = make([][]edge, N)
	for i := 0; i < M; i++ {
		a, b, c, d := getI()-1, getI()-1, getI(), getI()
		node[a] = append(node[a], edge{b, c, d})
		node[b] = append(node[b], edge{a, c, d})
	}
	dijkstra(0)
}
