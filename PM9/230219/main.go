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

type RollingHash struct {
	base int
	mod  int
	pw   []int
	h    []int
}

// 作成　入力は文字列
func newRollingHash(s string) *RollingHash {
	var ret RollingHash
	ret.base = 37
	ret.mod = int(1e9 + 7)
	n := len(s) + 1
	ret.pw = make([]int, n)
	ret.h = make([]int, n)
	for i := 0; i < n; i++ {
		ret.pw[i] = 1
	}
	v := 0
	l := len(s)
	for i := 0; i < l; i++ {
		v = (v*ret.base + int(s[i])) % ret.mod
		ret.h[i+1] = v
	}
	v = 1
	for i := 0; i < l; i++ {
		v = v * ret.base % ret.mod
		ret.pw[i+1] = v
	}
	return &ret
}

// [l, r)の値を調べる
func (rh *RollingHash) get(l, r int) int {
	ret := (rh.h[r] - rh.h[l]*rh.pw[r-l]) % rh.mod
	if ret < 0 {
		ret += rh.mod
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()
	s := []byte(getS())
	t := make([]byte, 0)
	for i := N - 1; i >= 0; i-- {
		t = append(t, s[i])
	}

	rs := newRollingHash(string(s))
	rt := newRollingHash(string(t))

	for q := 0; q < Q; q++ {
		l, r := getI()-1, getI()-1
		v0 := rs.get(l, r+1)
		v1 := rt.get(N-1-r, N-1-l+1)
		if v0 == v1 {
			out("Yes")
		} else {
			out("No")
		}
	}
}
