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

func lowerBound(a []int, x int) int {
	l := 0
	r := len(a) - 1
	for l <= r {
		m := (l + r) / 2
		if a[m] >= x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(a []int, x int) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m] <= x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
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

type job struct {
	a, b int
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	jo := make([]job, N)
	for i := 0; i < N; i++ {
		a, b := getInt(), getInt()
		jo[i] = job{a, b}
	}
	sort.Slice(jo, func(i, j int) bool {
		return jo[i].a < jo[j].a
	})
	// out(jo)

	pq := make(PQ, 0)
	heap.Init(&pq)
	c := make([]int, N)
	idx := 0
	for i := 0; i < N; i++ {
		for idx != len(jo) {
			if jo[idx].a <= i+1 {
				heap.Push(&pq, &Item{jo[idx].b, 0})
				idx++
			} else {
				break
			}
		}
		if pq.Len() > 0 {
			item := heap.Pop(&pq).(*Item)
			c[i] = item.priority
		}
	}
	// out(c)
	cnt := 0
	for _, v := range c {
		cnt += v
		out(cnt)
	}
}
