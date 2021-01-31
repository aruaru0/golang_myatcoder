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

var node [][]int
var dist []int

const inf = int(1e15)

type pqi struct{ a, pos int }

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

func dijkstra(v int) {
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, v})
	dist[v] = 0
	for len(pq) != 0 {
		c := pq[0]
		heap.Pop(&pq)
		for _, e := range node[c.pos] {
			if dist[e] > dist[c.pos]+1 {
				dist[e] = dist[c.pos] + 1
				heap.Push(&pq, pqi{dist[e], e})
			}
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	node = make([][]int, N)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		node[a] = append(node[a], b)
		node[b] = append(node[b], a)
	}

	K := getI()
	c := make([]int, K)
	for i := 0; i < K; i++ {
		c[i] = getI() - 1
	}

	di := make([][]int, N)
	for from := 0; from < K; from++ {
		di[from] = make([]int, N)
		dist = di[from]
		for i := 0; i < N; i++ {
			dist[i] = inf
		}
		dijkstra(c[from])
	}

	dx := make([][]int, K)
	for from := 0; from < K; from++ {
		dx[from] = make([]int, K)
		for to := 0; to < K; to++ {
			dx[from][to] = di[from][c[to]]
		}
	}

	k := 1 << K
	dp := make([][]int, K)
	for i := 0; i < K; i++ {
		dp[i] = make([]int, k)
		for j := 0; j < k; j++ {
			dp[i][j] = inf
		}
		dp[i][1<<i] = 0
	}
	for bit := 0; bit < k; bit++ {
		for from := 0; from < K; from++ {
			if (bit>>from)&1 == 0 {
				continue
			}
			for to := 0; to < K; to++ {
				if (bit>>to)&1 == 1 {
					continue
				}
				dp[to][bit|(1<<to)] = min(dp[to][bit|(1<<to)], dp[from][bit]+dx[from][to])
			}
		}
	}

	ans := inf
	for i := 0; i < K; i++ {
		ans = min(ans, dp[i][k-1])
	}
	if ans == inf {
		out(-1)
		return
	}
	out(ans + 1)
}
