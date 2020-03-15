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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W := getInt(), getInt()
	b := make([][]int, H)
	r := make([][]int, H)
	for i := 0; i < H; i++ {
		b[i] = make([]int, W)
		r[i] = make([]int, W)
		s := getString()
		for j := 0; j < W; j++ {
			b[i][j] = int(s[j] - '0')
		}
	}

	for i := 1; i < H-1; i++ {
		for j := 1; j < W-1; j++ {
			if b[i-1][j] != 0 && b[i+1][j] != 0 &&
				b[i][j-1] != 0 && b[i][j+1] != 0 {
				s := min(b[i-1][j], min(b[i+1][j],
					min(b[i][j-1], b[i][j+1])))
				b[i-1][j] -= s
				b[i+1][j] -= s
				b[i][j-1] -= s
				b[i][j+1] -= s
				r[i][j] += s
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Print(r[i][j])
		}
		out()
	}
}
