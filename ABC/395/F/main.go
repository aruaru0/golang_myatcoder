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
	const inf = int(1e18)
	N, X := getI(), getI()

	u, d := make([]int, N), make([]int, N)
	sum := 0
	hmax := inf
	for i := 0; i < N; i++ {
		u[i], d[i] = getI(), getI()
		sum += u[i] + d[i]
		hmax = min(hmax, u[i]+d[i])
	}

	// 指定されたHで、上の歯を条件に合わせることが可能か？
	f := func(H int) bool {
		// １つ目の歯の範囲はmax(1, H-d[0])~min(u[0], H-1)
		// 歯は削れ、足してHにする必要がある
		// 削らない場合のはの高さ、またはH-1を超えない
		l, r := max(1, H-d[0]), min(u[0], H-1)
		if l > r {
			return false
		}
		for i := 1; i < N; i++ {
			L, R := max(1, H-d[i]), min(u[i], H-1)
			if L > R {
				return false
			}
			// 手前の歯の区間+-Xの範囲と、現在の歯の範囲を合成
			l = max(L, l-X)
			r = min(R, r+X)
			if l > r {
				return false
			}
		}
		// 問題なければ実現可能とする
		return true
	}

	// 削る方しかできない
	// u, d >=1なので、2であれば必ずokとなる
	ok, ng := 2, hmax+1
	for ok+1 != ng {
		mid := (ng + ok) / 2
		if f(mid) { // 成立したらOKにする
			ok = mid
		} else {
			ng = mid
		}
	}

	ans := sum - ok*N
	out(ans)
}
