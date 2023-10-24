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

type pqi struct{ a, to, st int }

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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, A, B, C := getI(), getI(), getI(), getI()
	d := make([][]int, N)
	for i := 0; i < N; i++ {
		d[i] = getInts(N)
	}

	const inf = int(1e18)

	dist := make([][2]int, N)
	for i := 0; i < N; i++ {
		dist[i][0] = inf
		dist[i][1] = inf
	}
	pq := priorityQueue{}
	dist[0][0] = 0
	heap.Push(&pq, pqi{0, 0, 0})

	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)

		from := cur.to
		cost := cur.a
		st := cur.st

		if cost > dist[from][st] {
			continue
		}
		if st == 0 { // ここまで社用車で来ている場合
			for to := 0; to < N; to++ {
				if from == to {
					continue
				}
				// 社用車ー＞社用車
				if dist[to][0] > dist[from][0]+d[from][to]*A {
					dist[to][0] = dist[from][0] + d[from][to]*A
					heap.Push(&pq, pqi{dist[to][0], to, 0})
				}
			}
		}
		// 社用車、電車ー＞電車
		for to := 0; to < N; to++ {
			if from == to {
				continue
			}
			if dist[to][1] > dist[from][st]+d[from][to]*B+C {
				dist[to][1] = dist[from][st] + d[from][to]*B + C
				heap.Push(&pq, pqi{dist[to][1], to, 1})
			}
		}
	}

	out(min(dist[N-1][0], dist[N-1][1]))

}
