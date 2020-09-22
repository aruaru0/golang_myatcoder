package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

//　セグメント木
const mod = 998244353

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, K := getInt(), getInt()
	l := make([]int, K)
	r := make([]int, K)
	for i := 0; i < K; i++ {
		l[i], r[i] = getInt(), getInt()
	}

	dp := make([]int, N+10)
	tot := make([]int, N+10)
	dp[1] = 1
	tot[1] = 1
	for i := 2; i <= N; i++ {
		for k := 0; k < K; k++ {
			lp := max(1, i-r[k])
			rp := i - l[k]
			if rp < 0 {
				continue
			}
			dp[i] += tot[rp] - tot[lp-1]
			if dp[i] < 0 {
				dp[i] += mod
			}
			dp[i] %= mod
		}
		tot[i] = tot[i-1] + dp[i]
		tot[i] %= mod
	}
	out(dp[N])
}
