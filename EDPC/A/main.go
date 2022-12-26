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

	N := getInt()
	h := make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = getInt()
	}
	dp := make([]int, N)
	dp[0] = 0
	for i := 1; i < N; i++ {
		dp[i] = math.MaxUint32
	}
	for i := 1; i < N; i++ {
		if i == 1 {
			c0 := asub(h[i-1], h[i])
			dp[i] = min(dp[i], dp[i-1]+c0)
		} else {
			c0 := dp[i-1] + asub(h[i-1], h[i])
			c1 := dp[i-2] + asub(h[i-2], h[i])
			dp[i] = min(dp[i], min(c0, c1))
		}
	}
	out(dp[N-1])
}

/*
func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	h := make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = getInt()
	}
	dp := make([]int, N)
	dp[0] = 0
	for i := 1; i < N; i++ {
		dp[i] = math.MaxUint32
	}
	for i := 0; i < N; i++ {
		if i+1 < N {
			c1 := asub(h[i+1], h[i])
			dp[i+1] = min(dp[i+1], dp[i]+c1)
		}
		if i+2 < N {
			c2 := asub(h[i+2], h[i])
			dp[i+2] = min(dp[i+2], dp[i]+c2)
		}
	}
	out(dp[N-1])
}
*/
