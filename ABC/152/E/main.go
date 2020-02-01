package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// PfsMap : 素因数分解し、マップを作成
func PfsMap(n int) map[int]int {
	pfs := make(map[int]int)
	for n%2 == 0 {
		pfs[2] = pfs[2] + 1
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i] = pfs[i] + 1
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = pfs[n] + 1
	}

	return pfs
}

func out(x ...interface{}) {
	//	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

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
	MOD := 1000000007
	N := nextInt()
	a := make(map[int]map[int]int)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		v := nextInt()
		A[i] = v
		a[i] = PfsMap(v)
		out(a[i])
	}

	// merge
	s := make(map[int]int)
	for i := 0; i < N; i++ {
		for n, m := range a[i] {
			if s[n] < m {
				s[n] = m
			}
		}
	}

	out(s)

	//
	sum := 1
	for n, m := range s {
		sum *= int(math.Pow(float64(n), float64(m))) % MOD
		sum = sum % MOD
	}

	out(sum)

	ans := 0
	for i := 0; i < N; i++ {
		ans += sum * modinv(A[i], MOD) % MOD
		ans = ans % 1000000007
	}

	fmt.Println(ans)
	/*
		xa := 12345678900000
		xb := 100000

		xa %= MOD
		fmt.Println(xa, xb, xa*modinv(xb, MOD)%MOD)
	*/
	/*
		sum := 0
		for i := 0; i < N; i++ {
			val := 1
			for n, m := range s {
				x := m - a[i][n]
				val *= int(math.Pow(float64(n), float64(x)))
			}
			sum += val
			sum = sum % 1000000007
			out(val)
		}
		fmt.Println(sum)
	*/
}
