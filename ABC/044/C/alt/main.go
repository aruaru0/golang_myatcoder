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
		x[i] = getInt() - A
	}

	var dp [51][5300]int
	offset := 2500

	dp[0][offset] = 1
	for i := 0; i < N; i++ {
		for j := 0; j < 5000; j++ {
			//			out(j, j+x[i])
			dp[i+1][j] += dp[i][j]
			if j+x[i] > 0 {
				dp[i+1][j+x[i]] += dp[i][j]
			}
		}
	}

	out(dp[N][offset] - 1)
}
