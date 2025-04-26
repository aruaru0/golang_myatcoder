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
	node := make([][]int, N)
	for i := 0; i < M; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	s := getI() - 1
	K := getI()
	t := make([]int, K+1)
	t[0] = s
	for i := 0; i < K; i++ {
		t[i+1] = getI() - 1
	}

	K++

	tbl := make([][]int, K)
	for i := 0; i < K; i++ {
		tbl[i] = make([]int, K)
	}

	const inf = int(1e18)
	bsf := func(cur int) []int {
		q := []int{cur}
		dist := make([]int, N)
		for i := 0; i < N; i++ {
			dist[i] = inf
		}
		dist[cur] = 0
		for len(q) != 0 {
			cur := q[0]
			q = q[1:]
			for _, e := range node[cur] {
				if dist[e] == inf {
					dist[e] = dist[cur] + 1
					q = append(q, e)
				}
			}
		}
		return dist
	}

	for i, e := range t {
		dist := bsf(e)
		for j, v := range t {
			tbl[i][j] = dist[v]
		}
	}

	n := 1 << K
	// dp[S][i] : 状態Sでi番目に到着している時のコストの最小値
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, K)
		for j := 0; j < K; j++ {
			dp[i][j] = inf
		}
	}

	dp[1][0] = 0
	// pq := priorityQueue{}
	// heap.Push(&pq, pqi{0, 0})

	for bit := 0; bit < n; bit++ {
		for from := 0; from < K; from++ {
			if (bit>>from)%2 == 0 {
				continue
			}
			for to := 0; to < K; to++ {
				if (bit>>to)%2 == 1 {
					continue
				}
				nxt := bit | (1 << to)
				chmin(&dp[nxt][to], dp[bit][from]+tbl[from][to])
				// if dp[nxt][to] > dp[bit][from]+tbl[from][to] {
				// 	dp[nxt][to] = dp[bit][from] + tbl[from][to]
				// }
			}
		}
	}

	// for i := 0; i < K; i++ {
	// 	out(tbl[i])
	// }
	ans := inf
	for i := 0; i < K; i++ {
		ans = min(ans, dp[n-1][i])
	}
	out(ans)
}
