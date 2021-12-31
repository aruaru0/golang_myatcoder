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

	n := 1 << N
	popcount := make([]int, n)
	for i := 0; i < n; i++ {
		popcount[i] = bits.OnesCount(uint(i))
	}

	dp := make([]int, n)
	dp[0] = 1

	for bit := 0; bit < n; bit++ {
		for next := 0; next < N; next++ {
			// 既にnextｎが立っている場合は、スキップ
			if (bit>>next)%2 == 1 {
				continue
			}
			// 遷移先をnext_bitに設定
			next_bit := bit | (1 << next)
			bitcnt := popcount[next_bit]
			ok := true
			// bitcnt未満のものについてのみ条件を満たすかチェック
			for _, e := range c[bitcnt-1] {
				mask := 1<<(e.y+1) - 1
				// 次の遷移状態のe.y以下の部分のbit数がe.zより大きい場合は条件を満たさない
				if popcount[next_bit&mask] > e.z {
					ok = false
				}
			}
			if ok {
				dp[next_bit] += dp[bit]
			}
		}
	}
	out(dp[n-1])
}
