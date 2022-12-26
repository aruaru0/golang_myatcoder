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

type pqi struct{ c, to int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].c < pq[j].c }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

type edge struct {
	to, cost int
}

var N, M, R int
var r []int
var node [][]edge

const inf = int(1e18)

func dijkstra(n int) []int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	dist[n] = 0
	heap.Push(&pq, pqi{0, n})
	for len(pq) != 0 {
		c := pq[0]
		heap.Pop(&pq)
		if c.c > dist[c.to] {
			continue
		}
		for _, e := range node[c.to] {
			if dist[e.to] > dist[c.to]+e.cost {
				dist[e.to] = dist[c.to] + e.cost
				heap.Push(&pq, pqi{dist[e.to], e.to})
			}
		}
	}
	ret := make([]int, 0)
	for i := 0; i < R; i++ {
		ret = append(ret, dist[r[i]])
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, R = getI(), getI(), getI()
	r = make([]int, R)
	for i := 0; i < R; i++ {
		r[i] = getI() - 1
	}
	node = make([][]edge, N)
	for i := 0; i < M; i++ {
		from, to, cost := getI()-1, getI()-1, getI()
		node[from] = append(node[from], edge{to, cost})
		node[to] = append(node[to], edge{from, cost})
	}

	d := make([][]int, R)
	for i := 0; i < R; i++ {
		d[i] = dijkstra(r[i])
	}

	n := 1 << R
	dp := make([][]int, R)
	for i := 0; i < R; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}
	for i := 0; i < R; i++ {
		dp[i][1<<i] = 0
	}

	for bit := 0; bit < n; bit++ {
		for from := 0; from < R; from++ {
			if (bit>>from)%2 == 0 {
				continue
			}
			for to := 0; to < R; to++ {
				if (bit>>to)%2 == 1 {
					continue
				}
				// out(from, "->", to, bit)
				chmin(&dp[to][bit|(1<<to)], dp[from][bit]+d[from][to])
			}
		}
	}

	// for i := 0; i < R; i++ {
	// 	out(dp[i])
	// }
	ans := inf
	for i := 1; i < R; i++ {
		ans = min(ans, dp[i][n-1])
	}
	out(ans)
}
