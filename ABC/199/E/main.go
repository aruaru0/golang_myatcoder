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
	y, z int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	c := make([][]pair, N)
	for i := 0; i < M; i++ {
		x, y, z := getI()-1, getI()-1, getI()
		c[x] = append(c[x], pair{y, z})
	}

	// bitDPで考える
	n := 1 << N
	dp := make([]int, n)

	// 何度も利用するので１の立っているビットの数をあらかじめ計算しておく
	popcount := make([]int, n)
	for i := 1; i < n; i++ {
		popcount[i] = popcount[i/2] + i%2
	}

	dp[0] = 1
	for bit := 0; bit < n; bit++ { // bitパターン網羅
		for next := 0; next < N; next++ { // 次に集合に入れるやつ
			// bitがすでに立っている場合は飛ばす
			if (bit>>next)%2 == 1 {
				continue
			}
			next_bit := bit | (1 << next) // 次のビットパターンを計算
			bitcnt := popcount[next_bit]  // 次のビットパターンの１の個数
			ok := true
			for _, e := range c[bitcnt-1] { // 条件すべてに対して
				mask := 1<<(e.y+1) - 1             // y以下の部分だけマスクを生成
				if popcount[next_bit&mask] > e.z { // y以下の部分がz個より大きければ
					ok = false // 条件を満たさない
				}
			}
			if ok { // 条件を満たした場合
				dp[next_bit] += dp[bit] // 条件を満たしたので追加する
			}
		}
	}
	out(dp[n-1])
}
