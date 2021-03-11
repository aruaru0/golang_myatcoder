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

type edge struct {
	to, cost int
	idx      int
}

type pqi struct{ c, to, idx int }

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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	node := make([][]edge, N)
	for i := 0; i < M; i++ {
		f, t, c := getI()-1, getI()-1, getI()
		node[f] = append(node[f], edge{t, c, i})
		node[t] = append(node[t], edge{f, c, i})
	}

	used := make([]bool, M)

	for i := 0; i < N; i++ {
		dist := make([]int, N)
		for j := 0; j < N; j++ {
			dist[j] = inf
		}
		pq := priorityQueue{}
		dist[i] = 0
		heap.Push(&pq, pqi{0, i, -1})
		for len(pq) != 0 {
			c := pq[0]
			heap.Pop(&pq)
			if dist[c.to] < c.c {
				continue
			}
			if c.idx >= 0 {
				used[c.idx] = true
			}
			for _, e := range node[c.to] {
				if dist[e.to] > dist[c.to]+e.cost {
					dist[e.to] = dist[c.to] + e.cost
					heap.Push(&pq, pqi{dist[e.to], e.to, e.idx})
				}
			}
		}
		// out(i, dist, used)
	}
	ans := 0
	for i := 0; i < M; i++ {
		if used[i] == false {
			ans++
		}
	}
	out(ans)
}
