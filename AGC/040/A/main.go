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
	sc.Buffer([]byte{}, 1000000)
	s := getString()
	N := len(s)
	l := make([]int, N+1)
	r := make([]int, N+1)

	for i := 0; i < N; i++ {
		if s[i] == '<' {
			l[i+1] = l[i] + 1
		}
	}
	for i := N - 1; i >= 0; i-- {
		if s[i] == '>' {
			r[i] = r[i+1] + 1
		}
	}
	// out(l)
	// out(r)
	ans := 0
	for i := 0; i < N+1; i++ {
		ans += max(l[i], r[i])
	}
	out(ans)
}
