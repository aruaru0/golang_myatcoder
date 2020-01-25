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

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	A := getInt()
	x := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = getInt()
	}

	var dp [51][51][3000]int // 面倒なので多めに確保

	dp[0][0][0] = 1
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < 2500; k++ {
				dp[i+1][j][k] += dp[i][j][k]
				dp[i+1][j+1][k+x[i]] += dp[i][j][k]
			}
		}
	}
	ans := 0
	for i := 1; i <= N; i++ {
		ans += dp[N][i][A*i]
	}

	out(ans)
}
