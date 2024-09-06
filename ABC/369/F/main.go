package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
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

func bisectRight(dp []pair, target pair) int {
	lo, hi := 0, len(dp)
	for lo < hi {
		mid := (lo + hi) / 2
		if dp[mid].lessThan(target) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

type pair struct {
	first, second int
}

func (p pair) lessThan(other pair) bool {
	if p.first != other.first {
		return p.first < other.first
	}
	return p.second < other.second
}

func repeatString(char string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(char, count)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	H, W, N := getI(), getI(), getI()
	RC := make([]pair, N)
	for i := 0; i < N; i++ {
		RC[i].first, RC[i].second = getI(), getI()
	}

	sort.Slice(RC, func(i, j int) bool {
		if RC[i].first != RC[j].first {
			return RC[i].first < RC[j].first
		}
		return RC[i].second < RC[j].second
	})

	INF := int(1e9)
	dp := make([]pair, N)
	for i := range dp {
		dp[i] = pair{INF, INF}
	}

	pre := make(map[pair]pair)

	for _, rc := range RC {
		r, c := rc.first, rc.second
		pos := bisectRight(dp, pair{c, r})
		dp[pos] = pair{c, r}
		if pos > 0 {
			pre[pair{c, r}] = dp[pos-1]
		} else {
			pre[pair{c, r}] = pair{1, 1}
		}
	}

	for dp[len(dp)-1] == (pair{INF, INF}) {
		dp = dp[:len(dp)-1]
	}

	// 最後の位置から解を組み立てる
	lastPair := dp[len(dp)-1]
	sol := []string{repeatString("R", W-lastPair.first) + repeatString("D", H-lastPair.second)}
	x, y := lastPair.first, lastPair.second

	for x != 1 || y != 1 {
		px, py := pre[pair{x, y}].first, pre[pair{x, y}].second
		sol = append(sol, repeatString("R", x-px)+repeatString("D", y-py))
		x, y = px, py
	}

	// 結果を反転して出力
	for i := len(sol) - 1; i >= 0; i-- {
		fmt.Fprint(wr, sol[i])
	}
	out()
}
