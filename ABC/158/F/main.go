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

type pair struct {
	first, second int
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	if p[i].first == p[j].first {
		return p[i].second > p[j].second
	}
	return p[i].first > p[j].first
}

// Item :
type Item struct {
	priority, index int
}

// PQ :
type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push :
func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop :
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// End Priority Queue

const inf = math.MaxInt64 / 16
const mod = 998244353

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	tails := make(pairs, N+1)
	index := make([]int, N+1)
	for i := 0; i < N; i++ {
		x, d := getInt(), getInt()
		tails[i] = pair{x + d, x}
		index[i] = x
	}
	tails[N] = pair{inf, inf + 1}
	index[N] = inf
	N++
	sort.Sort(tails)
	sort.Ints(index)
	out(tails)
	out(index)

	pq := make(PQ, 0)
	heap.Init(&pq)
	dict := make(map[int]int)

	r := 0
	l := make([]int, N)
	for i := 0; i < N; i++ {
		dict[index[i]] = i
		l[i] = -1
	}
	out(dict)

	for i := N - 1; i >= 0; i-- {
		out("+++++")
		for r < N && tails[r].first > index[i] {
			heap.Push(&pq, &Item{tails[r].second, 0})
			out("push ", index[i], tails[r])
			r++
		}
		out("-----")
		for pq.Len() > 0 {
			item := heap.Pop(&pq).(*Item)
			out("compare ", item.priority, "<", index[i])
			if item.priority < index[i] {
				l[i] = dict[item.priority]
				heap.Push(&pq, item)
				break
			}
		}
	}

	out(l)

	dp := make([]int, N)
	dpSum := make([]int, N)
	dp[0] = 1
	dpSum[0] = 1
	for i := 1; i < N; i++ {
		dp[i] = dpSum[i-1]
		if l[i]-1 >= 0 {
			dp[i] -= dpSum[l[i]-1]
		}
		if l[i] == -1 {
			dp[i]++
		}
		dp[i] %= mod
		dpSum[i] = (dpSum[i-1] + dp[i]) % mod
	}

	if dp[N-1] < 0 {
		dp[N-1] += mod
	}

	out(dp[N-1])
}
