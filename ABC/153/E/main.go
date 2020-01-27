package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func out(x ...interface{}) {
	fmt.Println(x...)
}

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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	H := getInt()
	N := getInt()
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = getInt()
		B[i] = getInt()
	}

	//out(H, N, A, B)
	dp := make([]int, 10001)
	max := 2147483647
	for i := 0; i < 10001; i++ {
		dp[i] = max
	}

	dp[0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j <= H; j++ {
			pos := min(j+A[i], H)
			if dp[pos] > dp[j]+B[i] {
				dp[pos] = dp[j] + B[i]
			}
		}
	}

	fmt.Println(dp[H])
}
