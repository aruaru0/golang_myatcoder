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

const mod = int(1e9 + 7)
const size = 100000

var frac [size]int
var ifrac [size]int

func mpow(x, p int) int {
	ret := 1
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

func initfrac() {
	frac[0], frac[1] = 1, 1
	for i := 2; i < size; i++ {
		frac[i] = frac[i-1] * i % mod
	}
	ifrac[size-1] = mpow(frac[size-1], mod-2)
	for i := size - 2; i >= 0; i-- {
		ifrac[i] = ifrac[i+1] * (i + 1) % mod
	}
}

func nCk(n, k int) int {
	if n < k || n == 0 {
		return 0
	}
	return frac[n] * ifrac[k] % mod * ifrac[n-k] % mod
}

const tblsize = 1000

// const mod = 1000000007

var sTable [tblsize][tblsize]int

func initSTable() {
	sTable[0][0] = 0
	sTable[1][0] = 0
	sTable[1][1] = 1
	for n := 2; n < tblsize; n++ {
		sTable[n][0] = 0
		for k := 1; k <= n; k++ {
			sTable[n][k] = sTable[n-1][k-1] + (k*sTable[n-1][k])%mod
			sTable[n][k] %= mod
		}
	}
}

func nSk(n, k int) int {
	if n >= tblsize || k >= tblsize {
		panic("nSk size overflow")
	}
	return sTable[n][k]
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()

	initfrac()
	initSTable()

	ans := 0
	for x := 1; x <= N; x++ {
		for y := 1; y < x+1; y++ {
			t := nCk(N, x) * nSk(x, y)
			t %= mod
			t *= mpow((y*(y-1))%mod, N-x)
			t %= mod
			// out("N:", N, " x:", x, "y:", y, N-x, t)
			ans += t
			ans %= mod
		}
	}
	out(ans)
}
