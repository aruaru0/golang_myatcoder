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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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
	N, M, L := getI(), getI(), getI()
	a := getInts(N)
	// S0 = a0 + a1 ... a{L-1} = 0 (mod M)
	// S1 = a1 + a2 ... aL = 0 (mod M)
	// S1 - S0 = aL - a0 = 0 (mod M)
	// aL = a0 (mod M)　<- Lステップで(mod M)で同じ値になる必要がある

	const inf = int(1e18)
	// dp[i] : Mで割ったあまりがiになるように操作する最小回数
	dp := make([]int, M)
	for i := 0; i < M; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	// AiをLとびで操作するので、走査回数はL回
	for i := 0; i < L; i++ {
		tmp := make([]int, M)
		for j := 0; j < M; j++ {
			tmp[j] = inf
		}

		//
		// pos i, i+L, i+2L....についてdp計算する
		//
		for j := 0; j < M; j++ {
			// jが0~Mまでループ（jに合わせるコストを計算する）
			cost := 0
			// L飛びで走査
			for k := i; k < N; k += L {
				// jがa[k]より大きい場合は、j-a[k]加算すればOK加算
				cost += (j - a[k] + M) % M
			}
			// costは(ai, ai+L, ai+2L...)をj(mode M)するときのコストの和
			// DPの更新
			for k := 0; k < M; k++ {
				// 状態kにコストを加えることで(k+j)mod Mにする最小コスト
				tmp[(k+j)%M] = min(tmp[(k+j)%M], dp[k]+cost)
			}
		}
		// i = 0の時は、j(mod M)にするために加算するコスト
		// i = 1の時は、i=0, 1の双方を合わせてj(mod M)にするためのコストとなる
		dp = tmp
		// out(dp)
	}

	out(dp[0])
}
