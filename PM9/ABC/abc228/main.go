//----------------------------------------
// modint
//----------------------------------------
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
