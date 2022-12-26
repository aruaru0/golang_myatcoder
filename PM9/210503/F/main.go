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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	w := make([]int, N)
	maxW := 0
	for i := 0; i < N; i++ {
		w[i] = getI()
		chmax(&maxW, w[i])
	}
	l := make([]int, M)
	v := make([]int, M)
	for i := 0; i < M; i++ {
		l[i], v[i] = getI(), getI()
		if v[i] < maxW {
			out(-1)
			return
		}
	}
	// すべての組み合わせの重さを計算し、必要距離を演算しておく
	dist := make([]int, 1<<N)
	for bit := 0; bit < 1<<N; bit++ {
		W := 0
		for i := 0; i < N; i++ {
			if (bit>>i)%2 == 1 {
				W += w[i]
			}
		}
		for i := 0; i < M; i++ {
			if v[i] < W { // もし橋の加重よりＷが大きければ。最大長を更新
				chmax(&dist[bit], l[i])
			}
		}
	}
	// dp
	res := int(1e18)
	ids := make([]int, N)
	for i := 0; i < N; i++ {
		ids[i] = i
	}
	for { // 総当たり
		dp := make([]int, N)
		for i := 1; i < N; i++ {
			bit := 1 << ids[i] // 現在位置のラクダのビットを立てる
			for j := i - 1; j >= 0; j-- {
				bit |= 1 << ids[j]             // 手前のラクダを追加していって
				chmax(&dp[i], dp[j]+dist[bit]) // 一番幅をとらなければならないパターンを検出
			}
		}
		chmin(&res, dp[N-1])

		if !NextPermutation(sort.IntSlice(ids)) {
			break
		}
	}
	out(res)
}
