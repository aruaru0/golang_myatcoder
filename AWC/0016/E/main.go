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

type Edge struct {
	to, cost int
}

type pqi struct {
	cost, to, mask int
}

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].cost > pq[j].cost }
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
	p := getInts(N)
	s, t := getI()-1, getI()-1

	node := make([][]Edge, N)
	for i := 0; i < M; i++ {
		u, v, w := getI()-1, getI()-1, getI()
		node[u] = append(node[u], Edge{v, w})
		node[v] = append(node[v], Edge{u, w})
	}

	const inf = int(1e18)
	size := 1 << N
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, size)
		for j := 0; j < size; j++ {
			dp[i][j] = -inf
		}
	}

	sm := 1 << s
	dp[s][sm] = p[s]
	pq := priorityQueue{}
	heap.Push(&pq, pqi{p[s], s, sm})

	for len(pq) != 0 {
		cur := pq[0]
		heap.Pop(&pq)
		if cur.cost < dp[cur.to][cur.mask] {
			continue
		}
		for _, e := range node[cur.to] {
			next := cur.mask | (1 << e.to)
			get := 0
			// 訪問済みでない場合のみ地点の値を加算する
			if next != cur.mask {
				get = p[e.to]
			}
			// コスト計算は、現在から移動コストを引いて、報酬となる
			nextCost := cur.cost - e.cost + get
			if nextCost > dp[e.to][next] {
				dp[e.to][next] = nextCost
				heap.Push(&pq, pqi{nextCost, e.to, next})
			}
		}
	}

	ans := -inf
	for j := 0; j < size; j++ {
		ans = max(ans, dp[t][j])
	}
	out(ans)
}
