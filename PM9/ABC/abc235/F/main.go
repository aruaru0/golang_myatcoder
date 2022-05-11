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
	sum, cnt int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getS()
	m := getI()
	nn := len(N)
	c := getInts(m)
	n := make([]int, nn)
	for i := 0; i < nn; i++ {
		n[i] = int(N[i] - '0')
	}
	const mx = 1 << 10
	const mod = 998244353
	dp := make([]pair, mx)
	k, spre := 0, 0
	for i := 0; i < nn; i++ {
		tmp := make([]pair, mx)
		// 各ビットパターンに対してdpを更新
		for s := 0; s < mx; s++ {
			for d := 0; d < 10; d++ {
				tmp[s|1<<d].cnt += dp[s].cnt
				// 和は現在の和の１０倍＋末尾の桁の値×現在までの個数
				tmp[s|1<<d].sum += dp[s].sum*10 + dp[s].cnt*d
			}
		}
		// i桁から開始する値
		if i != 0 {
			for d := 0; d < 10; d++ {
				if d != 0 {
					tmp[1<<d].cnt += 1
					tmp[1<<d].sum += d
				}
			}
		}
		// 上がspre(上の桁が上限値のため、それ以下の値しか使えない)
		for d := 0; d < n[i]; d++ {
			if i != 0 || d != 0 {
				tmp[k|1<<d].cnt += 1
				tmp[k|1<<d].sum += spre*10 + d
			}
		}

		for s := 0; s < mx; s++ {
			tmp[s].cnt %= mod
			tmp[s].sum %= mod
		}
		k |= 1 << n[i]
		spre = (spre*10 + n[i]) % mod
		tmp, dp = dp, tmp
	}

	p := 0
	for i := 0; i < m; i++ {
		p |= 1 << c[i]
	}

	ans := 0
	for s := 0; s < mx; s++ {
		if (s & p) == p {
			ans += dp[s].sum
			ans %= mod
		}
	}
	if (k & p) == p {
		ans += spre
		ans %= mod
	}

	out(ans)
}
