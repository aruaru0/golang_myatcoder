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

// Data :
type Data struct {
	day, pay int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i].day < p[j].day
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
func search(day, M int, used []bool) (int, bool) {
	for i := M - day; i >= 0; i-- {
		if used[i] == false {
			return i, true
		}
	}
	return -1, false
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()

	var J Datas
	J = make(Datas, N)
	for i := 0; i < N; i++ {
		a, b := getInt(), getInt()
		J[i] = Data{a, b}
	}

	sort.Sort(J)
	// out(J)
	pq := make(PQ, 0)
	heap.Init(&pq)

	idx := 0
	ans := 0
	for i := 1; i <= M; i++ {
		// out(i)
		for idx != len(J) && J[idx].day == i {
			heap.Push(&pq, &Item{J[idx].pay, 0})
			idx++
		}
		if pq.Len() == 0 {
			continue
		}
		item := heap.Pop(&pq).(*Item)
		ans += item.priority
	}
	out(ans)
}
