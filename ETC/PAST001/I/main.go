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

func bitYN(s string) int {
	ret := 0
	for _, v := range s {
		if v == 'Y' {
			ret = ret<<1 + 1
		} else {
			ret = ret << 1
		}
	}
	return ret
}

const inf = int(1e18 + 1)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, M := getInt(), getInt()
	s := make([]int, M)
	c := make([]int, M)
	for i := 0; i < M; i++ {
		a, b := getString(), getInt()
		s[i] = bitYN(a)
		c[i] = b
	}

	B := 1 << uint(N)
	dp := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		dp[i] = make([]int, B)
		for j := 0; j < B; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0
	for i := 1; i <= M; i++ {
		for j := 0; j < B; j++ {
			// 購入しない場合は、コピー
			dp[i][j] = min(dp[i][j], dp[i-1][j])
			mask := s[i-1] | j
			// 購入する場合
			dp[i][mask] = min(dp[i][mask], dp[i-1][j]+c[i-1])
		}
	}
	/*
		for i := 0; i <= M; i++ {
			out(dp[i])
		}
	*/
	if dp[M][B-1] == inf {
		out(-1)
	} else {
		out(dp[M][B-1])
	}
}
