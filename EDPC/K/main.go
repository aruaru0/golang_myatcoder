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
	N, K := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	var dp [100001]int

	for i := 1; i <= K; i++ {
		for _, v := range a {
			if i >= v {
				if dp[i-v] == 0 {
					dp[i] = 1
				}
			}
		}
	}

	ans := "First"
	if dp[K] == 0 {
		ans = "Second"
	}

	out(ans)
}

/*
var dp [100001][2]int

func loop(k, j int, a []int) int {
	if dp[k][j] != -1 {
		return dp[k][j]
	}
	ret := 0
	for _, v := range a {
		if k-v >= 0 {
			if loop(k-v, (j+1)%2, a) == 0 {
				ret = 1
			}
		}
	}
	dp[k][j] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	for i := 0; i <= 100000; i++ {
		dp[i][0] = -1
		dp[i][1] = -1
	}

	ret := loop(K, 0, a)

	ans := "First"
	if ret == 0 {
		ans = "Second"
	}

	out(ans)
}
*/
