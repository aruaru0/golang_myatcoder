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
	N := getI()
	a := getInts(N)

	// 平均を求める処理
	l, r := 0.0, 1e10
	for i := 0; i < 100; i++ {
		x := (l + r) / 2
		b := make([]float64, N)
		for i := 0; i < N; i++ {
			b[i] = float64(a[i]) - x
		}
		dp := make([][2]float64, N+1)
		for i := 0; i < N; i++ {
			// この桁は使わない場合は、手前の桁は使っている必要がある
			dp[i+1][0] = dp[i][1]
			// この桁を使う場合は、どちらでもよい
			dp[i+1][1] = math.Max(dp[i][0]+b[i], dp[i][1]+b[i])
		}
		v := math.Max(dp[N][0], dp[N][1])
		if v < 0 {
			r = x
		} else {
			l = x
		}
	}
	out(l)

	// 中央値を求める処理
	ll, rr := 0, int(1e10)
	for ll+1 != rr {
		m := (ll + rr) / 2
		b := make([]int, N)
		cnt1 := 0
		for i := 0; i < N; i++ {
			if a[i] >= m {
				b[i] = 1
				cnt1++
			}
		}
		cnt0 := 0
		cnt := 0
		for i := 0; i < N; i++ {
			if b[i] == 1 {
				cnt0 += cnt / 2
				cnt = 0
			} else {
				cnt++
			}
		}
		cnt0 += cnt / 2
		if cnt0 >= cnt1 {
			rr = m
		} else {
			ll = m
		}
		// out(m, b, "cnt", cnt0, cnt1, "lr", ll, rr)
	}
	out(ll)
}
