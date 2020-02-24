package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//fmt.Println(x...)
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

func getString() string {
	sc.Scan()
	return sc.Text()
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Priority Queue
type Item struct {
	priority, index int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// End Priority Queue

// Priority Queue

type PQ2 []*Item

func (pq PQ2) Len() int {
	return len(pq)
}

func (pq PQ2) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PQ2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ2) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ2) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// End Priority Queue

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N*3)
	for i := 0; i < N*3; i++ {
		a[i] = getInt()
	}

	out(N, a)

	pqf := make(PQ, 0)
	pqt := make(PQ2, 0)
	heap.Init(&pqf)
	heap.Init(&pqt)
	f := 0
	t := 0
	for i := 0; i < N; i++ {
		heap.Push(&pqf, &Item{a[i], 0})
		heap.Push(&pqt, &Item{a[3*N-1-i], 0})
		f += a[i]
		t += a[3*N-1-i]
	}

	nf := make([]int, N+1)
	nt := make([]int, N+1)
	nf[0] = f
	nt[0] = t
	for i := 0; i < N; i++ {
		f += a[N+i]
		heap.Push(&pqf, &Item{a[N+i], 0})
		item := heap.Pop(&pqf).(*Item)
		out("front", item.priority)
		f -= item.priority
		nf[i+1] = f

		t += a[2*N-1-i]
		heap.Push(&pqt, &Item{a[2*N-1-i], 0})
		item = heap.Pop(&pqt).(*Item)
		out("tail", item.priority)
		t -= item.priority
		nt[i+1] = t
	}

	out(nf)
	out(nt)

	ma := -(math.MaxInt64 >> 4)
	for i := 0; i <= N; i++ {
		d := nf[i] - nt[N-i]
		if d > ma {
			ma = d
		}
	}
	fmt.Println(ma)
}
