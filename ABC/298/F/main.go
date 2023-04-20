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

type line struct {
	r, c, x, i, j int
}

func solve(n int, r, c, x []int) int {
	mr, mc := make(map[int]int), make(map[int]int)
	field := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		mr[r[i]] += x[i]
		mc[c[i]] += x[i]
		if field[r[i]] == nil {
			field[r[i]] = make(map[int]int)
		}
		field[r[i]][c[i]] = x[i]
	}

	var rows, cols []line
	for k, v := range mr {
		rows = append(rows, line{k, 0, v, -1, -1})
	}
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].x > rows[j].x
	})
	for k, v := range mc {
		cols = append(cols, line{0, k, v, -1, -1})
	}
	sort.Slice(cols, func(i, j int) bool {
		return cols[i].x > cols[j].x
	})

	q := &PriorityQueue{}
	heap.Init(q)
	for i := range rows {
		heap.Push(q, line{rows[i].r, cols[0].c, rows[i].x + cols[0].x, i, 0})
	}

	var ans int
	for q.Len() > 0 {
		node := heap.Pop(q).(line)
		if node.x <= ans {
			break
		}
		ans = max(ans, node.x-field[node.r][node.c])
		if node.j+1 < len(cols) {
			ni, nj := node.i, node.j+1
			heap.Push(q, line{node.r, cols[nj].c, rows[ni].x + cols[nj].x, ni, nj})
		}
	}

	return ans
}

type PriorityQueue []line

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].x > pq[j].x
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(line))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	var r, c, x []int
	for i := 0; i < N; i++ {
		r = append(r, getI())
		c = append(c, getI())
		x = append(x, getI())
	}
	ans := solve(N, r, c, x)
	out(ans)
}
