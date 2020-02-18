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

/*
func readLine(r *bufio.Reader) []byte {
	buf := make([]byte, 0, 1024)
	for {
		l, p, e := r.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return buf
}
	r := bufio.NewReaderSize(os.Stdin, 4096)
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dp [3001][3001]int

func solve(n, f, t, N int, a []int) int {
	if n == N {
		return 0
	}
	if dp[f][t] != -1 {
		return dp[f][t]
	}
	ret1 := a[f] - solve(n+1, f+1, t, N, a)
	ret2 := a[t] - solve(n+1, f, t-1, N, a)

	ret := max(ret1, ret2)
	dp[f][t] = ret
	return (ret)
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	for i := 0; i <= 3000; i++ {
		for j := 0; j <= 3000; j++ {
			dp[i][j] = -1
		}
	}

	out(solve(0, 0, N-1, N, a))

}
