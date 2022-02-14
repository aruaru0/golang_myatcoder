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

const mod = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, K := getI(), getI(), getI()
	node := make([][]int, N)
	for i := 0; i < M; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}
	dp := make([][]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, N)
	}
	dp[0][0] = 1
	for day := 0; day < K; day++ {
		tot := 0
		// その日の合計を計算
		for lst := 0; lst < N; lst++ {
			tot += dp[day][lst]
			tot %= mod
		}
		// 次の位置の計算
		for nxt := 0; nxt < N; nxt++ {
			// 次の日は、全体合計から、そのノードスタートを引く
			dp[day+1][nxt] = tot - dp[day][nxt]
			dp[day+1][nxt] %= mod
			if dp[day+1][nxt] < 0 {
				dp[day+1][nxt] += mod
			}
			// つながっていないルートの数を引く
			for _, lst := range node[nxt] {
				dp[day+1][nxt] -= dp[day][lst]
				dp[day+1][nxt] %= mod
				if dp[day+1][nxt] < 0 {
					dp[day+1][nxt] += mod
				}
			}
		}
	}

	out(dp[K][0])
}
