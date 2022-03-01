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

// iFracの計算でmpowを一度しか呼び出さない

const mod = 998244353
const size = 2000100

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

	n, m := getI(), getI()
	n2 := 1 << n
	a := getInts(n)
	b := getInts(m)

	ba := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c := getI()
			if c != 0 {
				ba[j] |= 1 << i
			}
		}
	}

	bs := make([]int, n2)
	for i := 0; i < m; i++ {
		bs[ba[i]] += b[i]
	}
	// 高速ゼータ変換
	for i := 0; i < n; i++ {
		for j := 0; j < n2; j++ {
			if j>>i%2 == 1 {
				bs[j] += bs[j^1<<i]
			}
		}
	}

	as := make([]int, n2)
	for i := 0; i < n2; i++ {
		for j := 0; j < n; j++ {
			if i>>j%2 == 1 {
				as[i] += a[j]
			}
		}
	}

	const inf = int(1e15)
	x := inf
	for s := 0; s < n2; s++ {
		if bs[s] != 0 {
			now := as[s] - bs[s] + 1
			x = min(x, now)
		}
	}
	if x <= 0 {
		out("0 1")
		return
	}

	//-----
	cs := make([]int, n2)
	for s := 0; s < n2; s++ {
		if bs[s] != 0 {
			now := as[s] - bs[s] + 1
			if now == x {
				cs[s] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n2; j++ {
			if j>>i%2 == 1 {
				cs[j^1<<i] |= cs[j]
			}
		}
	}

	initFrac()
	d := make([]int, n2)
	for s := 0; s < n2; s++ {
		d[s] = nCk(as[s], x)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n2; j++ {
			if j>>i%2 == 1 {
				d[j] -= d[j^1<<i]
				d[j] %= mod
				if d[j] < 0 {
					d[j] += mod
				}
			}
		}
	}

	ans := 0
	for s := 0; s < n2; s++ {
		if cs[s] != 0 {
			ans += d[s]
			ans %= mod
		}
	}
	out(x, ans)
}
