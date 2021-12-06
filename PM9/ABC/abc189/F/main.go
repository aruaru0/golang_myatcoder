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

type L struct {
	a, b float64
}

func plus(x, y L) L {
	return L{x.a + y.a, x.b + y.b}
}

func minus(x, y L) L {
	return L{x.a - y.a, x.b - y.b}
}

func mul(x, y L) L {
	return L{x.a * y.a, x.b * y.b}
}

func div(x L, y float64) L {
	return L{x.a / y, x.b / y}
}

const eps = 1e-6

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, K := getI(), getI(), getI()
	a := make([]bool, N+1)
	// 振り出しに戻るマスをmapで管理
	for i := 0; i < K; i++ {
		a[getI()] = true
	}
	// dp[i] : iからゴールまでの期待値
	// dp[N] = 0
	// dp[i] = Σdp[i+j]/M + 1 :  1 <= j <= M
	// ただし、振り出しに戻るマス
	// dp[i] = dp[0]
	//  --> DAGじゃない（ループが発生している）
	// 解法： この問題の答えdp[0] = xとおく
	// dp[0] = x
	//       = ax + b
	// x = ax + b -> x - ax = b -> x = b/(1-a)

	dp := make([]L, N+1)
	var s L
	for i := N - 1; i >= 0; i-- {
		// dp[i] = sum(dp[i+1]-dp[i+m]) -> s
		if a[i] == true {
			dp[i] = L{1, 0}
		} else {
			dp[i] = plus(div(s, float64(M)), L{0, 1})
		}
		// 尺取り的にエリアの累積和を更新していく
		s = plus(s, dp[i])
		if i+M <= N {
			s = minus(s, dp[i+M])
		}
	}
	aa := dp[0].a
	bb := dp[0].b
	if aa+eps > 1 {
		out(-1)
		return
	}
	ans := bb / (1 - aa)
	out(ans)
}
