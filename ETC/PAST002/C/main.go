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

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	s := make([][]byte, N)
	for i := 0; i < N; i++ {
		s[i] = make([]byte, 2*N-1)
		t := getString()
		for j := 0; j < 2*N-1; j++ {
			s[i][j] = t[j]
		}
	}
	for i := N - 2; i >= 0; i-- {
		for j := 1; j < 2*N-2; j++ {
			if s[i][j] == '#' {
				if s[i+1][j-1] == 'X' || s[i+1][j] == 'X' || s[i+1][j+1] == 'X' {
					s[i][j] = 'X'
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		out(string(s[i]))
	}
}
