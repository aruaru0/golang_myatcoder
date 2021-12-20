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

const n = 1 << 20

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	A, B := getI(), getI()
	N := B - A + 1
	// 72までの素数を列挙
	primes := [20]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71}

	// 上記の素数の20個を使ってbitDP
	dp := make([][n]int, N+1)
	dp[0][0] = 1
	for x := A; x <= B; x++ {
		i := x - A
		pat := 0
		// xがどの素数て割り切れるかのパターンを作成
		for j := 0; j < len(primes); j++ {
			if x%primes[j] == 0 {
				pat |= 1 << j
			}
		}
		// bitを0～1<<20まで変えてテーブルを更新
		for bit := 0; bit < n; bit++ {
			// まず、コピー
			dp[i+1][bit] += dp[i][bit]
			// patとbitの&が0なら（重複する素数倍ではないなら）、DPテーブルを更新
			if pat&bit == 0 {
				dp[i+1][bit|pat] += dp[i][bit]
			}
		}
	}

	ans := 0
	// すべてのパターンを足しこむ
	for bit := 0; bit < n; bit++ {
		ans += dp[N][bit]
	}
	out(ans)
}
