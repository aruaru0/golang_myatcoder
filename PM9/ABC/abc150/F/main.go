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

func f(a, b string) []int {
	// out(a, b)
	rb := newRollingHash(b)
	hb := rb.get(0, len(b))
	a = a + a
	ra := newRollingHash(a)
	ret := make([]int, 0)
	for i := 0; i < len(b); i++ {
		ha := ra.get(i, i+len(b))
		if hb == ha {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	b := getInts(N)

	ok := make([]int, N)
	ans := make([]int, N)
	cnt := 0
	for i := 0; i < 30; i++ {
		bita := make([]byte, 0)
		bitarev := make([]byte, 0)
		bitb := make([]byte, 0)
		zero := true
		for j := 0; j < N; j++ {
			x := (a[j] >> i) % 2
			y := (b[j] >> i) % 2
			if x != 0 || y != 0 {
				zero = false
			}
			bita = append(bita, byte(x+'0'))
			bitarev = append(bitarev, byte(x^1+'0'))
			bitb = append(bitb, byte(y+'0'))
		}
		if zero {
			continue
		}
		cnt++
		// out(i, "bit")
		z := f(string(bita), string(bitb))
		o := f(string(bitarev), string(bitb))
		for _, e := range z {
			ok[e]++
		}
		for _, e := range o {
			ok[e]++
			ans[e] |= 1 << i
		}
	}
	// out(ok, cnt)
	// out(ans)
	for i := 0; i < N; i++ {
		if ok[i] == cnt {
			out(i, ans[i])
		}
	}
}
