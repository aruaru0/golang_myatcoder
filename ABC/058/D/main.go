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

const MOD = 1000000007

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	sc.Split(bufio.ScanWords)

	n, m := getInt(), getInt()
	x := make([]int, n)
	y := make([]int, m)

	for i := 0; i < n; i++ {
		x[i] = getInt()
	}
	for i := 0; i < m; i++ {
		y[i] = getInt()
	}

	l := 0
	r := n - 1
	t := n - 1
	sx := 0
	for t > 0 {
		sx = (sx + (asub(x[r], x[l])*t)%MOD) % MOD
		l++
		r--
		t -= 2
	}

	l = 0
	r = m - 1
	t = m - 1
	sy := 0
	for t > 0 {
		sy = (sy + (asub(y[r], y[l])*t)%MOD) % MOD
		l++
		r--
		t -= 2
	}
	out((sx * sy) % MOD)
}
