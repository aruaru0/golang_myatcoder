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

func out(x ...interface{}) { fmt.Println(x...) }
func getInt() int          { sc.Scan(); i, _ := strconv.Atoi(sc.Text()); return i }
func getString() string    { sc.Scan(); return sc.Text() }

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

func main() {
	sc.Split(bufio.ScanWords)

	pq := priorityQueue{}

	heap.Push(&pq, pqi{10})
	heap.Push(&pq, pqi{2})
	heap.Push(&pq, pqi{100})
	heap.Push(&pq, pqi{5})
	out(len(pq), pq)
	a := heap.Pop(&pq).(pqi)
	out(a.a)
	i := pq[0]
	out(i.a)
	out(len(pq), pq)
}
