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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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
	N, D := getI(), getI()
	a := getInts(N)
	const maxn = int(1e6 + 10)
	b := make([]int, maxn)
	for i := 0; i < N; i++ {
		b[a[i]]++
	}

	if D == 0 {
		// 同じ値は１つだけしか残せない
		ans := 0
		for i := 0; i < maxn; i++ {
			ans += max(0, b[i]-1)
		}
		out(ans)
		return
	}
	const inf = int(1e18)
	ans := 0
	for i := 0; i < D; i++ {
		cnt := make([]int, 0)
		for j := i; j < maxn; j += D {
			cnt = append(cnt, b[j])
		}
		dp := make([][2]int, len(cnt)+1)
		for j := 0; j < len(cnt); j++ {
			dp[j][0] = inf
			dp[j][1] = inf
		}
		dp[0][0] = 0
		for j := 0; j < len(cnt); j++ {
			// j個目の値を消す場合は、cnt[j]個消す。遷移はどちらからでもOK
			dp[j+1][0] = min(dp[j][0]+cnt[j], dp[j][1]+cnt[j])
			// j個目を残す場合は、1つ前は消してなければダメ
			dp[j+1][1] = dp[j][0]
		}
		ans += min(dp[len(cnt)][0], dp[len(cnt)][1])
		// out(i, cnt)
	}
	out(ans)
}
