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

	N, A, B := getInt(), getInt(), getInt()

	if A > B {
		out(0)
		return
	}

	if N == 1 {
		if A == B {
			out(1)
			return
		}
		out(0)
		return
	}
	if N == 2 {
		out(1)
		return
	}

	n := N - 2
	x := A*(n+1) + B
	y := B*(n+1) + A
	out(y - x + 1)
}
