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

func f(a []int) float64 {
	tot := 0.0
	for i := 1; i < len(a); i++ {
		dy := a[i] - a[i-1]
		dx := 1
		d := dx*dx + dy*dy
		tot += math.Sqrt(float64(d))
	}
	return tot
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)

	tot := 0
	for i := 0; i < N; i++ {
		tot += a[i]
	}

	var dist [110][110]float64
	for i := 0; i <= tot; i++ {
		for j := 0; j <= tot; j++ {
			dist[i][j] = math.Sqrt(float64((i-j)*(i-j) + 1))
		}
	}

	const inf = 1e10

	dp := make([][110][110]float64, N+1)
	for i := 0; i <= N; i++ {
		for rest := 0; rest <= tot; rest++ {
			for last := 0; last <= tot; last++ {
				dp[i][rest][last] = inf
			}
		}
	}
	dp[1][tot][0] = 0

	for i := 1; i < N-1; i++ {
		for rest := 0; rest <= tot; rest++ {
			for last := 0; last <= tot; last++ {
				if dp[i][rest][last] < inf {
					for j := 0; j <= rest; j++ {
						dp[i+1][rest-j][j] = math.Min(dp[i+1][rest-j][j], dp[i][rest][last]+dist[last][j])
					}
				}
			}
		}
	}

	ans := inf
	for i := 0; i <= tot; i++ {
		ans = math.Min(ans, dp[N-1][0][i]+dist[i][0])
	}
	out(ans)
}
