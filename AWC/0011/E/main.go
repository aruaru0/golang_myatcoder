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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N, M := getI(), getI()

	A := make([]int, N+1)
	B := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i], B[i] = getI(), getI()
	}

	// dpL[i][j] : 先頭からi個目までの宝箱のうち、重さ合計j以下で選んだときの価値の最大値
	dpL := make([][]int, N+2)
	for i := 0; i <= N+1; i++ {
		dpL[i] = make([]int, M+1)
	}

	for i := 1; i <= N; i++ {
		for j := 0; j <= M; j++ {
			// i番目を選ばない場合
			dpL[i][j] = dpL[i-1][j]
			// i番目を選ぶ場合
			if j >= A[i] {
				dpL[i][j] = max(dpL[i][j], dpL[i-1][j-A[i]]+B[i])
			}
		}
	}

	// dpR[i][j] : 後ろから(i番目以降の)宝箱のうち、重さ合計j以下で選んだときの価値の最大値
	dpR := make([][]int, N+2)
	for i := 0; i <= N+1; i++ {
		dpR[i] = make([]int, M+1)
	}

	for i := N; i >= 1; i-- {
		for j := 0; j <= M; j++ {
			dpR[i][j] = dpR[i+1][j]
			if j >= A[i] {
				dpR[i][j] = max(dpR[i][j], dpR[i+1][j-A[i]]+B[i])
			}
		}
	}

	// 全体の最適解（最大価値）
	maxVal := dpL[N][M]

	for i := 1; i <= N; i++ {
		canBeOptimal := false
		// i番目の宝箱を必ず選ぶ場合の最大価値を求める
		// 残りの容量 (M - A[i]) を、左側と右側でどう分割するかを全探索
		for w := 0; w <= M-A[i]; w++ {
			val := dpL[i-1][w] + dpR[i+1][M-A[i]-w] + B[i]
			if val == maxVal {
				canBeOptimal = true
				break
			}
		}

		if canBeOptimal {
			out("Yes")
		} else {
			out("No")
		}
	}
}
