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

func fib(n int) int {
	a, b, c, d := 1, 1, 1, 0
	e, f, g, h := 1, 0, 1, 0
	n = n - 1
	for n > 0 {
		if n%2 == 1 {
			i := a*e%mod + b*g%mod
			j := a*f%mod + b*f%mod
			k := c*e%mod + d*g%mod
			l := c*f%mod + d*h%mod
			e, f, g, h = i%mod, j%mod, k%mod, l%mod
		}
		i := a*a%mod + b*c%mod
		j := a*b%mod + b*d%mod
		k := c*a%mod + d*c%mod
		l := c*b%mod + d*d%mod
		a, b, c, d = i%mod, j%mod, k%mod, l%mod
		n /= 2
	}
	return e + f
}

// 乗数計算（MOD)
func mpow(x, n int) int {
	if x == 0 {
		return 0
	}
	ans := 1
	for n != 0 {
		if n&1 == 1 {
			ans = ans * x % mod
		}
		x = x * x % mod
		n = n >> 1
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	ans := 1
	for i := 0; i < N; i++ {
		c, d := getInt(), getString()
		x := int(d[0] - '0')
		for j := 1; j < len(d); j++ {
			x = x*10 + int(d[j]-'0')
			x %= (mod - 1)
		}
		ans *= mpow(fib(c+1), x)
		// out(fib(c+1), x, mpow(fib(c+1), x))
		ans %= mod
	}
	out(ans)
}
