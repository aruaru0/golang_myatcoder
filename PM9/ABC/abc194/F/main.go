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

const mod = int(1e9 + 7)

func to_int(c byte) int {
	if 47 < c && c < 58 {
		return int(c) - 48
	}
	return int(c) - 55
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s, k := getS(), getI()

	n := len(s)
	dp := make([][17]int, n+1)
	m := map[int]struct{}{}
	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		dp[i+1][0]++
		dp[i+1][0] %= mod
		for j := 0; j < 17; j++ {
			dp[i+1][j] += dp[i][j] * j
			dp[i+1][j] %= mod
			if j < 16 {
				if j == 0 {
					dp[i+1][j+1] += dp[i][j] * (16 - j - 1)
				} else {
					dp[i+1][j+1] += dp[i][j] * (16 - j)
				}
				dp[i+1][j+1] %= mod
			}
		}
		j := 0
		if i == 0 {
			j = 1
		}
		for ; j < to_int(s[i]); j++ {
			cnt := 1
			if _, ok := m[j]; ok {
				cnt = 0
			}
			dp[i+1][len(m)+cnt]++
			dp[i+1][len(m)+cnt] %= mod
		}
		m[to_int(s[i])] = struct{}{}
	}

	if len(m) == k {
		out((dp[n][k] + 1) % mod)
	} else {
		out(dp[n][k] % mod)
	}
}
