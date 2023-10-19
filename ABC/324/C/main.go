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

func calc(t, s string) int {
	n := len(t)
	m := len(s)

	if abs(n-m) >= 2 {
		return abs(n - m)
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < m+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			cost := 0
			if t[i-1] != s[j-1] {
				cost = 1
			}
			dp[i][j] = nmin(
				dp[i-1][j]+1,
				dp[i][j-1]+1,
				dp[i-1][j-1]+cost)
			if i == j && dp[i][j] >= 3 {
				return 100
			}

			// out("----")
			// for k := 0; k < n; k++ {
			// 	out(dp[k])
			// }
		}
	}
	return dp[n][m]
}

func calc2(t, s string) bool {
	n := len(t)
	m := len(s)

	diff := abs(n - m)
	// ２つの文字列の長さが２以上
	if diff >= 2 {
		return false
	}
	if diff == 0 {
		// 文字列の違いをカウント
		cnt := 0
		for i := 0; i < n; i++ {
			if s[i] != t[i] {
				cnt++
			}
		}
		return cnt < 2
	}
	// out("pass")
	// out(t)
	// out(s)
	// 1文字違い
	l, r := 0, 0
	for i := 0; l < n && r < m; i++ {
		if t[l] == s[r] {
			l++
			r++
		} else {
			if n > m {
				l++
			} else {
				r++
			}
		}
	}
	// out("l, r", l, r)
	if abs(l-r) <= 1 {
		return true
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, T := getI(), getS()

	ans := make([]int, 0)
	for i := 0; i < N; i++ {
		s := getS()
		ret := calc2(T, s)
		// out(ret)
		if ret {
			ans = append(ans, i)
		}
	}

	out(len(ans))
	for _, e := range ans {
		fmt.Fprint(wr, e+1, " ")
	}
	out()
}
