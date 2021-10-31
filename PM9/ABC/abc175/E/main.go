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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	R, C, K := getI(), getI(), getI()
	s := make([][]int, R)
	dp := make([][][4]int, R)
	for i := 0; i < R; i++ {
		s[i] = make([]int, C)
		dp[i] = make([][4]int, C)
	}
	value := make([]int, K+1)
	for i := 0; i < K; i++ {
		r, c, v := getI()-1, getI()-1, getI()
		value[i+1] = v
		s[r][c] = i + 1
	}
	// out(s)
	if s[0][0] != 0 {
		dp[0][0][1] = value[s[0][0]]
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			for k := 0; k < 4; k++ {
				val := 0
				if c != 0 {
					val = max(val, dp[r][c-1][k])
					if s[r][c] != 0 && k != 0 { //　アイテムが存在
						val = max(val, dp[r][c-1][k-1]+value[s[r][c]])
					}
				}
				dp[r][c][k] = max(dp[r][c][k], val)
			}
			if r != 0 {
				val := nmax(dp[r-1][c][0], dp[r-1][c][1], dp[r-1][c][2], dp[r-1][c][3])
				dp[r][c][0] = max(dp[r][c][0], val)
				if s[r][c] != 0 {
					dp[r][c][1] = max(dp[r][c][1], val+value[s[r][c]])
				}
			}
		}
		// out(dp[r])
	}

	// out(dp[R-1][C-1])
	ans := 0
	for i := 0; i < 4; i++ {
		ans = max(ans, dp[R-1][C-1][i])
	}
	out(ans)
}
