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

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	a := getInts(N)

	var dp [101][101][10100]int32
	// for i := 0; i <= N; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		for k := 0; k < 10100; k++ {
	// 			dp[i][j][k] = -1
	// 		}
	// 	}
	// }

	dp[0][0][0] = 1
	for i := 1; i <= N; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k <= 10000; k++ {
				if dp[i-1][j][k] != 0 {
					//					out("pass", i, j, k, dp[i-1][j][k], a[i-1])
					dp[i][j][k] += dp[i-1][j][k]
					dp[i][j][k] %= int32(mod)
					dp[i][j+1][k+a[i-1]] += dp[i-1][j][k]
					dp[i][j+1][k+a[i-1]] %= int32(mod)
				}
			}
			//			out(i, j, dp[i][j][:100*N])
		}
	}
	ans := 0
	for i := 1; i <= N; i++ {
		for j := K * i; j <= 10000; j++ {
			ans += int(dp[N][i][j])
			ans %= mod
		}
	}
	out(ans)
}
