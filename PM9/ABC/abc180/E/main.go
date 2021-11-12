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

const inf = int(1e18)

func calc(from, to int) int {
	return abs(x[to]-x[from]) + abs(y[to]-y[from]) + max(0, z[to]-z[from])
}

var x, y, z []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x, y, z = make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i], z[i] = getI(), getI(), getI()
	}

	dist := make([][]int, N)
	for from := 0; from < N; from++ {
		dist[from] = make([]int, N)
		for to := 0; to < N; to++ {
			dist[from][to] = calc(from, to)
		}
	}

	// ワーシャルフロイド法で最短経路探索
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	n := 1 << N

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}

	// from := 0
	// for to := 0; to < N; to++ {
	// 	dp[from][1|1<<to] = dist[from][to]
	// }
	dp[0][1] = 0
	// out(dist)

	for bit := 0; bit < n; bit++ {
		for from := 0; from < N; from++ {
			if (bit>>from)%2 == 0 {
				continue
			}
			for to := 0; to < N; to++ {
				if (bit>>to)%2 == 1 {
					continue
				}
				if dp[from][bit] == inf {
					continue
				}
				// out(from, to, bit, dp[from][bit], dist[from][to])
				chmin(&dp[to][bit|(1<<to)], dp[from][bit]+dist[from][to])
			}
		}
		// out(dp)
	}
	ans := inf
	for i := 1; i < N; i++ {
		from, to := i, 0
		d := dist[from][to]
		ans = min(ans, dp[from][(1<<N)-1]+d)
	}
	out(ans)
}
