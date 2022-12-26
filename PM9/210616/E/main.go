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

type pos struct {
	r, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	R, C, K := getI(), getI(), getI()
	p := make(map[pos]int)
	for i := 0; i < K; i++ {
		r, c, v := getI()-1, getI()-1, getI()
		p[pos{r, c}] = v
	}
	dp := make([][][4]int, R)
	for i := 0; i < R; i++ {
		dp[i] = make([][4]int, C)
	}
	v, ok := p[pos{0, 0}]
	if ok {
		dp[0][0][1] = v
	}

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			u := 0
			if r > 0 {
				u = nmax(dp[r-1][c][0], dp[r-1][c][1], dp[r-1][c][2], dp[r-1][c][3])
			}
			dp[r][c][0] = u
			if c > 0 {
				for k := 0; k < 4; k++ {
					v, ok := p[pos{r, c}]
					if ok {
						if k == 1 {
							dp[r][c][k] = nmax(dp[r][c-1][k], dp[r][c][k-1]+v, dp[r][c-1][k-1]+v)
						} else if k > 0 {
							dp[r][c][k] = nmax(dp[r][c-1][k], dp[r][c-1][k-1]+v)
						} else {
							dp[r][c][k] = max(dp[r][c-1][k], dp[r][c][k])
						}
					} else {
						dp[r][c][k] = max(dp[r][c][k], dp[r][c-1][k])
					}
				}
			} else {
				v, ok := p[pos{r, c}]
				if ok {
					dp[r][c][1] = dp[r][c][0] + v
				}
			}
		}
	}
	ans := 0
	for i := 0; i < 4; i++ {
		ans = max(ans, dp[R-1][C-1][i])
	}
	out(ans)
}
