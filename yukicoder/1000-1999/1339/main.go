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
// func (m *modint) modinv(a int) int {
// 	b := m.mod
// 	u := 1
// 	v := 0
// 	for b != 0 {
// 		t := a / b
// 		a -= t * b
// 		a, b = b, a
// 		u -= t * v
// 		u, v = v, u
// 	}
// 	u %= m.mod
// 	if u < 0 {
// 		u += m.mod
// 	}
// 	return u
// }

// 拡張オイラーの互除法で逆元を求める
func (mm *modint) modinv(a int) int {
	m := mm.mod
	b, u, v := m, 1, 0
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

func f(x, n int) int {
	ret := 1
	p := 10 % n
	for x > 0 {
		if x&1 > 0 {
			ret = (ret * p) % n
		}
		p = (p * p) % n
		x >>= 1
	}
	return ret % n
}

func solve() {
	N := getI()
	for N%2 == 0 {
		N /= 2
	}
	for N%5 == 0 {
		N /= 5
	}
	if N == 1 {
		out(1)
		return
	}
	x := PfsMap(N)
	m := N
	for e := range x {
		m = m * (e - 1) / e
	}
	if m == 1 {
		out(1)
		return
	}
	// out("m=", m, N)
	ans := m
	for i := 1; i*i <= m; i++ {
		if m%i == 0 {
			// out(i, m/i, f(i, N), f(m/i, N))
			if f(i, N) == 1 {
				ans = min(ans, i)
			}
			if f(m/i, N) == 1 {
				ans = min(ans, m/i)
			}
		}
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		solve()
	}
}