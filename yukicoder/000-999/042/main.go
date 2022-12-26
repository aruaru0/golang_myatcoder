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

const mod = 1e9 + 9
const maxn = 3100

var c = []int{1, 5, 10, 50, 100, 500}

func mpow(x, p int) int {
	if p == 0 {
		return 1
	}
	if p == 1 {
		return x
	}
	if p%2 == 1 {
		return (mpow(x, p-1) * x) % mod
	}
	ret := mpow(x, p/2)
	return ret * ret % mod
}

func modi(a, m int) int {
	return mpow(a, m-2)
}

var dp [maxn]int

func solve(M int) {
	if M < maxn {
		out(dp[M])
		return
	}
	ans := 0
	q := M % 500
	m := (M / 500) % mod
	for i := 0; i < 6; i++ {
		tmp := 1
		for j := 0; j < 6; j++ {
			if i == j {
				continue
			}
			tmp *= m - j
			tmp %= mod
			tmp *= modi(i-j, mod) % mod
			tmp %= mod
		}
		tmp *= dp[i*500+q]
		tmp %= mod
		ans += tmp
		ans %= mod
	}
	ans += mod
	ans %= mod
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)

	for i := 0; i < maxn; i++ {
		dp[i] = 1
	}
	for i := 1; i < 6; i++ {
		for j := c[i]; j < maxn; j++ {
			dp[j] += dp[j-c[i]]
			dp[j] %= mod
		}
	}

	T := getInt()
	for i := 0; i < T; i++ {
		M := getInt()
		solve(M)
	}

}
