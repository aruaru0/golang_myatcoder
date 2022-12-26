package main

import (
	"bufio"
	"fmt"
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

func f() {
	n1, n2 := getI(), getI()
	m := getI()
	a := getInts(m)
	sort.Ints(a)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n1+1)
		for j := 0; j <= n1; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	sum, ans := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j <= n1; j++ {
			if dp[i][j] == -1 {
				continue
			}
			ans = max(ans, dp[i][j])
			jj := sum - j
			if j+a[i] <= n1 {
				dp[i+1][j+a[i]] = dp[i][j] + 1
			}
			if jj >= 0 && jj+a[i] <= n2 {
				dp[i+1][j] = dp[i][j] + 1
			}
		}
		sum += a[i]
	}
	for i := 0; i <= n1; i++ {
		ans = max(ans, dp[m][i])
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	d := getI()
	for i := 0; i < d; i++ {
		f()
	}
}
