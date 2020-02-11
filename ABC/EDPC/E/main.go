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
		for j := 0; j <= 100000; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[0][0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j <= 100000; j++ {
			if j-v[i] < 0 {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = min(dp[i][j], dp[i][j-v[i]]+w[i])
			}
		}
	}

	max := 0
	for i := 0; i <= 100000; i++ {
		if dp[N][i] <= W && i > max {
			max = i
		}
	}

	out(max)
}
