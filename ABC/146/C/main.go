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

func digit(n int) int {
	ret := 0
	for n > 0 {
		ret++
		n /= 10
	}
	return ret
}

const inf = 1000000000

func solve(A, B, X int) int {
	l := 0
	r := inf + 1

	for l+1 != r {
		m := (l + r) / 2
		x := m*A + digit(m)*B
		// out(x, X, l, r)
		if X >= x {
			l = m
		} else {
			r = m
		}
	}
	return l
}

func main() {
	sc.Split(bufio.ScanWords)

	A, B, X := getInt(), getInt(), getInt()

	ans := solve(A, B, X)

	out(ans)
}
