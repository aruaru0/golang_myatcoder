package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
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

const mod = 1000000007

func mpow(a, b int) int {
	if b == 0 {
		return 1
	} else if b%2 == 0 {
		tmp := mpow(a, b/2)
		return tmp * tmp % mod
	}
	return mpow(a, b-1) * a % mod
}

var fracMemo = []int{1, 1}

func mfrac(n int) int {
	if len(fracMemo) > n {
		return fracMemo[n]
	}
	if len(fracMemo) == 0 {
		fracMemo = append(fracMemo, 1)
	}
	for len(fracMemo) <= n {
		size := len(fracMemo)
		fracMemo = append(fracMemo, fracMemo[size-1]*size%mod)
	}
	return fracMemo[n]
}

var ifracMemo = []int{1, 1}

func mifrac(n int) int {
	if len(ifracMemo) > n {
		return ifracMemo[n]
	}
	if len(ifracMemo) == 0 {
		fracMemo = append(ifracMemo, 1)
	}
	for len(ifracMemo) <= n {
		size := len(ifracMemo)
		ifracMemo = append(ifracMemo, ifracMemo[size-1]*mpow(size, mod-2)%mod)
	}
	return ifracMemo[n]
}

func nCr(n, r int) int {
	if n == r {
		return 1
	}
	if n < r || r < 0 {
		return 0
	}
	ret := 1
	ret *= mfrac(n)
	ret %= mod
	ret *= mifrac(r)
	ret %= mod
	ret *= mifrac(n - r)
	ret %= mod
	return (ret)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type pair struct {
	f, s int
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getInt(), getInt()
	}

	zero := 0
	m := make(map[pair]pair)
	for i := 0; i < N; i++ {
		if a[i] == 0 && b[i] == 0 {
			zero++
			continue
		}
		g := gcd(abs(a[i]), abs(b[i]))
		a[i] /= g
		b[i] /= g
		if a[i] == 0 {
			x := m[pair{0, 1}]
			x.f++
			m[pair{0, 1}] = x
		} else if b[i] == 0 {
			x := m[pair{0, 1}]
			x.s++
			m[pair{0, 1}] = x
		} else if a[i] > 0 && b[i] > 0 {
			x := m[pair{a[i], b[i]}]
			x.f++
			m[pair{a[i], b[i]}] = x
		} else if a[i] < 0 && b[i] < 0 {
			a[i] = -a[i]
			b[i] = -b[i]
			x := m[pair{a[i], b[i]}]
			x.f++
			m[pair{a[i], b[i]}] = x
		} else if a[i] < 0 {
			a[i] = -a[i]
			b[i] = -b[i]
			x := m[pair{-b[i], a[i]}]
			x.s++
			m[pair{-b[i], a[i]}] = x
		} else {
			x := m[pair{-b[i], a[i]}]
			x.s++
			m[pair{-b[i], a[i]}] = x
		}
	}

	// out(zero)
	// out(m)
	// for i, v := range m {
	// 	out(i, v)
	// }
	ans := 1
	for _, v := range m {
		x := mpow(2, v.f)
		x += mpow(2, v.s)
		x %= mod
		x--
		if x < 0 {
			x += mod
		}
		x %= mod
		// out(x)
		ans *= x
		ans %= mod
	}
	ans += zero - 1
	if ans < 0 {
		ans += mod
	}
	ans %= mod
	out(ans)
}
