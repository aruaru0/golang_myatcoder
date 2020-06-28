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

func bitCount(bits int) int {

	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}

var node [][]int
var used []bool
var dist []int
var N int

const inf = 100100101001

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

func dijkstra(x int) {
	dist = make([]int, N+1)
	for i := 0; i <= N; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	heap.Push(&pq, pqi{x})
	dist[x] = 1
	for len(pq) != 0 {
		cur := heap.Pop(&pq).(pqi).a
		for _, to := range node[cur] {
			if dist[to] > dist[cur]+1 {
				dist[to] = dist[cur] + 1
				heap.Push(&pq, pqi{to})
			}
		}
	}
}

func bsf(x int) {
	dist = make([]int, N+1)
	for i := 0; i <= N; i++ {
		dist[i] = inf
	}
	q := make([]int, 0)
	q = append(q, 1)
	dist[1] = 1
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, to := range node[cur] {
			if dist[to] > dist[cur]+1 {
				dist[to] = dist[cur] + 1
				q = append(q, to)
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	node = make([][]int, N+1)
	for i := 1; i < N; i++ {
		n := bitCount(i)
		if i+n <= N {
			node[i] = append(node[i], i+n)
		}
		if i-n >= 0 {
			node[i] = append(node[i], i-n)
		}
	}
	used = make([]bool, N+1)
	//　ダイクストラでも幅優先でもどちらでもＯＫ
	//		dijkstra(1)
	bsf(1)
	if dist[N] == inf {
		out(-1)
		return
	}
	out(dist[N])
}
