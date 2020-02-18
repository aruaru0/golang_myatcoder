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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func show(dp [3001][3001]int, N int) {
	out("-------------------------------")
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if i == j {
				fmt.Print("*\t")
			} else {
				fmt.Printf("%d\t", dp[i][j])
			}
		}
		fmt.Println("")
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	var dp [3001][3001]int

	for len := 1; len <= N; len++ {
		for i := 0; i+len <= N; i++ {
			j := i + len
			//out(a[i:j], a[i]-dp[i+1][j], a[j-1]-dp[i][j-1])
			dp[i][j] = max(a[i]-dp[i+1][j], a[j-1]-dp[i][j-1])
		}
		//show(dp, N)
	}

	out(dp[0][N])
}
