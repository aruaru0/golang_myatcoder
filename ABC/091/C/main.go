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

	// solve(N, M)

	if M == 1 && N == 1 {
		out(1)
		return
	}
	if M == 2 && N == 2 {
		out(0)
		return
	}
	if M == 1 {
		out(N - 2)
		return
	}
	if N == 1 {
		out(M - 2)
		return
	}

	N = max(0, N-2)
	M = max(0, M-2)
	out(N * M)

}
