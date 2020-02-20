package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//fmt.Println(x...)
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
	out(N, W, w, v)
	w0 := w[0]
	var dp [101][101][330]int
	WW := 301

	dp[0][0][0] = 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			for k := 0; k <= WW; k++ {
				dw := w[i-1] - w0
				if k-dw < 0 {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					if dp[i-1][j][k] >= dp[i-1][j-1][k-dw]+v[i-1] {
						dp[i][j][k] = dp[i-1][j][k]
					} else {
						dp[i][j][k] = dp[i-1][j-1][k-dw] + v[i-1]
					}
				}
			}
		}
	}

	res := 0
	for i := 0; i <= N; i++ {
		out("--------")
		for j := 0; j <= WW; j++ {
			out("w = ", i, "v = ", dp[N][i][j], w0*i+j)
			if w0*i+j <= W && res < dp[N][i][j] {
				res = dp[N][i][j]
			}
		}
	}
	fmt.Println(res)

}
