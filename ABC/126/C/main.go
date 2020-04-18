package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()

	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		x := i
		cnt := 0
		for x < K {
			x *= 2
			cnt++
		}
		a[i] = cnt
	}
	p := 0.0
	for i := 1; i <= N; i++ {
		p += math.Pow(0.5, float64(a[i]))
	}
	// out(a)
	// fmt.Printf("%13.13f\n", p)
	ans := p / float64(N)
	fmt.Printf("%13.13f\n", ans)
}
