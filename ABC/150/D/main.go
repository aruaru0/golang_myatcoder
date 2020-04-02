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

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM : find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a / GCD(a, b) * b

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func x2(n int) int {
	ret := 0
	for n%2 == 0 {
		ret++
		n /= 2
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	//
	for i := 0; i < N; i++ {
		a[i] /= 2
	}

	t := x2(a[0])
	// out(t)
	a[0] >>= uint(t)
	for i := 1; i < N; i++ {
		if x2(a[i]) != t {
			out(0)
			return
		}
		a[i] >>= uint(t)
	}

	M >>= uint(t)

	l := LCM(1, 1, a...)

	M /= l
	out((M + 1) / 2)
}
