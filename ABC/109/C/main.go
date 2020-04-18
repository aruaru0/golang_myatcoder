package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	N, X := getInt(), getInt()
	x := make([]int, N+1)
	for i := 0; i < N; i++ {
		x[i] = getInt()
	}
	x[N] = X
	sort.Ints(x)

	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = x[i+1] - x[i]
	}
	ans := d[0]
	for i := 1; i < N; i++ {
		ans = GCD(ans, d[i])
	}
	out(ans)
}
