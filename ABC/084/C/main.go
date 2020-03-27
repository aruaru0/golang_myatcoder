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

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	c := make([]int, N)
	s := make([]int, N)
	f := make([]int, N)
	for i := 0; i < N-1; i++ {
		c[i], s[i], f[i] = getInt(), getInt(), getInt()
	}

	for k := 0; k < N; k++ {
		start := s[k] + c[k]
		for i := k + 1; i < N-1; i++ {
			//out("arrive ", i, " at ", start)
			if s[i] >= start {
				start = s[i] + c[i]
			} else {
				x := ((start + f[i] - 1) / f[i]) * f[i]
				start = x + c[i]
			}
		}

		out(start)
	}
}
