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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getI() - 1
	}
	b := make([][]int, M)
	for i := 0; i < M; i++ {
		b[i] = make([]int, N+1)
	}
	for i := 0; i < N; i++ {
		b[a[i]][i+1]++
	}
	tot := make([]int, M)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			b[i][j+1] += b[i][j]
		}
		tot[i] = b[i][N]
	}
	n := 1 << M
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for bit := 0; bit < n; bit++ {
		for i := 0; i < M; i++ {
			if (bit>>i)%2 == 1 {
				continue
			}
			// 総数をカウント
			sum := 0
			for j := 0; j < M; j++ {
				if (bit>>j)%2 == 1 {
					sum += tot[j]
				}
			}
			l, r := sum, sum+tot[i]
			// fmt.Fprintf(wr, "%b\n", bit)
			change := tot[i] - (b[i][r] - b[i][l])
			// out(tot[i], l, r, b[i][r], b[i][l], change)
			dp[bit|(1<<i)] = min(dp[bit|(1<<i)], dp[bit]+change)
		}
	}

	out(dp[n-1])
}
