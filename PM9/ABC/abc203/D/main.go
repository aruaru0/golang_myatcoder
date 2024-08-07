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

var N, K int
var a [][]int

func f(th int) bool {
	b := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		b[i] = make([]int, N+1)
	}
	// th以下の場合は1を立てる
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if a[i][j] <= th {
				b[i+1][j+1] = 1
			}
		}
	}

	// 累積和を計算
	for i := 0; i <= N; i++ {
		for j := 0; j < N; j++ {
			b[i][j+1] += b[i][j]
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j <= N; j++ {
			b[i+1][j] += b[i][j]
		}
	}

	// 中央値の場合の数を計算
	n := K*K - (K*K/2 + 1) + 1

	for i := 0; i <= N-K; i++ {
		for j := 0; j <= N-K; j++ {
			// K x K のサイズをずらしながら、個数が閾値以上かどうか調べる
			tot := b[i+K][j+K] - b[i+K][j] - b[i][j+K] + b[i][j]
			// out(tot)
			if tot >= n {
				// totがn以上、つまり、中央値はth以下の部分にある
				return true
			}
		}
	}

	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K = getI(), getI()
	a = make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInts(N)
	}

	l, r := -1, int(1e9+1)
	for l+1 != r {
		m := (l + r) / 2
		// 中央値がm以上かどうかをチェック
		ret := f(m)
		if ret {
			r = m
		} else {
			l = m
		}
	}

	out(r)
}
