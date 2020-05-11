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
	k, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return k
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
	idx := sort.Search(len(a), func(k int) bool {
		return a[k] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(k int) bool {
		return a[k] > x
	})
	return idx
}

const mod = 998244353

func mpow(n, p int) int {
	ret := 1
	x := n
	for p > 0 {
		if p%2 == 1 {
			ret *= x
			ret %= mod
		}
		x *= x
		x %= mod
		p /= 2
	}
	return ret
}

var fact [maxval]int
var ifact [maxval]int

func nCk(n, k int) int {
	if n < k || k < 0 {
		return 0
	}
	ret := (fact[n] * ifact[k]) % mod
	ret *= ifact[n-k]
	ret %= mod
	return ret
}

const maxval = 220000

func main() {
	sc.Split(bufio.ScanWords)

	fact[0] = 1
	for k := 1; k < maxval; k++ {
		fact[k] = (k * fact[k-1]) % mod
	}
	ifact[maxval-1] = mpow(fact[maxval-1], mod-2)
	for k := maxval - 2; k >= 0; k-- {
		ifact[k] = ((k + 1) * ifact[k+1]) % mod
	}

	N, M, K := getInt(), getInt(), getInt()

	ans := 0
	for k := 0; k <= K; k++ {
		if N < k {
			break
		}
		x := M * mpow(M-1, N-1-k)
		x %= mod
		x *= nCk(N-1, k)
		x %= mod
		ans += x
		ans %= mod
	}

	out(ans)
}
