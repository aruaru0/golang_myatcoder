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

var N, P int
var dp [3100][9100][2]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, P = getI(), getI()

	// [i][j][k] i列目まで決まっていて、削除した辺がｊ個で、
	// k = 0非連結　or 1連結の時の組み合わせの数
	dp[1][0][1] = 1 // １列目で削除が０ならつながっている（連結）
	dp[1][1][0] = 1 // １列目で削除が１ならつながっていない（非連結）

	for i := 1; i < N; i++ {
		for del := 0; del < 9000; del++ {
			// 下と縦の２本をカットした場合、非連結に
			dp[i+1][del+2][0] += dp[i][del][1]
			dp[i+1][del+2][0] %= P

			// 上と縦の２本をカットした場合、非連結に
			dp[i+1][del+2][0] += dp[i][del][1]
			dp[i+1][del+2][0] %= P

			// 上と縦を残した場合、連結なら連結に
			dp[i+1][del+1][1] += dp[i][del][1]
			dp[i+1][del+1][1] %= P

			// 下と縦を残した場合、連結なら連結に
			dp[i+1][del+1][1] += dp[i][del][1]
			dp[i+1][del+1][1] %= P

			// 上下を残す。連結状態は継承
			dp[i+1][del+1][1] += dp[i][del][1]
			dp[i+1][del+1][1] %= P
			dp[i+1][del+1][0] += dp[i][del][0]
			dp[i+1][del+1][0] %= P

			// 全部残す。連結状態に変化
			dp[i+1][del][1] += dp[i][del][0]
			dp[i+1][del][1] %= P
			dp[i+1][del][1] += dp[i][del][1]
			dp[i+1][del][1] %= P
		}
	}

	for i := 1; i < N; i++ {
		fmt.Fprint(wr, dp[N][i][1], " ")
	}
	out()
}
