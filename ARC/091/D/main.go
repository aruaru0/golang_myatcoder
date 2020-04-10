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

func solve2(N, K int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			fmt.Printf("%2d ", i%j)
		}
		out()
	}
}

func solve(N, K int) int {
	ans := 0
	for i := 1; i <= N; i++ {
		if i >= K {
			ans += N / i * (i - K)
			ans += max(0, N%i-K+1)
		}
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()

	if K == 0 {
		out(N * N)
		return
	}
	out(solve(N, K))

	//solve2(N, K)
}
