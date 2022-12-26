package main

import (
	"bufio"
	"fmt"
	"os"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 10000)

	s, _, _ := r.ReadLine()
	t, _, _ := r.ReadLine()

	var dp [3001][3001]int

	n := len(s)
	m := len(t)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	//	out(dp[n][m])

	x := n
	y := m
	ans := ""
	for x > 0 && y > 0 {
		if dp[x][y] == dp[x-1][y] {
			x--
		} else if dp[x][y] == dp[x][y-1] {
			y--
		} else {
			x--
			y--
			ans += string(s[x])
		}
	}

	for i := len(ans); i > 0; i-- {
		fmt.Print(string(ans[i-1]))
	}
	fmt.Println("")

}
