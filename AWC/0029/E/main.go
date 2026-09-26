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

type Edge struct {
	to, cost int
}

const inf = int(1e18)

func calcDist() ([]int, [][]int) {
	N, M := getI(), getI()
	node := make([][]Edge, N)
	for i := 0; i < M; i++ {
		u, v, w := getI()-1, getI()-1, getI()
		node[u] = append(node[u], Edge{v, w})
	}

	s, k := getI()-1, getI()
	t := make([]int, k+1)
	t[0] = s
	for i := 1; i <= k; i++ {
		t[i] = getI() - 1
	}

	d := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		d[i] = make([]int, k+1)
	}

	calcOneDist := func(cur int) []int {
		dist := make([]int, N)
		for i := 0; i < N; i++ {
			dist[i] = inf
		}
		pq := priorityQueue{}
		heap.Push(&pq, pqi{0, cur})
		dist[cur] = 0
		for len(pq) != 0 {
			cur := pq[0]
			heap.Pop(&pq)
			if cur.a > dist[cur.to] {
				continue
			}
			for _, e := range node[cur.to] {
				if dist[e.to] > cur.a+e.cost {
					dist[e.to] = cur.a + e.cost
					heap.Push(&pq, pqi{dist[e.to], e.to})
				}
			}
		}
		return dist
	}

	for i := 0; i <= k; i++ {
		dist := calcOneDist(t[i])
		for j := 0; j <= k; j++ {
			d[i][j] = dist[t[j]]
		}
	}

	return t, d
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	t, d := calcDist()

	n := len(t)
	dist := make([][]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = inf
		}
	}
	for i := 0; i < n; i++ {
		dist[1<<i][i] = d[0][i]
	}
	// out(dist)
	for bit := 0; bit < 1<<n; bit++ {
		for from := 0; from < n; from++ {
			if (bit>>from)%2 == 0 {
				continue
			}
			for to := 0; to < n; to++ {
				if (bit>>to)%2 == 1 {
					continue
				}
				chmin(&dist[bit+(1<<to)][to], dist[bit][from]+d[from][to])
			}
		}
	}

	// for i := 0; i < 1<<n; i++ {
	// 	out(dist[i])
	// }

	ans := dist[(1<<n)-1][0]

	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
