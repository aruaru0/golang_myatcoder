package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007
const size = 2100000

var fact [size + 1]int
var factInv [size + 1]int
var seqInv [size + 1]int

func createFactMod() {
	fact[0] = 1
	fact[1] = 1
	for i := 2; i < size; i++ {
		fact[i] = fact[i-1] * i % MOD
	}
}

func createInvMod() {
	seqInv[0] = 1
	seqInv[1] = 1
	for i := 2; i < size; i++ {
		seqInv[i] = (MOD - MOD/i) * seqInv[MOD%i] % MOD
	}
}

func createFactInvMod() {
	factInv[0] = 1
	factInv[1] = 1
	for i := 2; i < size; i++ {
		factInv[i] = factInv[i-1] * seqInv[i] % MOD
	}
}

func initTables() {
	createFactMod()
	createInvMod()
	createFactInvMod()
}

func comb(n, k int) int {
	return ((fact[n] * factInv[n-k]) % MOD * factInv[k]) % MOD
}

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

func main() {
	initTables()

	sc.Split(bufio.ScanWords)

	r1, c1, r2, c2 := getInt(), getInt(), getInt(), getInt()

	x00 := comb(r1+c1, c1)
	x10 := comb(r2+c1+1, c1)
	x01 := comb(r1+c2+1, c2+1)
	x11 := comb(r2+c2+2, c2+1)

	ans := (x11 + x00) % MOD
	ans = ans - x01
	if ans < 0 {
		ans += MOD
	}
	ans = ans - x10
	if ans < 0 {
		ans += MOD
	}
	out(ans)
}
