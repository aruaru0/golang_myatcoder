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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

//---------------------------------------------
// priority queue
//---------------------------------------------
type pqi struct{ x, n int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].n > pq[j].n }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

func solve() {

	N := getInt()
	l := make(map[int]int)
	for i := 0; i < N; i++ {
		x := getInt()
		l[x]++
	}
	pq := priorityQueue{}
	for i, v := range l {
		heap.Push(&pq, pqi{i, v})
	}

	ans := 0
	for len(pq) >= 3 {
		var x [3]pqi
		for i := 0; i < 3; i++ {
			x[i] = heap.Pop(&pq).(pqi)
		}
		ans++
		for i := 0; i < 3; i++ {
			if x[i].n > 1 {
				x[i].n--
				heap.Push(&pq, x[i])
			}
		}
	}
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	T := getInt()
	for i := 0; i < T; i++ {
		solve()
	}
}
