package main

import (
	"bufio"
	"fmt"
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

var D, G int
var p, c []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	D, G = getI(), getI()/100
	p = make([]int, D)
	c = make([]int, D)
	for i := 0; i < D; i++ {
		p[i], c[i] = getI(), getI()/100
	}

	const MAX = 200000
	const inf = int(1e10)
	dp := make([][MAX * 2]int, D+1)
	for i := 0; i <= D; i++ {
		for j := 1; j < MAX; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0
	for i := 0; i < D; i++ {
		for j := 0; j < MAX; j++ {
			for k := 1; k < p[i]; k++ {
				pos := k * (i + 1)
				dp[i+1][j+pos] = nmin(dp[i+1][j+pos], dp[i][j+pos], dp[i][j]+k)
			}
			pos := p[i]*(i+1) + c[i]
			dp[i+1][j+pos] = nmin(dp[i+1][j+pos], dp[i][j+pos], dp[i][j]+p[i])
		}
		// out(dp[i+1][:20])
	}

	ans := inf
	for i := G; i < MAX; i++ {
		ans = min(ans, dp[D][i])
	}
	out(ans)
}
