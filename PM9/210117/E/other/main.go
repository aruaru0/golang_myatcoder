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
	N, M := getI(), getI()
	a := getInts(N)
	// 　a[i]-a[j]までの総和は
	//   a[j] - a[0] - (a[i]-a[0])までの総和となり
	//   a[i]-a[j]までのＭで割ったあまりは
	//   (a[i]-a[j])%M =
	//   ((a[j]-a[0])-(a[i]-a[0]))%M =
	//   (a[j]-a[0])%M-(a[i]-a[0])%M と同じ
	//   つまり
	//     (a[j]-a[0])%M -(a[i]-a[0])%M = 0
	//      -->  (a[j]-a[0])%M = (a[i]-a[0])%M
	//   であればＭで割り切れることになる
	// なので、結局、0-iまでの総和のＭで割ったあまりを求めておき
	// あまりが同じ部分をカウントする問題になる
	a = append([]int{0}, a...)
	m := make(map[int]int)
	for i := 1; i <= N; i++ {
		a[i] += a[i-1]
		a[i] %= M
		m[a[i]]++
	}
	// 0は１つでもＯＫなので1足しておく
	// 0以外は2つあればＯＫ
	m[0]++
	cnt := 0
	for _, e := range m {
		// eC2を計算
		cnt += e * (e - 1) / 2
	}

	out(cnt)
}
