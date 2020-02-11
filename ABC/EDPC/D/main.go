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

func main() {
	sc.Split(bufio.ScanWords)
	N, W := getInt(), getInt()
	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		w[i], v[i] = getInt(), getInt()
	}

	var dp [101][100001]int
	for i := 0; i < N; i++ {
		for k := 0; k <= W; k++ {
			if k-w[i] < 0 {
				dp[i+1][k] = dp[i][k]
			} else {
				dp[i+1][k] = max(dp[i][k], dp[i][k-w[i]]+v[i])
			}
		}
	}

	out(dp[N][W])
}
