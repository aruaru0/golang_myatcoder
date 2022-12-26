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
	K, A, B := getInt(), getInt(), getInt()

	if A+1 >= B {
		out(K + 1)
		return
	}

	k := K - (A - 1)
	if k < 0 {
		out(K + 1)
		return
	}
	// out(k, B, A, k/2, k%2)

	// out(A, (B-A)*(k/2), k%2)
	ans := A + (B-A)*(k/2) + k%2

	out(ans)
}
