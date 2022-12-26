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

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	p := make([]float64, N)

	for i := 0; i < N; i++ {
		p[i] = getF()
	}

	var dp [3000][3000]float64

	dp[0][0] = 1
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if j != 0 {
				dp[i][j] = (1-p[i-1])*dp[i-1][j] + p[i-1]*dp[i-1][j-1]
			} else {
				dp[i][j] = (1 - p[i-1]) * dp[i-1][j]
			}
		}
	}

	ans := float64(0)
	for i := (N + 1) / 2; i <= N; i++ {
		ans += dp[N][i]
	}
	out(ans)

}
