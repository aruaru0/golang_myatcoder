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

const mod = 1000000007

var dp []int
P
// メモ化再帰バージョン
func rec(i, N int, m []int) int {
	if dp[i] != -1 {
		return dp[i]
	}
	if i == N {
		return 1
	}
	var r0, r1 int
	if i < N && m[i+1] != 1 {
		r0 = rec(i+1, N, m)
	}

	if i < N-1 && m[i+2] != 1 {
		r1 = rec(i+2, N, m)
	}

	dp[i] = (r0 + r1) % mod
	return dp[i]
}

// DPバージョン
func solve(N int, m []int) int {
	dp = make([]int, N+1)
	dp[0] = 1
	for i := 1; i <= N; i++ {
		if m[i-1] != 1 {
			dp[i] += dp[i-1]
			dp[i] %= mod
		}
		if i > 1 && m[i-2] != 1 {
			dp[i] += dp[i-2]
			dp[i] %= mod
		}
	}

	return dp[N]
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	m := make([]int, N+1)
	dp = make([]int, N+1)

	for i := 0; i <= N; i++ {
		dp[i] = -1
	}

	for i := 0; i < M; i++ {
		a := getInt()
		m[a] = 1
	}

	//out(rec(0, N, m))
	out(solve(N, m))
}
