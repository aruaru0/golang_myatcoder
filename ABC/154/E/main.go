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

	N := getString()
	m := getInt()
	n := len(N)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = int(N[i] - '0')
	}

	var dp [101][4][2]int
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 10; l++ {
					nj, nk := j, k
					if l != 0 {
						nj++
					}
					if nj > m {
						continue
					}
					if s[i] < l && k == 0 {
						continue
					}
					if s[i] > l {
						nk = 1
					}
					dp[i+1][nj][nk] += dp[i][j][k]
				}
			}
		}
	}
	out(dp[n][m][0] + dp[n][m][1])
}
