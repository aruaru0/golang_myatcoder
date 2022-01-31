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

const mod = 1000000007

var node = make([][]int, 2020)
var dp = [2020][2020][3]int{}

func dfs(u, v int) int {
	s := 1
	dp[v][0][0] = 1
	dp[v][1][1] = 1
	dp[v][1][2] = 1
	for _, w := range node[v] {
		if w != u {
			t := dfs(v, w)
			for k := s; k >= 0; k-- {
				for l := 1; l < t+1; l++ {
					dp[v][k+l][0] += dp[v][k][0] * ((dp[w][l][0] + dp[w][l][1] - dp[w][l-1][0] + mod) % mod) % mod
					dp[v][k+l][0] %= mod
					dp[v][k+l][1] += dp[v][k][1] * ((dp[w][l][0] + dp[w][l][1] + dp[w][l][2] - dp[w][l-1][0] + mod) % mod) % mod
					dp[v][k+l][1] %= mod
					dp[v][k+l][2] += dp[v][k][2] * (dp[w][l][1] + dp[w][l][2]) % mod
					dp[v][k+l][2] %= mod
				}
				dp[v][k][2] = 0
			}
			s += t
			s %= mod
		}
	}
	return s
}


// 現状のスキルでは厳しめ。とりあえず、コピーして内容を確認
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	for i := 1; i < n; i++ {
		u, v := getI(), getI()
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	dfs(0, 1)

	for k := 0; k < n+1; k++ {
		if k != 0 {
			out((dp[1][k][0] + dp[1][k][1] + dp[1][k][2] - dp[1][k-1][0] + mod) % mod)
		} else {
			out((dp[1][k][0] + dp[1][k][1] + dp[1][k][2]) % mod)
		}
	}
}
