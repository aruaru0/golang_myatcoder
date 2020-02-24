package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

// MOD :
var MOD = 1000000007

// Fac :
var Fac [200001]int

// IFac :
var IFac [200001]int

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
	for i := 0; i < 200000; i++ {
		Fac[i+1] = Fac[i] * (i + 1) % MOD            // n!(mod M)
		IFac[i+1] = IFac[i] * mpow(i+1, MOD-2) % MOD // k!^{M-2} (mod M) ←累乗にmpowを採用
	}
}

func comb2(a, b int) int {
	if (a == 0) && (b == 0) {
		return 1
	}
	if (a < b) || (a < 0) {
		return 0
	}

	x := 1
	for i := a; i > a-b; i-- {
		x = (x * i) % MOD
	}
	tmp := x * IFac[b] % MOD
	return tmp
}

func main() {
	sc.Split(bufio.ScanWords)
	initPow()

	N, a, b := getInt(), getInt(), getInt()

	ans := mpow(2, N) - 1

	x := comb2(N, a)
	y := comb2(N, b)

	ans = (ans - x) % MOD
	if ans < 0 {
		ans += MOD
	}
	ans = (ans - y) % MOD
	if ans < 0 {
		ans += MOD
	}

	out(ans)
}
