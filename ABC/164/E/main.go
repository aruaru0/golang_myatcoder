package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

type item struct {
	priority, value, index int
}

type pQ []*item

func (pq pQ) Len() int {
	return len(pq)
}

func (pq pQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq pQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pQ) Push(x interface{}) {
	n := len(*pq)
	it := x.(*item)
	it.index = n
	*pq = append(*pq, it)
}

func (pq *pQ) Pop() interface{} {
	old := *pq
	n := len(old)
	it := old[n-1]
	it.index = -1
	*pq = old[0 : n-1]
	return it
}

// End Priority Queue
type link struct {
	u, v, a, b int
}

type edge struct {
	to, cost int
}

type node struct {
	e []edge
}

const inf = 1001001001001001

func dijkstra(s, N int, n []node) []int {
	pq := make(pQ, 0)
	heap.Init(&pq)
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	dist[s] = 0
	heap.Push(&pq, &item{0, s, 0})
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*item)
		v := it.value
		if dist[v] < it.priority {
			continue
		}
		for _, e := range n[v].e {
			if dist[e.to] > dist[v]+e.cost {
				// out(v, e)
				dist[e.to] = dist[v] + e.cost
				heap.Push(&pq, &item{dist[e.to], e.to, 0})
			}
		}
	}
	return dist
}

func pos(n, c, m int) int {
	return n*m + c
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M, S := getInt(), getInt(), getInt()
	ma := 0
	l := make([]link, M*2)
	for i := 0; i < M; i++ {
		u, v, a, b := getInt()-1, getInt()-1, getInt(), getInt()
		l[i*2] = link{u, v, a, b}
		l[i*2+1] = link{v, u, a, b}
		ma = max(ma, a)
	}

	c := make([]int, N)
	d := make([]int, N)
	for i := 0; i < N; i++ {
		c[i], d[i] = getInt(), getInt()
	}

	// make nodes
	ma *= N
	n := make([]node, ma*N)
	for i := 0; i < N; i++ {
		// out("---", ma)
		for j := 0; j <= ma; j++ {
			to := j + c[i]
			cost := d[i]
			if to >= ma {
				break
			}
			f := pos(i, j, ma)
			t := pos(i, to, ma)
			// out(f, "->", t, cost)
			n[f].e = append(n[f].e, edge{t, cost})
		}
	}
	for i := 0; i < M*2; i++ {
		from := l[i].u
		to := l[i].v
		coin := l[i].a
		cost := l[i].b
		for j := 0; j < ma; j++ {
			if j-coin < 0 {
				continue
			}
			f := pos(from, j, ma)
			t := pos(to, j-coin, ma)
			// out(from, to, "f,t", f, t, cost, len(n))
			n[f].e = append(n[f].e, edge{t, cost})
		}
	}

	S = min(S, ma-1)
	dist := dijkstra(S, ma*N, n)

	for i := 1; i < N; i++ {
		m := inf
		for j := 0; j < ma; j++ {
			m = min(m, dist[i*ma+j])
		}
		// out(dist[i*ma : (i+1)*ma])
		out(m)
	}
	//out("dist=", dist)
}
