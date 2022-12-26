package main

import (
	"bufio"
	"fmt"
	"math"
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	h := make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = getInt()
	}

	dp := make([]int, N)
	dp[0] = 0
	for i := 1; i < N; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i < N; i++ {
		for j := 1; j <= K; j++ {
			if i-j >= 0 {
				c := asub(h[i], h[i-j])
				dp[i] = min(dp[i], dp[i-j]+c)
			}
		}
	}
	out(dp[N-1])
}
