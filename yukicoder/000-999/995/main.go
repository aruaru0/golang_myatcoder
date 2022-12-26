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
	mod int
}

func newModint(m int) *modint {
	var ret modint
	ret.mod = m
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

// A[][]のp乗を求める
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

var mod = newModint(int(1e9 + 7))

func mulS(a, b [][]int) [][]int {
	c := make([][]int, 2)
	c[0] = []int{
		mod.add(mod.mul(a[0][0], b[0][0]), mod.mul(a[0][1], b[1][0])), mod.add(mod.mul(a[0][0], b[0][1]), mod.mul(a[0][1], b[1][1]))}
	c[1] = []int{
		mod.add(mod.mul(a[1][0], b[0][0]), mod.mul(a[1][1], b[1][0])), mod.add(mod.mul(a[1][0], b[0][1]), mod.mul(a[1][1], b[1][1]))}
	return c
}

func mul(a [][]int, k int) [][]int {
	//	b := make([][]int, 2)
	ret := make([][]int, 2)
	for i := 0; i < 2; i++ {
		ret[i] = make([]int, 2)
	}
	ret[0][0] = 1
	ret[1][1] = 1
	// copy(b, a)
	for k > 0 {
		if k%2 == 1 {
			ret = mulS(ret, a)
		}
		a = mulS(a, a)
		k /= 2
	}
	// out(b)
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, K, p, q := getI(), getI(), getI(), getI(), getI()
	b := getInts(N)

	a := make([][]int, 2)
	a[0] = []int{mod.sub(1, mod.div(p, q)), mod.div(p, q)}
	a[1] = []int{mod.div(p, q), mod.sub(1, mod.div(p, q))}

	a = mod.powModMatrix(a, K)
	ans := 0
	for i := 0; i < N; i++ {
		if i < M {
			e := mod.mul(b[i], a[0][0])
			ans = mod.add(ans, e)
		} else {
			e := mod.mul(b[i], a[0][1])
			ans = mod.add(ans, e)
		}
	}
	out(ans)
}
