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

func choice(x, y byte) byte {
	var ret byte
	// out("chise", string(x), string(y))
	if x == 'p' {
		switch y {
		case 'r':
			ret = 's'
		case 's':
			ret = 'r'
		case 'p':
			ret = 'r'
		}
	}
	if x == 'r' {
		switch y {
		case 'r':
			ret = 's'
		case 's':
			ret = 'p'
		case 'p':
			ret = 's'
		}
	}
	if x == 's' {
		switch y {
		case 'r':
			ret = 'p'
		case 's':
			ret = 'r'
		case 'p':
			ret = 'r'
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, K := getInt(), getInt()
	R, S, P := getInt(), getInt(), getInt()
	T := getString()

	a := make([]byte, N)

	ans := 0
	for i := 0; i < K; i++ {
		x, pt := ch(T[i], R, S, P)
		a[i] = x
		ans += pt
	}
	for i := K; i < N; i++ {
		x, pt := ch(T[i], R, S, P)
		if a[i-K] != x {
			a[i] = x
			ans += pt
		} else {
			x := T[i]
			if i+K < N {
				x, _ = ch(T[i+K], R, S, P)
			}
			y := choice(x, T[i])
			a[i] = y
		}
	}

	out(ans)
}
