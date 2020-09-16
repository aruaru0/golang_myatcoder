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
type pqi struct{ cnt, val int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int      { return len(pq) }
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].cnt == pq[j].cnt {
		return pq[i].val > pq[j].val
	}
	return pq[i].cnt > pq[j].cnt
}
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

const maxN = 201000

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	A := make([]int, N)
	a := make([]int, maxN)
	b := make([]int, maxN)
	for i := 0; i < N; i++ {
		x := getInt()
		A[i] = x
		a[x]++
	}
	for i := 0; i < N; i++ {
		x := getInt()
		b[x]++
	}

	pq := priorityQueue{}
	for i := 0; i < maxN; i++ {
		if a[i]+b[i] > N {
			out("No")
			return
		}
		if b[i] != 0 {
			heap.Push(&pq, pqi{b[i], i})
		}
	}

	ans := make([]int, N)
	for i := 0; i < N; i++ {
		p0 := pq[0]
		heap.Pop(&pq)
		if p0.val != A[i] {
			ans[i] = p0.val
			p0.cnt--
			if p0.cnt > 0 {
				heap.Push(&pq, p0)
			}
		} else {
			if len(pq) == 0 {
				for j := 0; j < N-1; j++ {
					// out(A[j], A[i], ans[j], p0.val)
					if A[j] != p0.val && A[i] != ans[j] {
						ans[j], ans[i] = p0.val, ans[j]
						break
					}
				}
				p0.cnt--
				if p0.cnt > 0 {
					heap.Push(&pq, p0)
				}
				continue
			}
			p1 := heap.Pop(&pq).(pqi)
			ans[i] = p1.val
			heap.Push(&pq, p0)
			p1.cnt--
			if p1.cnt > 0 {
				heap.Push(&pq, p1)
			}
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, "Yes")
	for i := 0; i < N; i++ {
		fmt.Fprint(w, ans[i], " ")
	}
	fmt.Fprintln(w)
}
