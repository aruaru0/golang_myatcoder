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

func primitiveRoot(m int) int {
	if m == 2 {
		return 1
	}
	if m == 167772161 || m == 469762049 || m == 998244353 {
		return 3
	}
	if m == 754974721 {
		return 11
	}
	divs := make([]int, 20)
	divs[0] = 2
	cnt := 1
	x := (m - 1) / 2
	for x%2 == 0 {
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

func powMod(a, n, M int) int {
	if M == 1 {
		return 0
	}
	r := 1
	for n > 0 {
		if n&1 == 1 {
			r = r * a % M
		}
		a = a * a % M
		n >>= 1
	}
	return r
}

func invMod(a, M int) int {
	p, x, u := M, 1, 0
	for p != 0 {
		t := a / p
		a, p = p, a-t*p
		x, u = u, x-t*u
	}
	return x
}

func ceilPow2(n int) int {
	x := 0
	for 1<<x < n {
		x++
	}
	return x
}

func bsf(n int) int {
	x := 0
	for n&(1<<x) == 0 {
		x++
	}
	return x
}

func butterfly(a []int32, M int) {
	g := primitiveRoot(M)
	n := len(a)
	h := ceilPow2(n)

	se := make([]int, 30)
	es, ies := make([]int, 30), make([]int, 30)
	cnt2 := bsf(M - 1)
	e := powMod(g, (M-1)>>cnt2, M)
	ie := invMod(e, M)
	for i := cnt2; i >= 2; i-- {
		es[i-2] = e
		ies[i-2] = ie
		e = e * e % M
		ie = ie * ie % M
	}
	now := 1
	for i := 0; i <= cnt2-2; i++ {
		se[i] = es[i] * now % M
		now = now * ies[i] % M
	}
	mm := int32(M)
	for ph := 1; ph <= h; ph++ {
		w := 1 << (ph - 1)
		p := 1 << (h - ph)
		now := 1
		for s := 0; s < w; s++ {
			offset := s << (h - ph + 1)
			for i := 0; i < p; i++ {
				l := a[i+offset]
				r := int32(int(a[i+offset+p]) * now % M)
				a[i+offset] = (l + r) % mm
				a[i+offset+p] = (mm + l - r) % mm
			}
			now = now * se[bsf(^s)] % M
		}
	}
}

func butterflyInv(a []int32, M int) {
	g := primitiveRoot(M)
	n := len(a)
	h := ceilPow2(n)

	sie := make([]int32, 30)
	es, ies := make([]int32, 30), make([]int32, 30)
	cnt2 := bsf(M - 1)
	e := powMod(g, (M-1)>>cnt2, M)
	ie := invMod(e, M)
	for i := cnt2; i >= 2; i-- {
		es[i-2] = int32(e)
		ies[i-2] = int32(ie)
		e = e * e % M
		ie = ie * ie % M
	}
	mm := int32(M)
	now := 1
	for i := 0; i <= cnt2-2; i++ {
		sie[i] = int32(int(ies[i]) * now % M)
		now = now * int(es[i]) % M
	}

	for ph := h; ph >= 1; ph-- {
		w := 1 << (ph - 1)
		p := 1 << (h - ph)
		inow := 1
		for s := 0; s < w; s++ {
			offset := s << (h - ph + 1)
			for i := 0; i < p; i++ {
				l := int32(a[i+offset])
				r := int32(a[i+offset+p])
				a[i+offset] = (l + r) % mm
				a[i+offset+p] = int32(int(mm+l-r) * inow % M)
			}
			inow = inow * int(sie[bsf(^s)]) % M
		}
	}
}

func convolutionMod(a, b []int32, M int) []int32 {
	n1, n2 := len(a), len(b)
	n := n1 + n2 - 1
	if n1 == 0 || n2 == 0 {
		return []int32{}
	}

	z := 1 << ceilPow2(n)
	aa, bb := make([]int32, z), make([]int32, z)
	copy(aa, a)
	copy(bb, b)
	a, b = aa, bb

	butterfly(a, M)
	butterfly(b, M)
	for i := 0; i < z; i++ {
		a[i] = int32(int(a[i]) * int(b[i]) % M)
	}
	butterflyInv(a, M)
	a = a[:n]
	iz := invMod(z, M)
	for i := 0; i < n; i++ {
		a[i] = int32(int(a[i]) * iz % M)
		if a[i] < 0 {
			a[i] += int32(M)
		}
	}

	return a
}

const mod = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	S := []byte(getS())
	T := []byte(getS())
	N := len(S)
	M := len(T)
	// Tを反転（FFTで処理するため）
	for i := 0; i < len(T)/2; i++ {
		T[i], T[M-1-i] = T[M-1-i], T[i]
	}

	// 文字列を整数配列に変換する
	s0 := make([]int32, N)
	for i, e := range S {
		s0[i] = int32(e - '0')
	}
	t0 := make([]int32, N)
	for i, e := range T {
		t0[i] = int32(e - '0')
	}

	// 01を反転した配列を作成
	s1 := make([]int32, N)
	for i := 0; i < N; i++ {
		s1[i] = s0[i] ^ 1
	}
	t1 := make([]int32, M)
	for i := 0; i < M; i++ {
		t1[i] = t0[i] ^ 1
	}

	//  a b   ^  ~a ~b  ~a+b  a+~b
	//  0 0 : 0   1  1    0    0
	//  0 1 : 1   1  0    1    0
	//  1 0 : 1   0  1    0    1
	//  1 1 : 0   1  1    0    0
	// a xor b = ~a * b + a * ~b
	c0 := convolutionMod(s0, t1, mod)
	c1 := convolutionMod(s1, t0, mod)

	ans := int32(math.MaxInt32)
	for i := M - 1; i < N; i++ {
		ans = min32(ans, c0[i]+c1[i])
	}
	out(ans)
}

func min32(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}
