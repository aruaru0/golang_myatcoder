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

type pqi struct{ a, b int }

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

	// 0を最小化する問題として考えると
	// 蟻本p.104の問題に帰着できる
	n, m := getI(), getI()
	node := make([][]edge, n+1)
	for i := 0; i < n; i++ { // 次へは1のコストで到達可能
		node[i] = append(node[i], edge{i + 1, 1})
	}
	for i := 0; i < n; i++ { // 手前には0のコストで到達可能
		node[i+1] = append(node[i+1], edge{i, 0})
	}
	for i := 0; i < m; i++ { // つながった部分は、0の数（r-l)-xのコストで到達可能
		l, r, x := getI()-1, getI(), getI()
		node[l] = append(node[l], edge{r, (r - l) - x})
	}

	//　設定が追わればあとはダイクストラ法で最短経路を求めれば終わり
	const inf = int(1e18)
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = inf
	}
	dist[0] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, 0})
	for len(pq) != 0 {
		d, v := pq[0].a, pq[0].b
		heap.Pop(&pq)
		if dist[v] != d {
			continue
		}
		for _, e := range node[v] {
			nd := d + e.cost
			if dist[e.to] <= nd {
				continue
			}
			dist[e.to] = nd
			heap.Push(&pq, pqi{nd, e.to})
		}
	}
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		// コスト0で接続されている部分は1である
		ans[i] = (dist[i+1] - dist[i]) == 0
	}
	for i := 0; i < n; i++ {
		if ans[i] {
			fmt.Fprint(wr, "1 ")
		} else {
			fmt.Fprint(wr, "0 ")
		}
	}
	out()
}
