package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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

var N, M int
var X []float64
var Y []float64

func f(x0, y0, x1, y1 float64, bit int) float64 {
	bit = bit >> N
	cnt := bits.OnesCount(uint(bit))
	div := 1 << cnt

	dx := x0 - x1
	dy := y0 - y1
	ret := math.Sqrt(dx*dx+dy*dy) / float64(div)
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	n := M + N
	X = make([]float64, n)
	Y = make([]float64, n)
	for i := 0; i < n; i++ {
		X[i], Y[i] = getF(), getF()
	}
	dp := make([][]float64, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 1e18
		}
	}
	for i := 0; i < n; i++ {
		dp[1<<i][i] = f(0, 0, X[i], Y[i], 0)
	}

	mask := 1<<N - 1
	for bit := 0; bit < 1<<n; bit++ {
		for from := 0; from < n; from++ {
			if (bit>>from)%2 == 0 {
				continue
			}
			for to := 0; to < n; to++ {
				if (bit>>to)%2 == 1 {
					continue
				}
				pos := bit | (1 << to)
				dp[pos][to] = math.Min(dp[pos][to], dp[bit][from]+f(X[from], Y[from], X[to], Y[to], bit))
				// fmt.Fprintf(wr, "%b <- %b: from %d ->  %d %f -> %f\n", pos, bit, from, to, f(X[from], Y[from], X[to], Y[to], bit), dp[pos][to])

			}
		}
	}

	ans := 1e18
	for bit := 0; bit < 1<<n; bit++ {
		if bit&mask == mask {
			for i := 0; i < n; i++ {
				ans = math.Min(ans, dp[bit][i]+f(X[i], Y[i], 0, 0, bit))
			}
		}
	}
	out(ans)
}
