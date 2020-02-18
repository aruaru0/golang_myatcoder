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

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

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

var dp [1000000]int

const inf = 1 << 60

func solve(s string) int {
	var dp [10000100][2]int

	for i := 0; i < 10000100; i++ {
		dp[i][0] = inf
		dp[i][1] = inf
	}

	dp[0][0] = 0
	size := len(s)
	for i := 0; i < size; i++ {
		c := int(s[i] - '0')
		dp[i+1][0] = min(dp[i+1][0], dp[i][0]+c)
		dp[i+1][0] = min(dp[i+1][0], dp[i][1]+c+1)
		dp[i+1][1] = min(dp[i+1][1], dp[i][0]+10-c)
		dp[i+1][1] = min(dp[i+1][1], dp[i][1]+10-c-1)
	}
	return min(dp[size][0], dp[size][1]+1)
}

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
func main() {
	r := bufio.NewReaderSize(os.Stdin, 4096)

	s := reverse(string(readLine(r)))

	out(solve(s))
}
