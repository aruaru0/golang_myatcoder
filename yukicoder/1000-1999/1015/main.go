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
type pqi struct{ a int }

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

func calc(Z, Y int) bool {
	for Z != 0 {
		if len(pq) == 0 {
			return true
		}
		v := heap.Pop(&pq).(pqi).a
		n := v / Y
		m := v % Y
		if n == 0 {
			Z--
		} else {
			if n <= Z {
				Z -= n
				v = m
			} else {
				v -= Z * Y
				Z = 0
			}
			heap.Push(&pq, pqi{v})
		}
	}
	return false
}

var pq priorityQueue

func main() {
	sc.Split(bufio.ScanWords)
	N, X, Y, Z := getInt(), getInt(), getInt(), getInt()
	pq = priorityQueue{}
	for i := 0; i < N; i++ {
		a := getInt()
		if a%1000 == 999 {
			heap.Push(&pq, pqi{a})
		} else {
			heap.Push(&pq, pqi{a + 1})
		}
	}

	// out(pq)
	if calc(Z, 10000) == true {
		out("Yes")
		return
	}
	// out(pq)
	if calc(Y, 5000) == true {
		out("Yes")
		return
	}
	// out(pq)
	if calc(X, 1000) == true {
		out("Yes")
		return
	}
	// out(pq)
	if len(pq) == 0 {
		out("Yes")
		return
	}
	out("No")
}
