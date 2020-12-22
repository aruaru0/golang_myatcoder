package main

import "fmt"

func out(x ...interface{}) {
	fmt.Println(x...)
}

func extgcd(a, b int) (int, int, int) { // --> g, x, y
	if b == 0 {
		return a, 1, 0
	}
	// ax + by = gcd(a,b)
	// bX + (a%b)Y = gcd(a,b)
	// a%b + a/b * b = a から a%b Y + a/b * b Y = a Y
	// 						  a%b Y= (a - a/b*b)Y
	// a%bに代入して、
	// bX + (a-a/b*b)Y = bX + aY - a/b*b Y = b(X - a/b*Y) - aY
	// aの部分とbの部分を取り出す
	// aY = ax, x = Y
	// (X - a/b*Y)b = by, y = X - a/b * Y
	g, X, Y := extgcd(b, a%b)
	x := Y
	y := X - a/b*Y
	return g, x, y
}

// mod mの1..n-1までの逆元を計算する　O(N)
func invMOD(n, m int) []int {
	// M/x*x + M%x  = M
	// M/x*x + M%x = 0 (mod M)
	// M/x*x = -M%x (mod M)
	// M/x = -M%x * inv(x) (mod M)
	// M/x * -inv(M%x) = inv(x) (mod M)
	ret := make([]int, n)
	ret[1] = 1
	for i := 2; i < n; i++ {
		ret[i] = m / i * (-ret[m%i])
		ret[i] %= m
		if ret[i] < 0 {
			ret[i] += m
		}
	}
	return ret
}

const mod = int(1e9 + 7)

func main() {
	inv := invMOD(10000, mod)
	for i := 1; i < 10000; i++ {
		v := inv[i] * i % mod
		if v == 1 {
			out("OK")
		} else {
			out("NG", i, inv[i])
			return
		}
	}
	// for i := 1; i < 100; i++ {
	// 	_, x, _ := extgcd(i, mod)
	// 	v := i * x % mod
	// 	if v < 0 {
	// 		v += mod
	// 	}
	// 	if v == 1 {
	// 		// out("OK")
	// 	} else {
	// 		out("NG", v)
	// 	}
	// }
}
