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

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	x := make([]int, N)
	x[0] = a[0]
	for i := 1; i < N; i++ {
		x[i] = GCD(x[i-1], a[i])
	}

	y := make([]int, N)
	y[N-1] = a[N-1]
	for i := N - 2; i >= 0; i-- {
		y[i] = GCD(y[i+1], a[i])
	}
	// out(x)
	// out(y)

	ans := max(x[N-2], y[1])
	for i := 1; i < N-1; i++ {
		g := GCD(x[i-1], y[i+1])
		ans = max(ans, g)
	}
	out(ans)
}
