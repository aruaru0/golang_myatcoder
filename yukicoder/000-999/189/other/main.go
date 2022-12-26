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

var dp [2][210][2010][2]int

const mod = int(1e9 + 9)

func makeDP(s string, a int) {
	n := len(s)
	dp[a][0][0][1] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 2000; j++ {
			for k := 0; k < 2; k++ {
				if k == 0 {
					for b := 0; b < 10; b++ {
						dp[a][i+1][j+b][k] += dp[a][i][j][k]
						dp[a][i+1][j+b][k] %= mod
					}
				} else {
					t := int(s[i] - '0')
					for b := 0; b < t; b++ {
						dp[a][i+1][j+b][0] += dp[a][i][j][1]
						dp[a][i+1][j+b][0] %= mod
					}
					dp[a][i+1][j+t][1] += dp[a][i][j][1]
					dp[a][i+1][j+t][1] %= mod
				}
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s, t := getString(), getString()
	makeDP(s, 0)
	makeDP(t, 1)
	ans := 0
	for i := 1; i <= 2000; i++ {
		x := dp[0][len(s)][i][0] + dp[0][len(s)][i][1]
		y := dp[1][len(t)][i][0] + dp[1][len(t)][i][1]
		ans += x * y % mod
		ans %= mod
	}
	out(ans)
}
