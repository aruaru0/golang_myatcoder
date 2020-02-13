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

func show(dp [101][100001]int, N, K int) {
	for i := 0; i <= N; i++ {
		for j := 0; j <= K; j++ {
			fmt.Printf("%d\t", dp[i][j])
		}
		fmt.Println("")
	}
	out("-----------------------------")
}

const MOD = 1000000007

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	var dp [101][100001]int

	dp[0][0] = 1
	for i := 1; i <= N; i++ {
		s := 0
		for j := 0; j <= K; j++ {
			s += dp[i-1][j]
			s %= MOD
			dp[i][j] = s
			//			out(j, a[i-1], s)
			if j >= a[i-1] {
				s -= dp[i-1][j-a[i-1]]
				s %= MOD
				if s < 0 {
					s += MOD
				}
				//				out("--", j, a[i-1], dp[i-1][j-a[i-1]], s)
			}
		}
		//		show(dp, N, K)
	}
	out(dp[N][K])
}
