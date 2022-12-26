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

type nr struct {
	n, r int
}

var cm map[nr]int

func nCr(n, r int) int {
	v, ok := cm[nr{n, r}]
	if ok {
		return v
	}
	if r == 0 || r == n {
		return 1
	}
	if r == 1 || r == n-1 {
		return n
	}
	v = nCr(n-1, r) * n % mod * modinv(n-r, mod) % mod
	cm[nr{n, r}] = v
	return v
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	cm = make(map[nr]int)

	N, K := getInt(), getInt()
	a := getInts(N)

	ans := 0
	for i := 1; i <= N; i++ {
		x := nCr(i+K-1, K)
		x *= nCr(N-i+K, K)
		x %= mod
		x *= a[i-1]
		x %= mod
		ans += x
		ans %= mod
	}
	out(ans)
}
