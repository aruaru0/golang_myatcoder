package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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
	to, c, t int
}

const inf = 1001001001001

type pqi struct{ time, node, cost int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].time < pq[j].time }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	N, C, V := getInt(), getInt(), getInt()
	s := getInts(V)
	t := getInts(V)
	y := getInts(V)
	m := getInts(V)
	node := make([][]edge, N)
	for i := 0; i < V; i++ {
		from := s[i] - 1
		to := t[i] - 1
		node[from] = append(node[from], edge{to, y[i], m[i]})
		// 有向グラフなのに双方向にしてバグらせた。削除
		//		node[to] = append(node[to], edge{from, y[i], m[i]})
	}

	// テーブル初期化
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			dp[i][j] = inf
		}
	}

	// ダイクストラ法
	//  -優先キューは時間。コストを追加（拡張）
	dp[0][C] = 0
	pq := priorityQueue{}
	heap.Push(&pq, pqi{0, 0, C})
	for len(pq) != 0 {
		x := heap.Pop(&pq).(pqi)
		n := x.node
		m := x.cost
		t := x.time
		// 記録されている時間が既に小さければスキップ
		if dp[n][m] < t {
			continue
		}
		for _, e := range node[n] {
			if m-e.c < 0 {
				continue
			}
			if dp[e.to][m-e.c] > dp[n][m]+e.t {
				dp[e.to][m-e.c] = dp[n][m] + e.t
				heap.Push(&pq, pqi{dp[e.to][m-e.c], e.to, m - e.c})
			}
		}
	}
	//  結果から最小値を検索
	ans := inf
	for i := 0; i <= C; i++ {
		ans = min(ans, dp[N-1][i])
	}
	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
