package main

import (
	"bufio"
	"fmt"
	"math"
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

// 約数を列挙
func divisor(n int) []int {
	ret := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if i*i != n {
				ret = append(ret, n/i)
			}
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	x := divisor(2 * N)

	ans := math.MaxInt64
	for _, e := range x {
		// k = 0 mod e
		// k = -1 mod 2*N/2
		r := []int{0, -1}
		m := []int{e, 2 * N / e}
		k := crt(r, m)
		if k[0] != 0 {
			ans = min(ans, k[0])
		}
	}
	out(ans)
}

func floorSum(n, m, a, b int) int {
	ret := 0
	if a >= m {
		ret += (n - 1) * n * (a / m) / 2
		a %= m
	}
	if b >= m {
		ret += n * (b / m)
		b %= m
	}
	ymx := (a*n + b) / m
	xmx := ymx*m - b
	if ymx == 0 {
		return ret
	}
	ret += (n - (xmx+a-1)/a) * ymx
	ret += floorSum(ymx, a, m, (a-xmx%a)%a)
	return ret
}
func crt(r, m []int) [2]int {
	n := len(r)
	r0, m0 := 0, 1
	for i := 0; i < n; i++ {
		r1 := safeMod(r[i], m[i])
		m1 := m[i]
		if m0 < m1 {
			r0, r1 = r1, r0
			m0, m1 = m1, m0
		}
		if m0%m1 == 0 {
			if r0%m1 != r1 {
				return [2]int{0, 0}
			}
			continue
		}
		tmp := invGcd(m0, m1)
		g, im := tmp[0], tmp[1]
		u1 := m1 / g
		if (r1-r0)%g != 0 {
			return [2]int{0, 0}
		}
		x := (r1 - r0) / g % u1 * im % u1
		r0 += x * m0
		m0 *= u1
		if r0 < 0 {
			r0 += m0
		}
	}
	return [2]int{r0, m0}
}
func powMod(x, n, m int) int {
	if m == 1 {
		return 0
	}
	r := 1
	y := x % m
	if y < 0 {
		y += m
	}
	for n != 0 {
		if (n & 1) == 1 {
			r = (r * y) % m
		}
		y = (y * y) % m
		n >>= 1
	}
	return r
}
func safeMod(x, d int) int {
	x %= d
	if x < 0 {
		x += d
	}
	return x
}
func invMod(x, m int) int {
	z := invGcd(x, m)
	return z[1]
}
func invGcd(a, b int) [2]int {
	a = a % b
	if a < 0 {
		a += b
	}
	s, t := b, a
	m0, m1 := 0, 1
	for t != 0 {
		u := s / t
		s -= t * u
		m0 -= m1 * u
		tmp := s
		s = t
		t = tmp
		tmp = m0
		m0 = m1
		m1 = tmp
	}
	if m0 < 0 {
		m0 += b / s
	}
	return [2]int{s, m0}
}
func primitiveRoot(m int) int {
	if m == 2 {
		return 1
	} else if m == 167772161 {
		return 3
	} else if m == 469762049 {
		return 3
	} else if m == 754974721 {
		return 11
	} else if m == 998244353 {
		return 3
	}
	divs := make([]int, 20)
	divs[0] = 2
	cnt := 1
	x := (m - 1) / 2
	for (x % 2) == 0 {
		x /= 2
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			divs[cnt] = i
			cnt++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		divs[cnt] = x
		cnt++
	}
	for g := 2; ; g++ {
		ok := true
		for i := 0; i < cnt; i++ {
			if powMod(g, (m-1)/divs[i], m) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
	}
}
