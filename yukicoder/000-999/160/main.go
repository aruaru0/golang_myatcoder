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
type pqi struct {
	a, to int
}

type priorityQueue []pqi

func (pq priorityQueue) Len() int      { return len(pq) }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].a == pq[j].a {
		return pq[i].to < pq[j].to
	}
	return pq[i].a < pq[j].a
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

var N, M, S, G int
var node [][]edge

const inf = int(1e15)

func dijkstra(s int) []int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	dist[s] = 0
	heap.Push(&pq, pqi{0, s})
	for len(pq) != 0 {
		e := pq[0].to
		// out(e, pq)
		heap.Pop(&pq)
		for _, v := range node[e] {
			if dist[v.to] > dist[e]+v.cost {
				dist[v.to] = dist[e] + v.cost
				heap.Push(&pq, pqi{dist[v.to], v.to})
			}
		}
	}
	return dist
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M, S, G = getInt(), getInt(), getInt(), getInt()
	node = make([][]edge, N)
	for i := 0; i < M; i++ {
		a, b, c := getInt(), getInt(), getInt()
		node[a] = append(node[a], edge{b, c})
		node[b] = append(node[b], edge{a, c})
	}
	for i := 0; i < N; i++ {
		sort.Slice(node[i], func(j, k int) bool {
			return node[i][j].to < node[i][k].to
		})
	}
	// for i := 0; i < N; i++ {
	// 	out(i, node[i])
	// }

	dist := dijkstra(G)
	// out(dist)
	cur := S
	cost := dist[S]
	for {
		fmt.Print(cur, " ")
		for _, e := range node[cur] {
			if cost == dist[e.to]+e.cost {
				cost -= e.cost
				cur = e.to
				break
			}
		}
		if cur == G {
			fmt.Println(cur)
			break
		}
	}

}
