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

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	s := getString()
	N := len(s)
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, 13)
	}
	dp[0][0] = 1

	for i := 0; i < N; i++ {
		if s[i] == '?' {
			for j := 0; j < 13; j++ {
				for x := 0; x < 10; x++ {
					dp[i+1][(10*j+x)%13] += dp[i][j]
					dp[i+1][(10*j+x)%13] %= mod
				}
			}
		} else {
			for j := 0; j < 13; j++ {
				x := int(s[i] - '0')
				dp[i+1][(10*j+x)%13] += dp[i][j]
				dp[i+1][(10*j+x)%13] %= mod
			}
		}
	}
	// for i := 0; i <= N; i++ {
	// 	out(dp[i])
	// }

	out(dp[N][5])
}
