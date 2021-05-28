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

type pqi struct{ a, i, j, k int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].a > pq[j].a }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

type pos struct {
	i, j, k int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	X, Y, Z, K := getI(), getI(), getI(), getI()
	a := getInts(X)
	b := getInts(Y)
	c := getInts(Z)
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	m := make(map[pos]bool)
	pq := priorityQueue{}
	heap.Push(&pq, pqi{a[0] + b[0] + c[0], 0, 0, 0})
	for i := 0; i < K; i++ {
		p := pq[0]
		heap.Pop(&pq)
		if p.i+1 != len(a) && !m[pos{p.i + 1, p.j, p.k}] {
			m[pos{p.i + 1, p.j, p.k}] = true
			heap.Push(&pq, pqi{a[p.i+1] + b[p.j] + c[p.k], p.i + 1, p.j, p.k})
		}
		if p.j+1 != len(b) && !m[pos{p.i, p.j + 1, p.k}] {
			m[pos{p.i, p.j + 1, p.k}] = true
			heap.Push(&pq, pqi{a[p.i] + b[p.j+1] + c[p.k], p.i, p.j + 1, p.k})
		}
		if p.k+1 != len(c) && !m[pos{p.i, p.j, p.k + 1}] {
			m[pos{p.i, p.j, p.k + 1}] = true
			heap.Push(&pq, pqi{a[p.i] + b[p.j] + c[p.k+1], p.i, p.j, p.k + 1})
		}
		out(p.a)
	}
}
