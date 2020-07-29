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

const N = 10000

func main() {
	sc.Split(bufio.ScanWords)
	K := getInt()

	var dp [N][210]float64
	dp[0][0] = 1
	for i := 1; i < N; i++ {
		for j := 0; j < K; j++ {
			for k := 1; k <= 6; k++ {
				if j+k > K {
					dp[i][0] += dp[i-1][j] * 1.0 / 6.0
				} else {
					dp[i][j+k] += dp[i-1][j] * 1.0 / 6.0
				}
			}
		}
		// out(dp[i])
	}

	ans := float64(0.0)
	for i := 1; i < N; i++ {
		// out(dp[i][K])
		ans += float64(i) * dp[i][K]
	}
	out(ans)
}
