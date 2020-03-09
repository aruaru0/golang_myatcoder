package main

import "fmt"

// MOD :
var MOD = 1000000007

// Fac :
var Fac [300001]int

// IFac :
var IFac [300001]int

// 乗数計算（MOD)
func mpow(x, n int) int {
	ans := 1
	for n != 0 {
		if n&1 == 1 {
			ans = ans * x % MOD
		}
		x = x * x % MOD
		n = n >> 1
	}
	return ans
}

// コンビネーション計算
func comb(a, b int) int {
	if (a == 0) && (b == 0) {
		return 1
	}
	if (a < b) || (a < 0) {
		return 0
	}
	tmp := IFac[a-b] * IFac[b] % MOD
	return tmp * Fac[a] % MOD
}

func initPow() {
	Fac[0] = 1
	IFac[0] = 1
	for i := 0; i < 300000; i++ {
		Fac[i+1] = Fac[i] * (i + 1) % MOD            // n!(mod M)
		IFac[i+1] = IFac[i] * mpow(i+1, MOD-2) % MOD // k!^{M-2} (mod M) ←累乗にmpowを採用
	}
}

func main() {
	initPow()

	var N, K int

	fmt.Scanf("%d %d", &N, &K)

	ans := comb(N, K)

	fmt.Println(ans)
}
