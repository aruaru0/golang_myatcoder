package main

import (
	"bufio"
	"fmt"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	plus := make([]int, 0)
	minus := make([]int, 0)

	for i := 0; i < N; i++ {
		a := getI()
		if a >= 0 {
			plus = append(plus, a)
		} else {
			minus = append(minus, a)
		}
	}

	// 正がない場合で奇数選択の場合は、大きいものから順に選ぶ
	if len(plus) == 0 && K%2 == 1 {
		ans := 1
		sort.Slice(minus, func(i, j int) bool {
			return minus[i] > minus[j]
		})
		for i := 0; i < K; i++ {
			ans *= minus[i]
			ans %= mod
		}
		if ans < 0 {
			ans += mod
		}
		out(ans)
		return
	}

	sort.Slice(plus, func(i, j int) bool {
		return plus[i] > plus[j]
	})
	sort.Slice(minus, func(i, j int) bool {
		return minus[i] < minus[j]
	})

	// すべて正の場合は大きい順に選択
	if len(minus) == 0 {
		ans := 1
		for i := 0; i < K; i++ {
			ans *= plus[i]
			ans %= mod
		}
		out(ans)
		return
	}

	ans := 1
	ip, im := 0, 0

	// 奇数の場合は正から１つ選ぶ
	// 奇数で正がない場合は上でカバー済み
	if K%2 == 1 {
		ans = ans * plus[ip] % mod
		ip++
		K--
	}

	for K > 0 {
		x, y := -1, -1
		if ip+1 < len(plus) {
			x = plus[ip] * plus[ip+1]
		}
		if im+1 < len(minus) {
			y = minus[im] * minus[im+1]
		}
		if x == -1 && y == -1 {
			// １つずつあまりが発生した場合
			// 全部使い切る場合？
			x = plus[ip] * minus[im] % mod
			// 負になるのでmodを加えておく
			x += mod
			x %= mod
			ans = ans * x % mod
			im++
			ip++
		} else if x > y {
			ans = ans * (x % mod) % mod
			ip += 2
		} else {
			ans = ans * (y % mod) % mod
			im += 2
		}
		K -= 2
	}
	out(ans)
}
