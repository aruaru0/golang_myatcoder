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
	N, Ma, Mb := getInt(), getInt(), getInt()
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i], c[i] = getInt(), getInt(), getInt()
	}
	var dp [41][410][410]int
	for i := 0; i <= N; i++ {
		for j := 0; j < 410; j++ {
			for k := 0; k < 410; k++ {
				dp[i][j][k] = math.MaxUint32
			}
		}
	}
	dp[0][0][0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j < 400; j++ {
			for k := 0; k < 400; k++ {
				dp[i+1][j][k] = min(dp[i+1][j][k], dp[i][j][k])
				dp[i+1][j+a[i]][k+b[i]] = min(dp[i+1][j+a[i]][k+b[i]], dp[i][j][k]+c[i])
				//out(i, j, k, " /", j+a[i], j+b[i], dp[i+1][j+a[i]][k+b[i]])
			}
		}
	}

	min := math.MaxInt32
	for i := 1; i < 400; i++ {
		for j := 1; j < 400; j++ {
			if i*Mb != j*Ma {
				continue
			}
			if min > dp[N][i][j] {
				min = dp[N][i][j]
			}

		}
	}
	if min == math.MaxInt32 {
		min = -1
	}
	out(min)
}
