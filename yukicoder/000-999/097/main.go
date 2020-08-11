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

var xor128_x = uint32(123456789)
var xor128_y = uint32(362436069)
var xor128_z = uint32(521288629)
var xor128_w = uint32(88675123)

func xor128() uint32 {
	t := xor128_x ^ (xor128_x << 11)
	xor128_x = xor128_y
	xor128_y = xor128_z
	xor128_z = xor128_w
	xor128_w = xor128_w ^ (xor128_w >> 19) ^ (t ^ (t >> 8))
	return xor128_w
}

func generateA(N int) []int {
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = int(xor128() % 100003)
	}
	return A
}

const mod = 100003

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
	N, Q := getInt(), getInt()
	a := generateA(N)
	flg := make([]int, mod)
	for _, v := range a {
		flg[v]++
	}

	for i := 0; i < Q; i++ {
		q := getInt()
		if q == 0 {
			out(0)
			continue
		}
		if N <= 1000 {
			ans := 0
			for i := 0; i < N; i++ {
				ans = max(ans, (a[i]*q)%mod)
			}
			out(ans)
			continue
		}
		ans := mod - 1
		rev := modinv(q, mod)
		for ; ans >= 0; ans-- {
			if flg[ans*rev%mod] != 0 {
				break
			}
		}
		out(ans)
	}
}
