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

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	x := make([]int, M)
	y := make([]int, M)
	for i := 0; i < M; i++ {
		x[i], y[i] = getInt()-1, getInt()-1
	}

	a := make([]bool, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		b[i] = 1
	}
	a[0] = true
	for i := 0; i < M; i++ {
		f := x[i]
		t := y[i]
		if b[f] == 1 {
			a[t] = a[t] || a[f]
			a[f] = false
		} else {
			a[t] = a[t] || a[f]
		}
		b[f]--
		b[t]++
		// out(a)
	}
	// out(a)
	ans := 0
	for i := 0; i < N; i++ {
		if a[i] {
			ans++
		}
	}
	out(ans)
}
