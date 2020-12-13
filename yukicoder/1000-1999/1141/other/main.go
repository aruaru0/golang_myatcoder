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

//----------------------------------------
// modint
//----------------------------------------
type modint struct {
	mod       int
	fracMemo  []int
	ifracMemo []int
}

func newModint(m int) *modint {
	var ret modint
	ret.mod = m
	ret.fracMemo = []int{1, 1}
	ret.ifracMemo = []int{1, 1}
	return &ret
}

func (m *modint) add(a, b int) int {
	ret := (a + b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) sub(a, b int) int {
	ret := (a - b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) mul(a, b int) int {
	ret := a * b % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) div(a, b int) int {
	ret := a * m.modinv(b)
	ret %= m.mod
	return ret
}

func (m *modint) pow(p, n int) int {
	ret := 1
	x := p
	for n != 0 {
		if n%2 == 1 {
			ret *= x
			ret %= m.mod
		}
		n /= 2
		x = x * x % m.mod
	}
	return ret
}

// 逆元を使った割り算（MOD）
// mod. m での a の逆元 a^{-1} を計算する
func (m *modint) modinv(a int) int {
	b := m.mod
	u := 1
	v := 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= m.mod
	if u < 0 {
		u += m.mod
	}
	return u
}

//-----------------------------------------------
// 行列累乗
// 　A[][]のp乗を求める
//-----------------------------------------------
func (m *modint) powModMatrix(A [][]int, p int) [][]int {
	N := len(A)
	ret := make([][]int, N)
	for i := 0; i < N; i++ {
		ret[i] = make([]int, N)
		ret[i][i] = 1
	}

	for p > 0 {
		if p&1 == 1 {
			ret = m.mulMod(ret, A)
		}
		A = m.mulMod(A, A)
		p >>= 1
	}

	return ret
}

func (m *modint) mulMod(A, B [][]int) [][]int {
	H := len(A)
	W := len(B[0])
	K := len(A[0])
	C := make([][]int, W)
	for i := 0; i < W; i++ {
		C[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j]
				C[i][j] %= m.mod
			}
		}
	}

	return C
}

//---------------------------------------------------
// nCk 計算関連：　TELすることがあるかも
//                ※pow(x, p-2)を何度も取るので
// 厳しそうな場合は、ここを削除して高速なのを使う
//---------------------------------------------------
func (m *modint) mfrac(n int) int {
	if len(m.fracMemo) > n {
		return m.fracMemo[n]
	}
	if len(m.fracMemo) == 0 {
		m.fracMemo = append(m.fracMemo, 1)
	}
	for len(m.fracMemo) <= n {
		size := len(m.fracMemo)
		m.fracMemo = append(m.fracMemo, m.fracMemo[size-1]*size%m.mod)
	}
	return m.fracMemo[n]
}

func (m *modint) mifrac(n int) int {
	if len(m.ifracMemo) > n {
		return m.ifracMemo[n]
	}
	if len(m.ifracMemo) == 0 {
		m.fracMemo = append(m.ifracMemo, 1)
	}
	for len(m.ifracMemo) <= n {
		size := len(m.ifracMemo)
		m.ifracMemo = append(m.ifracMemo, m.ifracMemo[size-1]*m.pow(size, m.mod-2)%m.mod)
	}
	return m.ifracMemo[n]
}

func (m *modint) nCr(n, r int) int {
	if n == r {
		return 1
	}
	if n < r || r < 0 {
		return 0
	}
	ret := 1
	ret *= m.mfrac(n)
	ret %= m.mod
	ret *= m.mifrac(r)
	ret %= m.mod
	ret *= m.mifrac(n - r)
	ret %= m.mod
	return (ret)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = getInts(W)
	}
	mod := newModint(1e9 + 7)
	s0 := make([][]int, H)
	s1 := make([][]int, H)
	s2 := make([][]int, H)
	s3 := make([][]int, H)
	for i := 0; i < H; i++ {
		s0[i] = make([]int, W)
		s1[i] = make([]int, W)
		s2[i] = make([]int, W)
		s3[i] = make([]int, W)
	}

	// s0
	for i := 0; i < H; i++ {
		tot := 1
		for j := 0; j < W; j++ {
			tot = mod.mul(tot, a[i][j])
			s0[i][j] = tot
		}
	}
	for j := 0; j < W; j++ {
		tot := 1
		for i := 0; i < H; i++ {
			tot = mod.mul(tot, s0[i][j])
			s0[i][j] = tot
		}
	}
	// s1
	for i := 0; i < H; i++ {
		tot := 1
		for j := 0; j < W; j++ {
			tot = mod.mul(tot, a[i][j])
			s1[i][j] = tot
		}
	}
	for j := 0; j < W; j++ {
		tot := 1
		for i := H - 1; i >= 0; i-- {
			tot = mod.mul(tot, s1[i][j])
			s1[i][j] = tot
		}
	}
	// s2
	for i := 0; i < H; i++ {
		tot := 1
		for j := W - 1; j >= 0; j-- {
			tot = mod.mul(tot, a[i][j])
			s2[i][j] = tot
		}
	}
	for j := 0; j < W; j++ {
		tot := 1
		for i := 0; i < H; i++ {
			tot = mod.mul(tot, s2[i][j])
			s2[i][j] = tot
		}
	}
	// s3
	for i := 0; i < H; i++ {
		tot := 1
		for j := W - 1; j >= 0; j-- {
			tot = mod.mul(tot, a[i][j])
			s3[i][j] = tot
		}
	}
	for j := 0; j < W; j++ {
		tot := 1
		for i := H - 1; i >= 0; i-- {
			tot = mod.mul(tot, s3[i][j])
			s3[i][j] = tot
		}
	}

	Q := getI()
	for i := 0; i < Q; i++ {
		r, c := getI()-1, getI()-1
		x0 := 1
		if r != 0 && c != 0 {
			x0 = s0[r-1][c-1]
		}
		x1 := 1
		if r != H-1 && c != 0 {
			x1 = s1[r+1][c-1]
		}
		x2 := 1
		if r != 0 && c != W-1 {
			x2 = s2[r-1][c+1]
		}
		x3 := 1
		if r != H-1 && c != W-1 {
			x3 = s3[r+1][c+1]
		}
		ans := mod.mul(mod.mul(x0, x1), mod.mul(x2, x3))
		out(ans)
	}
}
