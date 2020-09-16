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

var a []int
var l, r []int

const maxN = 110

func solve() {
	n := getInt()
	a = getInts(n)

	var dp [maxN][maxN]int
	var dp2 [maxN][maxN]int

	for i := 0; i < maxN; i++ {
		for j := 0; j < maxN; j++ {
			dp[i][j] = -1
			dp2[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
		dp2[i][i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			for k := i + 1; k < n; k++ {
				if a[k]-a[i] > 0 && a[k]-a[i] > a[i]-a[j] {
					dp[k][i] = max(dp[k][i], dp[i][j]+1)
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			for k := i - 1; k >= 0; k-- {
				if a[k]-a[i] > 0 && abs(a[k]-a[i]) > abs(a[i]-a[j]) {
					dp2[k][i] = max(dp2[k][i], dp2[i][j]+1)
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k < n; k++ {
				ans = max(ans, dp[i][j]+dp2[i][k])
			}
		}
	}
	out(ans - 1)
}

func main() {
	sc.Split(bufio.ScanWords)
	T := getInt()
	for i := 0; i < T; i++ {
		solve()
	}
}
