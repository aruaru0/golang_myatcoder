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

type edge struct {
	to, cost int
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

// 解き方は理解。ただ、実装は解説より
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	node := make([][]edge, N)
	for i := 0; i < M; i++ {
		u, v, w := getI()-1, getI()-1, getI()
		node[u] = append(node[u], edge{v, w})
		node[v] = append(node[v], edge{u, w})
	}

	K := getI()
	a := getInts(K)

	ans := make([]int, N)
	for i := 0; i < N; i++ {
		ans[i] = -1
	}

	pq := priorityQueue{}
	pq2 := priorityQueue{}

	for _, e := range a {
		e--
		ans[e] = 0
		for _, v := range node[e] {
			if ans[v.to] < 0 {
				heap.Push(&pq, pqi{v.cost, v.to})
			}
		}
	}

	D := getI()
	for i := 0; i < D; i++ {
		x := getI()
		// キュー１から、探索頂点候補をキュー２に入れる
		for len(pq) != 0 {
			p := pq[0]
			if p.a > x { // 今回の候補はここまで
				break
			}
			heap.Pop(&pq)
			if ans[p.to] >= 0 {
				continue
			}
			heap.Push(&pq2, p)
		}

		// ダイクストラの変形
		for len(pq2) != 0 {
			p := pq2[0]
			heap.Pop(&pq2)
			if ans[p.to] >= 0 {
				continue
			}
			y := p.to
			ans[y] = i + 1
			for _, v := range node[y] {
				if ans[v.to] < 0 {
					if p.a+v.cost <= x {
						heap.Push(&pq2, pqi{p.a + v.cost, v.to})
					} else {
						// 未到達なノードで、距離条件が合わないものはキュー１に追加
						heap.Push(&pq, pqi{v.cost, v.to})
					}
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		out(ans[i])
	}
}
