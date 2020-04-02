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

func ch(a byte, R, S, P int) (byte, int) {
	var ret byte
	var pt int
	switch a {
	case 'r':
		ret = 'p'
		pt = P
	case 's':
		ret = 'r'
		pt = R
	case 'p':
		ret = 's'
		pt = S
	}
	return ret, pt
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, K := getInt(), getInt()
	R, S, P := getInt(), getInt(), getInt()
	T := getString()

	a := make([]byte, N)

	ans := 0
	for i := 0; i < N; i++ {
		x, pt := ch(T[i], R, S, P)
		if i-K >= 0 && a[i-K] == x {
			x = 'x'
			pt = 0
		}
		a[i] = x
		ans += pt
	}
	out(ans)
}
