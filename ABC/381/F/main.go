package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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
	N := getI()
	a := getInts(N)
	for i := 0; i < N; i++ {
		a[i]--
	}

	m := 20
	m2 := 1 << m

	is := make([][]int, m)
	for i := 0; i < N; i++ {
		is[a[i]] = append(is[a[i]], i)
	}

	const inf = int(1e18)

	// cをi番目以降で検索し、２つとった場合の場所を返す
	getNext := func(c, i int) int {
		nis := is[c]
		j := lowerBound(nis, i)
		j += 1
		if j < len(nis) {
			return nis[j] + 1
		}
		return inf
	}

	// dpはSを作成した場合の最小のインデックス
	dp := make([]int, m2)
	for i := 0; i < m2; i++ {
		dp[i] = inf
	}
	dp[0] = 0

	ans := 0
	for s := 0; s < m2; s++ {
		// 集合Sの1の数を数える
		pc := bits.OnesCount(uint(s))
		// もし集合Sがすでに計算済みなら、答えを更新
		if dp[s] != inf {
			ans = max(ans, pc)
		}
		// ビットを確認
		for c := 0; c < m; c++ {
			// すでにbitが立っていればスキップ
			if (s>>c)%2 == 1 {
				continue
			}
			// bitが立っていない場合は、getNextを呼び出す
			chmin(&dp[s|1<<c], getNext(c, dp[s]))
		}
	}
	out(ans * 2)
}
