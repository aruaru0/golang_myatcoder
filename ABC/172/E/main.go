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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

func mpow(x, y int) int {
	a := x
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret *= a
			ret %= mod
		}
		a = a * a
		a %= mod
		y /= 2
	}
	return ret
}

var fact [500100]int
var ifact [500100]int

func nCk(n, k int) int {
	return fact[n] * ifact[k] % mod * ifact[n-k] % mod
}

func nPk(n, k int) int {
	return fact[n] * ifact[n-k] % mod
}

func main() {
	sc.Split(bufio.ScanWords)

	fact[0] = 1
	for i := 1; i <= 500001; i++ {
		fact[i] = (fact[i-1] * i) % mod
	}
	ifact[500000] = mpow(fact[500000], mod-2)
	for i := 500000; i > 0; i-- {
		ifact[i-1] = (ifact[i] * i) % mod
	}

	N, M := getInt(), getInt()
	ans := 0
	for k := 0; k <= N; k++ {
		x := nCk(N, k) * nPk(M, k) % mod * nPk(M-k, N-k) % mod * nPk(M-k, N-k) % mod
		//out(x)
		if k%2 == 0 {
			ans += x
			ans %= mod
		} else {
			ans -= x
			ans %= mod
			if ans < 0 {
				ans += mod
			}
		}
	}
	out(ans)
}
