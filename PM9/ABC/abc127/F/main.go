package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

type State = int

type SmallHeap []State

func (q SmallHeap) Less(i, j int) bool    { return q[i] > q[j] }
func (q SmallHeap) Len() int              { return len(q) }
func (q SmallHeap) Swap(i, j int)         { q[i], q[j] = q[j], q[i] }
func (q *SmallHeap) Push(x interface{})   { *q = append(*q, x.(State)) }
func (q *SmallHeap) Pop() (x interface{}) { *q, x = (*q)[:len(*q)-1], (*q)[len(*q)-1]; return }
func (q *SmallHeap) push(v State)         { heap.Push(q, v) }
func (q *SmallHeap) pop() State           { return heap.Pop(q).(State) }

type LargeHeap []State

func (q LargeHeap) Less(i, j int) bool    { return q[i] < q[j] }
func (q LargeHeap) Len() int              { return len(q) }
func (q LargeHeap) Swap(i, j int)         { q[i], q[j] = q[j], q[i] }
func (q *LargeHeap) Push(x interface{})   { *q = append(*q, x.(State)) }
func (q *LargeHeap) Pop() (x interface{}) { *q, x = (*q)[:len(*q)-1], (*q)[len(*q)-1]; return }
func (q *LargeHeap) push(v State)         { heap.Push(q, v) }
func (q *LargeHeap) pop() State           { return heap.Pop(q).(State) }

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	h1 := SmallHeap{}
	h2 := LargeHeap{}
	sum := 0
	Q := getI()
	for q := 0; q < Q; q++ {
		t := getI()
		if t == 2 {
			out(h1[0], sum)
		} else {
			a, b := getI(), getI()
			sum += b
			h1.push(a)
			h2.push(a)
			if h1[0] > h2[0] {
				x := h1.pop()
				y := h2.pop()
				h1.push(y)
				h2.push(x)
				// out(x, y, h1, h2, x-y)
				sum += x - y
			}
		}
	}
}
