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

func solve(N, K int, a []int) {
	x := 0
	for i := 0; i < K; i++ {
		out(x%N, a[x%N])
		x += a[x%N]
	}
	out(x)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	a := getInts(N)

	s := []int{0}
	mm := make(map[int]bool)
	m := make(map[int]int)
	x := 0
	lpos := -1
	curK := 0
	for i := 0; i < K; i++ {
		// out(x, x%N, mm[x%N], s)
		if mm[x%N] {
			lpos = m[x%N]
			break
		}
		mm[x%N] = true
		m[x%N] = len(s)
		x += a[x%N]
		s = append(s, x)
		curK++
	}

	// 同じ値が出た時点でストップし、
	// Kまできていれば終了
	if curK == K {
		out(s[len(s)-1])
		return
	}
	// そうで無ければ残りを計算し、
	// ループ回数×ループで加算される値　＋　残りを加算
	diff := K - curK
	cnt := len(s) - lpos
	tot := s[len(s)-1] - s[lpos-1]
	num := diff / cnt
	rest := diff % cnt
	ans := s[len(s)-1] + num*tot + (s[lpos+rest-1] - s[lpos-1])
	out(ans)
}
