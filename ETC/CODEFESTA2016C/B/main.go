package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
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

func getString() string {
	sc.Scan()
	return sc.Text()
}

// Priority Queue
type Item struct {
	priority, value, index int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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

func main() {
	sc.Split(bufio.ScanWords)

	K, T := getInt(), getInt()
	pq := make(PQ, 0)
	heap.Init(&pq)

	for i := 0; i < T; i++ {
		a := getInt()
		heap.Push(&pq, &Item{a, i, 0})
	}

	if K == 1 {
		out(0)
		return
	}

	item0 := heap.Pop(&pq).(*Item)
	item0.priority--

	for pq.Len() > 0 {
		item1 := heap.Pop(&pq).(*Item)
		item1.priority--
		if item0.priority != 0 {
			heap.Push(&pq, &Item{item0.priority, item0.value, 0})
		}
		item0 = item1
	}

	ans := item0.priority
	out(ans)
}
