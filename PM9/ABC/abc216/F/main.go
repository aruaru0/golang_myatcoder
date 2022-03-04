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

type pair struct {
	a, b int
}

const mod = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	b := getInts(N)
	p := make([]pair, N)
	for i := 0; i < N; i++ {
		p[i] = pair{a[i], b[i]}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].a < p[j].a
	})

	// dp[i][j] i番目まで決めていてi番目を選択しているときに、Ｂの要素の総和がjである組み合わせ数
	dp := make([][]int, 5100)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, 5100)
	}
	tot := make([]int, 5100)
	dp[0][0] = 1
	tot[0] = 1
	for to := 1; to <= N; to++ {
		for b := 0; b < 5100; b++ {
			if p[to-1].b <= b {
				dp[to][b] = tot[b-p[to-1].b]
			}
		}
		for b := 0; b < 5100; b++ {
			tot[b] += dp[to][b]
			tot[b] %= mod
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		for b := 0; b <= p[i-1].a; b++ {
			ans += dp[i][b]
			ans %= mod
		}
	}
	out(ans)
}
