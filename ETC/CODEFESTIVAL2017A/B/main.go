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

func solve(n, m, k int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			x := (m-j)*i + (n-i)*j
			// out(i, j, x)
			if x == k {
				out("Yes")
				return
			}
		}
	}
	out("No")
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M, K := getInt(), getInt(), getInt()

	solve(N, M, K)
}
