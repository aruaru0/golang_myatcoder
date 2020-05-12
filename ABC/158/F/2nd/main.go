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

type robot struct {
	x, y int
}

type robots []robot

func (p robots) Len() int {
	return len(p)
}

func (p robots) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p robots) Less(i, j int) bool {
	return p[i].x < p[j].x
}

// Priority Queue
// Item :
type Item struct {
	priority, x, index int
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
const mod = 998244353

type pair struct {
	first, second int
}

var to [][]int

func dfs(v int) int {
	res := 1
	for _, u := range to[v] {
		res *= dfs(u)
		res %= mod
	}
	return (res + 1) % mod
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	p := make(robots, N)
	for i := 0; i < N; i++ {
		x, d := getInt(), getInt()
		p[i] = robot{x, d}
	}
	sort.Sort(p)
	// out(p)
	to = make([][]int, N)
	pq := make(PQ, 0)
	heap.Init(&pq)
	for i := N - 1; i >= 0; i-- {
		x := p[i].x
		d := p[i].y
		// out("----")
		for pq.Len() > 0 {
			item := heap.Pop(&pq).(*Item)
			// out("Pop ", item.priority, item.x, "x+d", x, d, x+d)
			if item.priority < x+d {
				to[i] = append(to[i], item.x)
				// out("append ", i, item.x)
			} else {
				heap.Push(&pq, item)
				// out("push back")
				break
			}
		}
		heap.Push(&pq, &Item{x, i, 0})
		// out("Push ", x, i)
	}

	// out(to)
	ans := 1
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		// out(item.x)
		ans *= dfs(item.x)
		ans %= mod
	}
	out(ans)
}
