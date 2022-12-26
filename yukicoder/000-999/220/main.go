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

func main() {
	sc.Split(bufio.ScanWords)
	p := getInt()

	var ck = func(n, d int) int {
		if d == 3 {
			return 3
		}
		if n == 3 {
			return 3
		}
		return (n + d) % 3
	}

	s := "1"
	for i := 0; i < p; i++ {
		s += "0"
	}
	// out(s)
	var dp [21][2][4]int
	dp[0][0][0] = 1
	for i := 0; i < len(s); i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 10; k++ {
				dp[i+1][1][ck(j, k)] += dp[i][1][j]
			}
			d := int(s[i] - '0')
			for k := 0; k < d; k++ {
				dp[i+1][1][ck(j, k)] += dp[i][0][j]
			}
			dp[i+1][0][ck(j, d)] += dp[i][0][j]
		}
	}

	// for i := 0; i <= len(s); i++ {
	// 	out(dp[i])
	// }
	out(dp[len(s)][1][0] + dp[len(s)][1][3] - 1)
}
