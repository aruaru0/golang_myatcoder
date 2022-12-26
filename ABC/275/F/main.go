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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	a := getInts(N)

	dp := make([][2]int, M+1)
	for i := 0; i < M+1; i++ {
		dp[i][0] = inf
		dp[i][1] = inf
	}
	dp[0][1] = 0
	for i := 0; i < N; i++ {
		tmp := make([][2]int, M+1)
		for j := 0; j < M+1; j++ {
			tmp[j][0] = inf
			tmp[j][1] = inf
		}
		// out("--------")
		for j := 0; j <= M; j++ {
			// 選択×→選択×
			chmin(&tmp[j][0], dp[j][0])
			// 選択〇→選択×
			chmin(&tmp[j][0], dp[j][1]+1)
			// 選択×→選択〇
			if j+a[i] <= M {
				chmin(&tmp[(j + a[i])][1], dp[j][0])
				// 選択〇→選択〇
				chmin(&tmp[(j + a[i])][1], dp[j][1])
			}
		}
		// for j := 0; j <= M; j++ {
		// 	out(j, ":", tmp[j])
		// }
		dp = tmp
	}

	for i := 1; i <= M; i++ {
		ans := min(dp[i][0], dp[i][1])
		if ans == inf {
			out(-1)
		} else {
			out(ans)
		}
	}
}
