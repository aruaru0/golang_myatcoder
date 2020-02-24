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
	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = getInt()
	}
	var dp [101][11001]int
	dp[0][0] = 1
	for i := 1; i <= N; i++ {
		for j := 0; j < 11000; j++ {
			if dp[i-1][j] == 1 {
				dp[i][j] = 1
			}
			if j >= s[i-1] && dp[i-1][j-s[i-1]] == 1 {
				dp[i][j] = 1
			}
		}
	}

	ans := 0
	for i := 11000; i >= 0; i-- {
		if i%10 != 0 && dp[N][i] == 1 {
			ans = i
			break
		}
	}
	out(ans)
}
