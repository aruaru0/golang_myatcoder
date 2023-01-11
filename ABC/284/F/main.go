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

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	T := getS()
	rT := Reverse(T)

	r0 := newRollingHash(T)
	r1 := newRollingHash(rT)

	// abcbac  ab...c  2
	// cabcba  c...ba  3-2 = 1
	ans := ""
	cnt := -1
	for i := 0; i <= N; i++ {
		l, r := i, N-i
		// out(i, "--------")
		s0 := r0.get(0, l)
		s1 := r0.get(2*N-r, 2*N)
		t0 := r1.get(r, N)
		t1 := r1.get(N, N+r)
		// out(l, r, "-->", 0, l, ":", 2*N-r, 2*N, ":", r, N, ":", N, N+r)
		// out(s0, s1, t0, t1)
		if s0 == t0 && s1 == t1 {
			ans = T[:l] + T[2*N-r:]
			cnt = l
			break
		}
	}

	// out(r0.get(0, 1))
	// out(r0.get(4, 6))
	// out(r1.get(2, 3))
	// out(r1.get(3, 5))

	if cnt != -1 {
		out(ans)
	}
	out(cnt)
	// out(r0.get(0, 2))
	// out(r0.get(5, 6))
	// out(r1.get(1, 3))
	// out(r1.get(3, 4))
}

// 0 123 456
// a bcb ac  l....r
// 01 234 5
// ca bcb a  r..
// 97 a
// 98 b
// 99 c
