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

func check(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	T := getInt()

	for i := 0; i < T; i++ {
		n, m := getInt(), getInt()
		if m <= n {
			out(0)
			continue
		}
		if n < 200000 {
			x := 1 % m
			for i := 1; i <= n; i++ {
				x *= i
				x %= m
			}
			out(x)
			continue
		}
		if check(m) {
			out(0)
			continue
		}
		x := m - 1
		for i := m - 2; i >= n; i-- {
			x *= modinv(i+1, m)
			x %= m
		}
		out(x)
	}
}
