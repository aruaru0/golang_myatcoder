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
	N := getInt()
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i], c[i] = getInt(), getInt(), getInt()
	}

	var dp [100000][3]int
	dp[0][0] = a[0]
	dp[0][1] = b[0]
	dp[0][2] = c[0]
	for i := 1; i < N; i++ {
		dp[i][0] = max(dp[i-1][1]+a[i], dp[i-1][2]+a[i])
		dp[i][1] = max(dp[i-1][0]+b[i], dp[i-1][2]+b[i])
		dp[i][2] = max(dp[i-1][0]+c[i], dp[i-1][1]+c[i])
	}

	out(max(dp[N-1][0], max(dp[N-1][1], dp[N-1][2])))

}
