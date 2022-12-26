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

type pqi struct{ cost, to int }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

type edge struct {
	to, cost int
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x := make([]int, N)
	y := make([]int, N)
	z := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i], z[i] = getI(), getI(), getI()
	}
	node := make([][]edge, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			cost := abs(x[j]-x[i]) + abs(y[j]-y[i]) + max(0, z[j]-z[i])
			node[i] = append(node[i], edge{j, cost})
		}
	}

	n := 1 << N
	dp := make([][1 << 17]int, N)
	for bit := 0; bit < n; bit++ {
		for i := 0; i < N; i++ {
			dp[i][bit] = inf
		}
	}
	dp[0][1] = 0
	for bit := 0; bit < n; bit++ {
		for from := 0; from < N; from++ {
			if (bit>>from)%2 == 0 {
				continue
			}
			for to := 0; to < N; to++ {
				if (bit>>to)%2 == 1 {
					continue
				}
				// out(from, to, bit, dp[to][bit|(1<<to)], dp[from][bit], node[from][to].cost)
				dp[to][bit|(1<<to)] = min(dp[to][bit|(1<<to)], dp[from][bit]+node[from][to].cost)
			}
		}
	}

	ans := inf
	for i := 1; i < N; i++ {
		ans = min(ans, dp[i][n-1]+node[i][0].cost)
	}
	out(ans)

}
