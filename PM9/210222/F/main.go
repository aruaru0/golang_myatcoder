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

const inf = int(1e18)

func bsf(cur int) []int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = inf
	}
	q := []int{cur}
	dist[cur] = 0
	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		for _, e := range node[c] {
			if dist[e] > dist[c]+1 {
				dist[e] = dist[c] + 1
				q = append(q, e)
			}
		}
	}
	return dist
}

var N, M int
var node [][]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node = make([][]int, N)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		node[a] = append(node[a], b)
		node[b] = append(node[b], a)
	}
	K := getI()
	c := make([]int, K)
	for i := 0; i < K; i++ {
		c[i] = getI() - 1
	}
	d := make([][]int, K)
	for i := 0; i < K; i++ {
		dist := bsf(c[i])
		d[i] = make([]int, 0)
		for _, e := range c {
			d[i] = append(d[i], dist[e])
		}
	}

	n := 1 << K
	dp := make([][]int, K)
	for i := 0; i < K; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}
	for i := 0; i < K; i++ {
		dp[i][1<<i] = 0
	}
	for k := 0; k < n; k++ {
		for from := 0; from < K; from++ {
			if (k>>from)%2 == 0 {
				continue
			}
			for to := 0; to < K; to++ {
				if (k>>to)%2 == 1 {
					continue
				}
				dp[to][k|(1<<to)] = min(dp[to][k|(1<<to)], dp[from][k]+d[from][to])
			}
		}
	}
	ans := inf
	for i := 0; i < K; i++ {
		ans = min(ans, dp[i][n-1])
	}
	if ans == inf {
		out(-1)
		return
	}
	out(ans + 1)
}
