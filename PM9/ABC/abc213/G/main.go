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

const mod = int(998244353)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		a, b := getI()-1, getI()-1
		g[a][b] = 1
		g[b][a] = 1
	}

	n2 := 1 << n
	// e[s] : sに含まれる辺の数
	e := make([]int, n2)
	for s := 0; s < n2; s++ {
		if s == 0 {
			continue
		}
		// 1が立っている最下位を調べる
		b := s & -s
		// 最下位が立っていないやつをコピー
		e[s] = e[s^b]
		// 1が立っているビット→辺の番号へ変換
		v := bits.OnesCount(uint(b - 1))
		for i := 0; i < n; i++ {
			if (s>>i)&1 == 1 { // 辺iが含まれている場合
				e[s] += g[i][v] //　辺を足す
			}
		}
	}

	// 2のn乗を計算しておく
	two := make([]int, m+1)
	two[0] = 1
	for i := 0; i < m; i++ {
		two[i+1] = two[i] * 2 % mod
	}

	// bitDP本体
	dp := make([]int, n2)
	for s := 0; s < n2; s++ {
		if s&1 == 0 {
			continue
		}
		// 辺の数の組み合わせ数すべて
		dp[s] = two[e[s]]
		// 辺を削除した場合の数を引く
		for t := s; t != 0; {
			t = (t - 1) & s
			if t&1 == 1 {
				// two[e[s^t]] t含まれない点は自由
				dp[s] -= dp[t] * two[e[s^t]]
				dp[s] %= mod
				if dp[s] < 0 {
					dp[s] += mod
				}
			}
		}
	}

	ans := make([]int, n)
	for s := 0; s < n2; s++ {
		if s&1 == 0 {
			continue
		}
		for i := 0; i < n; i++ {
			if (s>>i)&1 == 1 {
				ans[i] += dp[s] * two[e[(n2-1)^s]]
				ans[i] %= mod
			}
		}
	}
	for i := 1; i < n; i++ {
		out(ans[i])
	}
}
