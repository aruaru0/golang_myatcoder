package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// MOD :
var MOD = 1000000007

// Fac :
var Fac [300001]int

// IFac :
var IFac [300001]int

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
func out(x ...interface{}) {
	//	fmt.Println(x...)
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

func main() {
	initPow()
	sc.Split(bufio.ScanWords)

	H := getInt()
	W := getInt()
	A := getInt()
	B := getInt()

	H0 := H - A
	W0 := W
	W1 := W - B
	H1 := A

	out(H0, W0, H1, W1)

	idx := 0
	sum := 0
	for i := B + 1; i <= W; i++ {
		out(H0+i-2, "C", H0-1)
		out(W1+H1-idx-2, "C", H1-1)
		x := comb(H0+i-2, H0-1)
		y := comb(W1+H1-idx-2, H1-1)
		sum = (sum + (x * y % MOD)) % MOD
		idx++
	}
	fmt.Println(sum)
}
