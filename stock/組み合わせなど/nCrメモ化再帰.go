//半端な区間のみ利用する場合にＴＥＬしないことがある
//有用性は低い　yukicoder 1102で利用
const mod = int(1e9 + 7)

type nr struct {
	n, r int
}

var cm map[nr]int

func nCr(n, r int) int {
	v, ok := cm[nr{n, r}]
	if ok {
		return v
	}
	if r == 0 || r == n {
		return 1
	}
	if r == 1 || r == n-1 {
		return n
	}
	v = nCr(n-1, r) * n % mod * modinv(n-r, mod) % mod
	cm[nr{n, r}] = v
	return v
}

// 逆元を使った割り算（MOD）
// mod. m での a の逆元 a^{-1} を計算する
func modinv(a, m int) int {
	b := m
	u := 1
	v := 0
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
