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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	X, Y, Z := getI(), getI(), getI()
	s := getS()

	n := len(s)
	// dp[i][j] : i文字目まで入力してCapsLockの状態がjときに要した入力時間の最小値
	dp := make([][2]int, n+1)
	dp[0][0] = 0
	dp[0][1] = Z
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			// 次caps = offになる場合、　caps:offでaを押す、 onでshift+aの後切り替える、onで切り替えてaを押す
			dp[i+1][0] = nmin(dp[i][0]+X, dp[i][1]+Y+Z, dp[i][1]+X+Z)
			// 次caps = onになる場合、 offで切り替えてaを押す、onでshift+aを押す
			dp[i+1][1] = nmin(dp[i][0]+X+Z, dp[i][1]+Y)
		} else {
			dp[i+1][1] = nmin(dp[i][1]+X, dp[i][0]+Y+Z, dp[i][0]+X+Z)
			dp[i+1][0] = nmin(dp[i][1]+X+Z, dp[i][0]+Y)
		}
	}
	out(min(dp[n][0], dp[n][1]))
}
