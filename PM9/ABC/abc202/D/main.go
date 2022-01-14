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

func nCr(n, r int) int {
	ret := 1
	for i := 1; i <= r; i++ {
		ret *= n - i + 1
		ret /= i
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	A, B, K := getI(), getI(), getI()-1

	n := A + B
	ans := ""
	// 先頭の文字がaになるかどうか判定し、それによって処理を変える
	//  axxxxxxとしてxxxxxxの部分パターン数を求める(nCr(A+B-1, B)個)
	//  Kより大きい場合は、K番目は前半部に存在＝桁をaと決める
	//  Kより小さい場合は、K番目は後半部に存在＝桁はbと決める
	//   ※ｂと決めた場合、K番目を前半部分だけ減らしておく
	//  これを繰り返すと、最後の桁まで決まる
	for i := 0; i < n; i++ {
		if 0 < A {
			if K < nCr(A+B-1, B) {
				ans += "a"
				A--
			} else {
				ans += "b"
				K -= nCr(A+B-1, B)
				B--
			}
		} else {
			ans += "b"
			B--
		}
	}
	out(ans)
}
