package main

import (
	"bufio"
	"fmt"
	"os"
)

func out(x ...interface{}) {
	fmt.Println(x...)
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

const mod = 1000000007

func main() {
	r := bufio.NewReaderSize(os.Stdin, 4096)

	var H, W int
	fmt.Scan(&H, &W)

	m := make([][]byte, H)
	for i := 0; i < H; i++ {
		m[i] = readLine(r)
	}

	var dp [1001][1001]int
	dp[0][1] = 1
	for y := 1; y <= H; y++ {
		for x := 1; x <= W; x++ {
			if m[y-1][x-1] == '.' {
				dp[y][x] = (dp[y-1][x] + dp[y][x-1]) % mod
			}
		}
	}

	out(dp[H][W])
}
