package main

import (
	"container/heap"
	"fmt"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

// Priority Queue
// Item :
type Item struct {
	priority, value, index int
}

// PQ :
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

// Push :
func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop :
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// End Priority Queue

func main() {
	pq := make(PQ, 0)

	heap.Init(&pq)
	heap.Push(&pq, &Item{10, 2, 0})
	heap.Push(&pq, &Item{30, 4, 0})
	heap.Push(&pq, &Item{1, 1, 0})
	heap.Push(&pq, &Item{14, 3, 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		out(item.value, item.priority)
	}
}
