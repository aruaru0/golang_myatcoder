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
type pqi struct{ cost, x, y int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

// N : N
var N int
var cost [][]int
var dist [][]int

const inf = 1001001001

func dijkstra(sx, sy int) {
	// out("------------------------")
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	dist = make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = inf
		}
	}
	dist[sy][sx] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, sx, sy})
	for len(pq) > 0 {
		p := heap.Pop(&pq).(pqi)
		cx := p.x
		cy := p.y

		// out("c", cx, cy)
		for i := 0; i < 4; i++ {
			x := cx + dx[i]
			y := cy + dy[i]
			if x < 0 || x >= N || y < 0 || y >= N {
				continue
			}
			if dist[y][x] > dist[cy][cx]+cost[y][x] {
				// out("  -- ", x, y)
				dist[y][x] = dist[cy][cx] + cost[y][x]
				heap.Push(&pq, pqi{dist[y][x], x, y})
			}
		}
	}

}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	V, Ox, Oy := getInt(), getInt()-1, getInt()-1
	cost = make([][]int, N)
	for i := 0; i < N; i++ {
		cost[i] = getInts(N)
	}

	dijkstra(0, 0)
	ans := dist[N-1][N-1]
	// out(ans, V)
	if ans < V {
		out("YES")
		return
	}

	// out(ans, V)
	if Ox != -1 && Oy != -1 {
		sum := dist[Oy][Ox]
		// out(sum, V)
		if sum >= V {
			out("NO")
			return
		}
		V = (V - sum) * 2
		dijkstra(Ox, Oy)
		ans = dist[N-1][N-1]
		// out(ans, V)
	}

	if ans < V {
		out("YES")
		return
	}
	out("NO")

	// for i := 0; i < N; i++ {
	// 	out(dist[i])
	// }
	// out(V, Ox, Oy)
}
