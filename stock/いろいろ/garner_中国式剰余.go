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

func invmod(a, m int) int {
	s := a % m
	t := m
	sx, sy, tx, ty := 1, 0, 0, 1
	for s%t != 0 {
		f := s / t
		u := s - t*f
		ux := sx - tx*f
		uy := sy - ty*f
		s = t
		sx = tx
		sy = ty
		t = u
		tx = ux
		ty = uy
	}
	if tx < 0 {
		tx += m
	}
	return tx
}

// a = x (mod m)
func garner(x []int, m []int, mod int) int {
	if len(x) == 0 {
		return 0
	}
	v := make([]int, len(x))
	v[0] = x[0]
	for i := 1; i < len(x); i++ {
		X := x[i]
		M := 1
		for j := 0; j < i; j++ {
			X -= v[j] * M
			X %= m[i]
			M *= m[j]
			M %= m[i]
		}
		if X < 0 {
			X += m[i]
		}
		v[i] = X * invmod(M, m[i]) % m[i]
	}
	ret := v[0]
	p := 1
	if mod == 0 {
		for i := 1; i < len(x); i++ {
			p *= m[i-1]
			ret += p * v[i]
		}
	} else {
		ret %= mod
		for i := 1; i < len(x); i++ {
			p *= m[i-1]
			p %= mod
			ret += p * v[i]
			ret %= mod
		}
	}
	return ret
}

type pair struct {
	first, second int
}

// 剰余(B)が互いに素でない場合の処理
func garnerHelper(A, B []int, mod int, zero bool) int {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		b := B[i]
		for j := 2; j*j <= b; j++ {
			if b%j == 0 {
				m[j]++
				for b%j == 0 {
					b /= j
				}
			}
		}
		if b > 1 {
			m[b]++
		}
	}
	primes := make([]int, 0)
	for i := range m {
		primes = append(primes, i)
	}
	sort.Ints(primes)
	a := make([]int, 0)
	b := make([]int, 0)
	flag := !zero
	ret := 1
	for _, p := range primes {
		ma := 1
		res := 0
		x := make([]pair, 0)
		for i := 0; i < len(A); i++ {
			if B[i]%p == 0 {
				t := 1
				for B[i]%p == 0 {
					B[i] /= p
					t *= p
				}
				x = append(x, pair{t, A[i] % t})
				if ma < t {
					ma = t
					res = x[len(x)-1].second
				}
				A[i] %= B[i]
			}
		}
		for _, q := range x {
			if res%q.first != q.second {
				return -1
			}
		}
		a = append(a, res)
		b = append(b, ma)
		flag = flag && (res == 0)
		if flag {
			ret *= ma
			if mod != 0 {
				ret %= mod
			}
		}
	}
	if !flag {
		ret = garner(a, b, mod)
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt(), getInt()
	}
	out(garnerHelper(x, y, 1e9+7, false))
}
