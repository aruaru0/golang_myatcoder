package main

import (
	"bufio"
	"fmt"
	"os"
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

const mod = 2147483647

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

func comb(n, k int) int {
	r := 1
	for d := 1; d <= k; d++ {
		r *= n
		n--
		r /= d
	}
	return r
}

func main() {
	sc.Split(bufio.ScanWords)

	A := getInt()

	out(A+1, 2)

}
