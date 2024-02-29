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
	N, L, D := getI(), getI(), getI()

	// DP
	dp1 := make([]float64, N+1)
	dp1[0] = 1.0
	// 累積和
	s1 := make([]float64, N+2)
	s1[1] = 1.0

	// ディーラー側は、サイコロを振るだけ
	for i := 1; i < N+1; i++ {
		// iになる確率は、1つ前の出目の合計確率 / サイコロの面数
		dp1[i] = math.Max(0.0, s1[min(L, i)]-s1[max(0, i-D)]) / float64(D)
		s1[i+1] = s1[i] + dp1[i]
	}

	// プレーヤー側
	// Nを超えたら負け（dp2[N+1] = 0)
	dp2 := make([]float64, N+1)
	s2 := make([]float64, N+2)
	for i := N; i >= 0; i-- {
		// ディーラー側が負ける条件
		p1 := 1.0 - (s1[N+1] - s1[max(L, i)])
		// プレーヤー側が勝つ条件
		p2 := (s2[i+1] - s2[min(N+1, i+D+1)]) / float64(D)
		// どちらか大きい方
		dp2[i] = math.Max(p1, p2)
		s2[i] = s2[i+1] + dp2[i]
	}

	out(dp2[0])
}
