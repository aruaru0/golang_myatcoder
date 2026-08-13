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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	n            int
}

// newDsu :
func newDsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
	}
	return &d
}

// Merge :
func (d DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

// Same :
func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader :
func (d DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

// Size :
func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

// Groups : original implement
func (d DSU) Groups() [][]int {
	m := make(map[int][]int)
	for i := 0; i < d.n; i++ {
		x := d.Leader(i)
		if x < 0 {
			m[i] = append(m[i], i)
		} else {
			m[x] = append(m[x], i)
		}
	}
	ret := make([][]int, len(m))
	idx := 0
	for _, e := range m {
		ret[idx] = make([]int, len(e))
		copy(ret[idx], e)
		idx++
	}
	return ret
}

// ----------------------------------------
// modint
// ----------------------------------------
type modint struct {
	mod       int
	factMemo  []int
	ifactMemo []int
}

func newModint(m int) *modint {
	var ret modint
	ret.mod = m
	ret.factMemo = []int{1, 1}
	ret.ifactMemo = []int{1, 1}
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
	a %= m.mod
	b %= m.mod
	ret := a * b % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) div(a, b int) int {
	a %= m.mod
	ret := a * m.modinv(b)
	ret %= m.mod
	return ret
}

func (m *modint) pow(p, n int) int {
	ret := 1
	x := p % m.mod
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

// -----------------------------------------------
// 行列累乗
// 　A[][]のp乗を求める
// -----------------------------------------------
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

// ---------------------------------------------------
// nCk 計算関連：　TELすることがあるかも
//
//	※pow(x, p-2)を何度も取るので
//
// 厳しそうな場合は、ここを削除して高速なのを使う
// ---------------------------------------------------
func (m *modint) mfact(n int) int {
	if len(m.factMemo) > n {
		return m.factMemo[n]
	}
	if len(m.factMemo) == 0 {
		m.factMemo = append(m.factMemo, 1)
	}
	for len(m.factMemo) <= n {
		size := len(m.factMemo)
		m.factMemo = append(m.factMemo, m.factMemo[size-1]*size%m.mod)
	}
	return m.factMemo[n]
}

func (m *modint) mifact(n int) int {
	if len(m.ifactMemo) > n {
		return m.ifactMemo[n]
	}
	if len(m.ifactMemo) == 0 {
		m.factMemo = append(m.ifactMemo, 1)
	}
	for len(m.ifactMemo) <= n {
		size := len(m.ifactMemo)
		m.ifactMemo = append(m.ifactMemo, m.ifactMemo[size-1]*m.pow(size, m.mod-2)%m.mod)
	}
	return m.ifactMemo[n]
}

func (m *modint) nCr(n, r int) int {
	if n == r {
		return 1
	}
	if n < r || r < 0 {
		return 0
	}
	ret := 1
	ret *= m.mfact(n)
	ret %= m.mod
	ret *= m.mifact(r)
	ret %= m.mod
	ret *= m.mifact(n - r)
	ret %= m.mod
	return (ret)
}

// iFracの計算でmpowを一度しか呼び出さない

const mod = 998244353
const size = 2000100

func mpow(p, n int) int {
	ret := 1
	x := p
	for n != 0 {
		if n%2 == 1 {
			ret *= x
			ret %= mod
		}
		n /= 2
		x = x * x % mod
	}
	return ret
}

var frac [size]int
var ifrac [size]int

func initFrac() {
	frac[0] = 1
	for i := 1; i < size; i++ {
		frac[i] = frac[i-1] * i % mod
	}
	ifrac[size-1] = mpow(frac[size-1], mod-2)
	for i := size - 2; i >= 0; i-- {
		ifrac[i] = ifrac[i+1] * (i + 1) % mod
	}
}

func nCk(n, k int) int {
	if n < k || k < 0 {
		return 0
	}
	return frac[n] * ifrac[k] % mod * ifrac[n-k] % mod
}

func nPk(n, k int) int {
	if k < 0 || n < k {
		return 0
	}
	return frac[n] * ifrac[n-k] % mod
}

func nHk(n, k int) int {
	if n == 0 && k == 0 {
		return 1
	}
	return nCk(n+k-1, k)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	s := getS()
	a := make([]int, M)
	b := make([]int, M)

	uf := newDsu(N)
	for i := 0; i < M; i++ {
		a[i], b[i] = getI()-1, getI()-1
		uf.Merge(a[i], b[i])
	}

	mint := newModint(998244353)
	initFrac()
	ans := 1
	single := true
	for _, e := range uf.Groups() {
		m := make(map[byte]int)
		for _, idx := range e {
			m[s[idx]]++
		}
		cnt := frac[len(e)]
		for v := range m {
			cnt = mint.mul(cnt, ifrac[m[v]])
			if m[v] > 1 {
				single = false
			}
			// out(m[v])
		}
		// out(m, e, cnt)
		ans = mint.mul(ans, cnt)
	}
	if single {
		ans = mint.div(ans, 2)
	}
	out(ans)
}
