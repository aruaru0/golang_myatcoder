package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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
func prime(R int) []int {
	// Step 1: p^k (k>=2) の形の数を数える
	// p^k <= R (k>=2) より p <= sqrt(R)。
	sqrtR := int(math.Sqrt(float64(R)))

	// エラトステネスの篩で sqrt(R) までの素数を列挙
	is_prime := make([]bool, sqrtR+1)
	for i := range is_prime {
		is_prime[i] = true
	}
	is_prime[0], is_prime[1] = false, false
	for p := 2; p*p <= sqrtR; p++ {
		if is_prime[p] {
			for i := p * p; i <= sqrtR; i += p {
				is_prime[i] = false
			}
		}
	}

	primes := make([]int, 0)
	for p := 2; p <= sqrtR; p++ {
		if is_prime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	L := getI()
	R := getI()

	// 区間に含まれる数が1つだけなので、種類は1つ。
	if L == R {
		out(1)
		return
	}

	// A_n の値が A_{n-1} から変わるのは、n が素数べき(p^k, p:素数, k>=1)のとき。
	// 求める種類数は、A_L の1種類に、(L, R] の範囲の素数べきの個数を加えたもの。

	count := 0

	primes := prime(R)

	// p^k (k>=2) で (L, R] に入るものを数える
	for _, p := range primes {
		pk := p * p
		for {
			if pk > R {
				break
			}
			if pk > L {
				count++
			}
			// オーバーフロー防止: pk * p > R は R/p < pk とほぼ同値
			if R/p < pk {
				break
			}
			pk *= p
		}
	}

	// Step 2: 素数 (p^1) を数える
	// 区間篩で [L+1, R] に含まれる素数を数える
	offset := L + 1
	size := R - offset + 1

	if size > 0 {
		is_composite_seg := make([]bool, size) // false: 素数候補, true: 合成数

		for _, p := range primes {
			// 篩の開始地点を計算: offset 以上の最初の p の倍数
			start_val := (offset + p - 1) / p * p
			// 最適化: p*p 未満の p の倍数は、p より小さい素因数を持つため、
			// それらの素数によって既に篩われている。よって p*p からでよい。
			if start_val < p*p {
				start_val = p * p
			}

			for i := start_val; i <= R; i += p {
				if i >= offset {
					is_composite_seg[i-offset] = true
				}
			}
		}

		prime_count := 0
		for i := 0; i < size; i++ {
			if !is_composite_seg[i] {
				// L+1+i は素数
				prime_count++
			}
		}
		count += prime_count
	}

	// A_Lの分(1)と、(L, R]の素数べきの数を足す
	out(1 + count)
}
