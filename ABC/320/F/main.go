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

func calc(a, b int) bool {
	out(p, f, a, b)
	rest := H
	cur := 0
	// 0 to N
	for i := 0; i < N; i++ {
		out(rest, cur, x[i])
		if rest < abs(cur-x[i]) {
			return false
		}
		rest -= abs(cur - x[i])
		out("-->", rest, cur, x[i])
		if i == a {
			rest = min(H, rest+f[i])
		}
		cur = x[i]
		out("-->", rest, cur, x[i])
	}
	// N to 0
	out("cur", cur, "rest", rest)
	for i := N - 1; i >= 0; i-- {
		out(rest, cur, x[i])
		if rest < abs(cur-x[i]) {
			return false
		}
		rest -= abs(cur - x[i])
		out("-->", rest, cur, x[i])
		if i == b {
			rest = min(H, rest+f[i])
		}
		out("-->", rest, cur, x[i])
		cur = x[i]
	}
	return true
}

var N, H int
var x []int
var p, f []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, H = getI(), getI()
	x = append([]int{0}, getInts(N)...)
	p = make([]int, N)
	f = make([]int, N)
	for i := 0; i < N-1; i++ {
		p[i], f[i] = getI(), getI()
	}

	const inf = int(1e18)
	dp := make([][]int, H+1)
	for i := 0; i < H+1; i++ {
		dp[i] = make([]int, H+1)
		for j := 0; j < H+1; j++ {
			dp[i][j] = inf
		}
	}
	for i := 0; i < H+1; i++ {
		dp[H][i] = 0
	}

	for i := 0; i < N; i++ {
		pre := make([][]int, H+1)
		for ii := 0; ii < H+1; ii++ {
			pre[ii] = make([]int, H+1)
			for j := 0; j < H+1; j++ {
				pre[ii][j] = inf
			}
		}
		dp, pre = pre, dp
		dx := x[i+1] - x[i]
		for j := 0; j <= H; j++ {
			for k := 0; k <= H; k++ {
				nj := j - dx
				nk := k + dx
				if nj < 0 {
					continue
				}
				if nk > H {
					continue
				}
				chmin(&dp[nj][nk], pre[j][k])
				chmin(&dp[min(nj+f[i], H)][nk], pre[j][k]+p[i])
				if nk-f[i] >= 0 {
					chmin(&dp[nj][nk-f[i]], pre[j][k]+p[i])
				}
				if nk == H {
					for l := 0; l < f[i]; l++ {
						chmin(&dp[nj][nk-l], pre[j][k]+p[i])
					}
				}
			}
		}
	}

	ans := inf
	for i := 0; i < H+1; i++ {
		ans = min(ans, dp[i][i])
	}
	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}
}
