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

func f(i, j, k, N int, v []int) int {
	if i+j > N {
		return 0
	}
	if i+j < k {
		return 0
	}
	// out(i, j, k)

	pq := make(PQ, 0)
	heap.Init(&pq)
	ret := 0
	// out left
	for x := 0; x < i; x++ {
		heap.Push(&pq, &Item{v[x], 0})
		ret += v[x]
	}
	// out right
	for x := 0; x < j; x++ {
		heap.Push(&pq, &Item{v[len(v)-1-x], 0})
		ret += v[len(v)-1-x]
	}
	// drop
	// out("drop")
	for x := 0; x < k; x++ {
		item := heap.Pop(&pq).(*Item)
		// out(item.priority)
		ret -= item.priority
	}
	// out("--------", ret)

	return ret
}

const inf = 1001001001001

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	v := make([]int, N)
	for i := 0; i < N; i++ {
		v[i] = getInt()
	}

	ans := -inf

	for i := 0; i <= K; i++ {
		for j := 0; j <= K-i; j++ {
			x := K - i - j
			for k := 0; k <= x; k++ {
				ret := f(i, j, k, N, v)
				// out("ret", ret)
				ans = max(ans, ret)
			}
		}
	}
	out(ans)
}
