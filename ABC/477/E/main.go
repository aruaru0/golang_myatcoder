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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

type pqi struct{ a, to int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].a < pq[j].a }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

type edge struct {
	to, cost int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()
	a := getInts(N)
	b := getInts(N)

	node := make([][]edge, N+1)
	for i := 0; i < N; i++ {
		node[i] = append(node[i], edge{(i + 1) % N, a[i]})
		node[(i+1)%N] = append(node[(i+1)%N], edge{i, a[i]})
		node[i] = append(node[i], edge{N, b[i]})
		node[N] = append(node[N], edge{i, b[i]})
	}

	const inf = int(1e18)
	dist := make([]int, N+1)
	for i := 0; i < N+1; i++ {
		dist[i] = inf
	}
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, N})
	dist[N] = 0
	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if cur.a > dist[cur.to] {
			continue
		}
		for _, e := range node[cur.to] {
			if dist[e.to] > dist[cur.to]+e.cost {
				dist[e.to] = dist[cur.to] + e.cost
				heap.Push(&pq, pqi{dist[e.to], e.to})
			}
		}
	}

	sum := make([]int, N*2+1)
	for i := 0; i < 2*N; i++ {
		sum[i+1] = a[i%N]
	}
	for i := 1; i < 2*N; i++ {
		sum[i+1] += sum[i]
	}

	for qi := 0; qi < Q; qi++ {
		s, t := getI()-1, getI()-1
		if s > t {
			s, t = t, s
		}
		if t == N {
			out(dist[s])
		} else {
			l, r := sum[t]-sum[s], sum[s+N]-sum[t]
			z := dist[t] + dist[s]
			out(nmin(l, r, z))
		}
	}

}
