package main

import (
	"bufio"
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

func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

// Queue ... Priority Queue
type Queue[T any] struct {
	data []T
	less func(i, j T) bool
}

// New ... Create new priority queue
func New[T any](less func(i, j T) bool) Queue[T] {
	var ret Queue[T]
	ret.data = make([]T, 0)
	ret.less = less
	return ret
}

func (q Queue[T]) Len() int {
	return len(q.data)
}

func (q Queue[T]) Swap(i, j int) {
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

func (q *Queue[T]) Push(x T) {
	q.data = append(q.data, x)
	cur := q.Len()
	parent := cur / 2
	for cur != 1 {
		if q.less(q.data[cur-1], q.data[parent-1]) {
			q.Swap(cur-1, parent-1)
		} else {
			break
		}
		cur = parent
		parent = cur / 2
	}
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.Len() == 0 {
		var item T
		return item, false
	}
	old := *q
	n := len(old.data)
	item := old.data[0]

	old.data[0] = old.data[n-1]
	old.data = old.data[:n-1]
	cur := 1
	for {
		nxt0 := cur * 2
		nxt1 := cur*2 + 1
		if nxt0 > len(old.data) {
			break
		}
		nxt := nxt0
		if nxt1 <= len(old.data) && old.less(q.data[nxt1-1], q.data[nxt0-1]) {
			nxt = nxt1
		}
		if old.less(q.data[nxt-1], q.data[cur-1]) {
			old.Swap(nxt-1, cur-1)
		} else {
			break
		}

		cur = nxt
	}

	*q = old
	return item, true
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)

	pq := New[int](func(i, j int) bool {
		return i > j
	})

	for i := 0; i < N; i++ {
		pq.Push(a[i])
	}

	cnt := 0
	for {
		x, _ := pq.Pop()
		y, _ := pq.Pop()
		if y <= 0 {
			break
		}
		x--
		y--
		pq.Push(x)
		pq.Push(y)
		cnt++
	}

	out(cnt)

}
