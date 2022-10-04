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

func calc(x1, y1, x2, y2 int, flg bool) map[int]int {
	// out("calc", x1, y1, x2, y2)
	// (x1, y1)から、(x2, y2)に行くまでのパターンを列挙 flgがtrueの場合、x1, y1を含める
	dp := make([][]map[int]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]map[int]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make(map[int]int)
		}
	}
	if flg {
		dp[y1][x1][a[y1][x1]] = 1
	} else {
		dp[y1][x1][0] = 1
	}
	for h := y1; h <= y2; h++ {
		for w := x1; w <= x2; w++ {
			x := a[h][w]
			if h > y1 {
				for e, i := range dp[h-1][w] {
					dp[h][w][e^x] += i
				}
			}
			if w > x1 {
				for e, i := range dp[h][w-1] {
					dp[h][w][e^x] += i
				}
			}
			// out(h, w, dp[h][w])
		}
	}
	return dp[y2][x2]
}

var N int
var a [][]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	a = make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInts(N)
	}

	ans := 0
	for i := 0; i < N; i++ {
		x, y := i, N-i-1
		p := calc(0, 0, x, y, true)
		q := calc(x, y, N-1, N-1, false) //, x, y)
		// p ^ q = 0 になるのは、p==qの時のみなので、片方はmapを探索するだけでよい
		for e, n := range p {
			ans += q[e] * n
		}
	}
	out(ans)
}
