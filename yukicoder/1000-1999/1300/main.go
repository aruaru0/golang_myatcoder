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

type BIT struct {
	v []int
}

func newBIT(n int) *BIT {
	b := new(BIT)
	b.v = make([]int, n)
	return b
}
func (b BIT) sum(a int) int {
	ret := 0
	for i := a + 1; i > 0; i -= i & -i {
		ret += b.v[i-1]
		ret %= mod.mod
	}
	return ret
}
func (b BIT) rangeSum(x, y int) int {
	if y == 0 {
		return 0
	}
	y--
	if x == 0 {
		return b.sum(y)
	} else {
		return b.sum(y) - b.sum(x-1)
	}
}
func (b BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
		b.v[i-1] %= mod.mod
	}
}

type obj struct {
	a, i int
}

var mod = newModint(998244353)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	b := make([]obj, N)
	for i := 0; i < N; i++ {
		b[i] = obj{a[i], i + 1}
	}
	sort.Slice(b, func(i, j int) bool {
		if b[i].a == b[j].a {
			return b[i].i < b[j].i
		}
		return b[i].a < b[j].a
	})

	bit1 := newBIT(N + 1)
	bit1c := newBIT(N + 1)
	bit2 := newBIT(N + 1)
	bit2c := newBIT(N + 1)
	ans := 0
	for i := 0; i < N; i++ {
		o := b[i]
		v2 := mod.sub(bit2.sum(N), bit2.sum(o.i))
		c2 := mod.sub(bit2c.sum(N), bit2c.sum(o.i))
		val := mod.add(mod.mul(o.a, c2), v2)
		ans = mod.add(ans, val)

		v1 := mod.sub(bit1.sum(N), bit1.sum(o.i))
		c1 := mod.sub(bit1c.sum(N), bit1c.sum(o.i))
		p := mod.add(mod.mul(o.a, c1), v1)
		bit2.add(o.i, p)
		bit2c.add(o.i, c1)

		bit1.add(o.i, o.a)
		bit1c.add(o.i, 1)
		// fmt.Println("------------------")
		// view(bit2, N+1)
		// view(bit2c, N+1)
		// view(bit1, N+1)
		// view(bit1c, N+1)
	}

	out(ans)
}

func view(bit *BIT, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(bit.rangeSum(i, i+1), " ")
	}
	fmt.Println()
}
