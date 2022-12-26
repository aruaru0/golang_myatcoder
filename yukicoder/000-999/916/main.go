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

// Writers解答を写経：数え上げでこの手のは不得意
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	d, l, r, k := getI(), getI(), getI(), getI()
	mod := newModint(1e9 + 7)

	pow2 := make([]int, 21)
	sum := make([]int, 21)
	fac := make([]int, 1<<20+1)
	pow2[0], pow2[1] = 1, 1
	for i := 2; i < 21; i++ {
		pow2[i] = pow2[i-1] * 2
	}
	for i := 1; i < 21; i++ {
		sum[i] = sum[i-1] + pow2[i]
	}
	fac[0] = 1
	for i := 1; i < 1<<20+1; i++ {
		fac[i] = mod.mul(fac[i-1], i)
	}

	//-----
	l = lowerBound(sum, l)
	r = lowerBound(sum, r)

	lca := -1
	//lcaが存在するならそのlcaの深さを計算
	if (l+r-k) > 1 && (l+r-k)%2 == 0 {
		lca = (l + r - k) / 2
	}

	//lcaが条件を満たしていない場合コーナー
	if lca == -1 || lca > l || lca > r {
		out(0)
		return
	}

	ans := 1
	//上の段からl,r以外の頂点の順列の数え上げ
	for i := 1; i <= d; i++ {
		cnt := pow2[i]
		if i == l {
			cnt--
		}
		if i == r {
			cnt--
		}
		ans = mod.mul(ans, fac[cnt])
	}

	//lcaとなるような頂点の位置についての数え上げ
	ans = mod.mul(ans, pow2[lca])
	//lcaを決め打ちした時のlの位置についての数え上げ
	ans = mod.mul(ans, pow2[l-lca])
	//lcaを決め打ちした時のrの位置についての数え上げ
	ans = mod.mul(ans, pow2[r-lca])
	//lcaについて、子は左右について2通りのパターンがあるため
	ans = mod.mul(ans, 2)

	out(ans)
}
