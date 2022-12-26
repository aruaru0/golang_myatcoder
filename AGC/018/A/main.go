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

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	a := make([]int, N)
	m := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		m = max(m, a[i])
	}

	g := a[0]
	for i := 1; i < N; i++ {
		g = GCD(g, a[i])
	}

	if K <= m && K%g == 0 {
		out("POSSIBLE")
	} else {
		out("IMPOSSIBLE")
	}

}
