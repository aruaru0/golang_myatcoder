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

type Edge struct {
	to, cost int
}

const inf = int(1e15)

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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	node := make([][]Edge, N)
	for i := 0; i < M; i++ {
		u, v, w := getI()-1, getI()-1, getI()
		node[u] = append(node[u], Edge{v, w})
	}

	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = inf
		}
	}

	for i := 0; i < N; i++ {
		pq := priorityQueue{}
		heap.Push(&pq, pqi{0, i})
		dist[i][i] = 0
		for len(pq) != 0 {
			cur := pq[0]
			heap.Pop(&pq)
			if dist[i][cur.to] < cur.a {
				continue
			}
			for _, e := range node[cur.to] {
				if dist[i][e.to] > dist[i][cur.to]+e.cost {
					dist[i][e.to] = dist[i][cur.to] + e.cost
					heap.Push(&pq, pqi{dist[i][e.to], e.to})
				}
			}
		}
		// out(dist[i])
	}

	d := make([][]int, 1<<N)
	for i := 0; i < 1<<N; i++ {
		d[i] = make([]int, N)
		for j := 0; j < N; j++ {
			d[i][j] = inf
		}
	}
	for i := 0; i < N; i++ {
		d[1<<i][i] = 0
	}

	for bit := 0; bit < 1<<N; bit++ {
		for from := 0; from < N; from++ {
			if (bit>>from)%2 == 0 { // fromに訪れていなければNG
				continue
			}
			for to := 0; to < N; to++ {
				if (bit>>to)%2 == 1 { // toに訪れていればスキップ
					continue
				}
				if dist[from][to] == inf { // fromとtoが繋がっていない場合は処理しない
					continue
				}
				if d[bit][from] == inf {
					continue
				}
				chmin(&d[bit+(1<<to)][to], d[bit][from]+dist[from][to])
			}
		}
	}

	ans := inf
	for i := 0; i < N; i++ {
		ans = min(ans, d[1<<N-1][i])
	}
	// out(d)
	if ans == inf {
		out("No")
	} else {
		out(ans)
	}
}
