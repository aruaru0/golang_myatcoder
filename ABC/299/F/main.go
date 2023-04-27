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
const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getS()
	n := len(s)

	w := make([][130]int, n+1)
	for i := 0; i <= n; i++ {
		for j := 0; j < 130; j++ {
			w[i][j] = inf
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := int('a'); j <= int('z'); j++ {
			if int(s[i]) != j {
				w[i][j] = w[i+1][j]
			} else {
				w[i][j] = i + 1
			}
		}
	}

	ans := 0
	for x := 1; x < n; x++ {
		dp := make([][]int, n+1)
		for i := 0; i <= n; i++ {
			dp[i] = make([]int, n+1)
		}
		t1, t2 := w[0][int(s[x])], x+1
		if t1 < t2 && t2 < inf {
			dp[t1][t2]++
		}
		for i := 0; i <= x; i++ {
			for j := x + 1; j <= n; j++ {
				for k := int('a'); k <= int('z'); k++ {
					t1, t2 := w[i][k], w[j][k]
					if t1 < t2 && t2 < inf {
						dp[t1][t2] += dp[i][j]
						dp[t1][t2] %= mod

					}
				}
			}
		}
		for i := 0; i <= x; i++ {
			for j := x + 1; j <= n; j++ {
				if w[i][int(s[x])] == x+1 {
					ans += dp[i][j]
					ans %= mod
				}
			}
		}
	}

	out(ans)
}
