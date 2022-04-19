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
	h, w, k := getI(), getI(), getI()
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = getInts(w)
	}
	ans := inf
	for pi := 0; pi < h; pi++ {
		for pj := 0; pj < w; pj++ {
			x := a[pi][pj]
			dp := make([][]int, h)
			for i := 0; i < h; i++ {
				dp[i] = make([]int, w)
				for j := 0; j < w; j++ {
					dp[i][j] = inf
				}
			}
			dp[0][0] = max(a[0][0]-x, 0)
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if i+1 < h {
						chmin(&dp[i+1][j], dp[i][j]+max(a[i+1][j]-x, 0))
					}
					if j+1 < w {
						chmin(&dp[i][j+1], dp[i][j]+max(a[i][j+1]-x, 0))
					}
				}
			}
			chmin(&ans, dp[h-1][w-1]+x*k)
		}
	}
	out(ans)
}
