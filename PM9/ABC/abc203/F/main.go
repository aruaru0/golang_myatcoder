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
const M = 31

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	a := getInts(N)
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	// i個目まで見たときに、j回高橋くんが取ったときの青木君の回数の最小値
	dp := make([][M + 1]int, N+1)
	for i := 0; i <= N; i++ {
		for j := 0; j < M; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	r := 0
	for i := 0; i < N; i++ {
		// a[i]を高橋くんが取る場合、a[i]/2以上の草はすべてなくなる
		// ので、次に残る部分までrを進める
		for r < N && a[r]*2 > a[i] {
			r++
		}

		for j := 0; j < M; j++ {
			// 青木君が取る場合は、a[i]が１つ減って回数が１増える
			chmin(&dp[i+1][j], dp[i][j]+1)
			// 高橋くんが１つ取った場合、次はr個になるので以下の式となる
			chmin(&dp[r][j+1], dp[i][j])
		}
	}

	// 条件を満たす最小の回数を出力
	for j := 0; j < M; j++ {
		if dp[N][j] <= K {
			out(j, dp[N][j])
			return
		}
	}
}
