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

// 逆元を使った割り算（MOD）
// mod. m での a の逆元 a^{-1} を計算する
func modinv(a, m int) int {
	b := m
	u := 1
	v := 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= m
	if u < 0 {
		u += m
	}
	return u
}

func nCr(n, r int) int {
	if n == r {
		return 1
	}
	if n < r || r < 0 {
		return 0
	}
	f0 := 1
	f1 := 1
	for i := 0; i < r; i++ {
		f0 *= (n - i)
		f0 %= mod
		f1 *= (r - i)
		f1 %= mod
	}

	f1 = modinv(f1, mod)

	ret := f0 * f1
	ret %= mod

	return (ret)
}

func main() {
	sc.Split(bufio.ScanWords)
	N, a, b := getInt(), getInt(), getInt()

	ans := mpow(2, N)

	ans -= nCr(N, a)
	if ans < 0 {
		ans += mod
	}
	ans -= nCr(N, b)
	if ans < 0 {
		ans += mod
	}

	out(ans - 1)
}
