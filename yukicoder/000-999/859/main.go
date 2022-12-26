package main

import (
	"bufio"
	"container/heap"
	"fmt"
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
}

type connect struct {
	from, to, cost int
}

var node [][]edge
var dist []int

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ a, e int }

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

func dijkstra(s int) {
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, s})
	dist[s] = 0
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		// out(cur, node[cur.e])
		for _, e := range node[cur.e] {
			if dist[e.to] > dist[cur.e]+e.cost {
				dist[e.to] = dist[cur.e] + e.cost
				heap.Push(&pq, pqi{dist[e.to], e.to})
			}
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	x, y, z := getI(), getI(), getI()
	s0, t0 := getS(), getI()
	s1, t1 := getS(), getI()

	a := []connect{
		{0, 2, 1}, {0, 4, 1}, {2, 4, 1},
		{1, 3, 1}, {1, 5, 1}, {3, 5, 1}}

	// swap
	if t0 > t1 {
		s0, s1 = s1, s0
		t0, t1 = t1, t0
	}

	if s0 == "A" && s1 == "A" {
		a = append(a, connect{0, 6, t0 - 1})
		a = append(a, connect{6, 7, t1 - t0})
		a = append(a, connect{7, 1, x - t1})
		a = append(a, connect{2, 3, y - 1})
		a = append(a, connect{4, 5, z - 1})
	} else if s0 == "B" && s1 == "B" {
		a = append(a, connect{2, 6, t0 - 1})
		a = append(a, connect{6, 7, t1 - t0})
		a = append(a, connect{7, 3, y - t1})
		a = append(a, connect{0, 1, x - 1})
		a = append(a, connect{4, 5, z - 1})
	} else if s0 == "C" && s1 == "C" {
		a = append(a, connect{4, 6, t0 - 1})
		a = append(a, connect{6, 7, t1 - t0})
		a = append(a, connect{7, 5, z - t1})
		a = append(a, connect{0, 1, x - 1})
		a = append(a, connect{2, 3, y - 1})
	} else {
		flgA := false
		flgB := false
		flgC := false
		switch s0 {
		case "A":
			a = append(a, connect{0, 6, t0 - 1})
			a = append(a, connect{6, 1, x - t0})
			flgA = true
		case "B":
			a = append(a, connect{2, 6, t0 - 1})
			a = append(a, connect{6, 3, y - t0})
			flgB = true
		case "C":
			a = append(a, connect{4, 6, t0 - 1})
			a = append(a, connect{6, 5, z - t0})
			flgC = true
		}
		switch s1 {
		case "A":
			a = append(a, connect{0, 7, t1 - 1})
			a = append(a, connect{7, 1, x - t1})
			flgA = true
		case "B":
			a = append(a, connect{2, 7, t1 - 1})
			a = append(a, connect{7, 3, y - t1})
			flgB = true
		case "C":
			a = append(a, connect{4, 7, t1 - 1})
			a = append(a, connect{7, 5, z - t1})
			flgC = true
		}
		if flgA == false {
			a = append(a, connect{0, 1, x - 1})
		}
		if flgB == false {
			a = append(a, connect{2, 3, y - 1})
		}
		if flgC == false {
			a = append(a, connect{4, 5, z - 1})
		}
	}

	node = make([][]edge, 8)
	for _, e := range a {
		node[e.from] = append(node[e.from], edge{e.to, e.cost})
		node[e.to] = append(node[e.to], edge{e.from, e.cost})
	}
	// out(a)
	// out(node)

	dist = make([]int, 8)
	for i := 0; i < 8; i++ {
		dist[i] = inf
	}
	dijkstra(7)
	// out(dist)
	out(dist[6])
}
