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

const inf = math.MaxInt64

var node [][]edge
var dist []int

func dijkstra(s int) {
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, s})
	dist[s] = 0
	for len(pq) != 0 {
		cur := pq[0].to
		heap.Pop(&pq)
		for _, e := range node[cur] {
			if dist[e.to] > max(dist[cur], e.cost) {
				dist[e.to] = max(dist[cur], e.cost)
				heap.Push(&pq, pqi{dist[e.to], e.to})
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt(), getInt()
	}

	node = make([][]edge, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			dx := abs(x[i] - x[j])
			dy := abs(y[i] - y[j])
			node[i] = append(node[i], edge{j, dx*dx + dy*dy})
		}
	}

	dist = make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	dijkstra(0)
	// out(dist)
	ans := int(math.Sqrt(float64(dist[N-1])))
	if ans*ans < dist[N-1] {
		ans++
	}
	ans = (ans + 9) / 10 * 10
	out(ans)
}
