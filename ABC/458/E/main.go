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

// iFracの計算でmpowを一度しか呼び出さない

const mod = 998244353
const size = int(3e6 + 100)

func mpow(p, n int) int {
	ret := 1
	x := p
	for n != 0 {
		if n%2 == 1 {
			ret *= x
			ret %= mod
		}
		n /= 2
		x = x * x % mod
	}
	return ret
}

var frac [size]int
var ifrac [size]int

func initFrac() {
	frac[0] = 1
	for i := 1; i < size; i++ {
		frac[i] = frac[i-1] * i % mod
	}
	ifrac[size-1] = mpow(frac[size-1], mod-2)
	for i := size - 2; i >= 0; i-- {
		ifrac[i] = ifrac[i+1] * (i + 1) % mod
	}
}

func nCk(n, k int) int {
	if n < k || k < 0 {
		return 0
	}
	return frac[n] * ifrac[k] % mod * ifrac[n-k] % mod
}

func nPk(n, k int) int {
	if k < 0 || n < k {
		return 0
	}
	return frac[n] * ifrac[n-k] % mod
}

func nHk(n, k int) int {
	if n == 0 && k == 0 {
		return 1
	}
	return nCk(n+k-1, k)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	x1, x2, x3 := getI(), getI(), getI()

	initFrac()

	ans := 0
	limit := min(x1-1, x2)
	for i := 0; i <= limit; i++ {
		// c1 := nCk(x1-1, i)
		c1 := nHk(i+1, x1-i-1)
		c2 := nCk(x2+1, i+1)
		// c3 := nCk(x2+x3-i-1, x2-i-1)
		c3 := nHk(x2-i, x3)
		ans += (c1 * c2 % mod * c3) % mod
		ans %= mod
	}
	out(ans)
}
