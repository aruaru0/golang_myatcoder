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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func powmod(a, p int) int {
	if p == 0 {
		return 1
	}
	if p%2 == 1 {
		return powmod(a, p-1) * a % mod
	}
	mid := powmod(a, p/2)
	return mid * mid % mod
}

func rev(a int) int {
	return powmod(a, mod-2)
}

//　解説をもとに書き起こしたが、きちんと理解できず。
// 包除原理はわかるんだけど、ピントこない・・・
func main() {
	sc.Split(bufio.ScanWords)
	K := getInt()
	C := make([]int, K)
	sum := 0
	for i := 0; i < K; i++ {
		C[i] = getInt()
		sum += C[i]
	}
	g := C[0]
	for i := 0; i < K; i++ {
		g = gcd(g, C[i])
	}
	ans := 0
	factor := make([]int, 0)
	num := make([]int, 0)

	for i := g; i >= 1; i-- {
		if g%i != 0 {
			continue
		}
		remain := sum/i - 1

		mul := 1
		div := 1
		for j := 0; j < K; j++ {
			end := 0
			if j == 0 {
				end = 1
			}
			for l := 0; l < C[j]/i-end; l++ {
				mul *= remain
				remain--
				mul %= mod
				div *= l + 1
				div %= mod
			}
		}
		mul *= rev(div)

		for j := 0; j < len(factor); j++ {
			if factor[j]%i != 0 {
				continue
			}
			mul -= num[j]
		}
		mul %= mod
		mul += mod
		mul %= mod
		num = append(num, mul)

		mul *= rev(C[0] / i)
		mul %= mod
		factor = append(factor, i)
		ans += mul
		ans %= mod
	}
	out(ans)
}
