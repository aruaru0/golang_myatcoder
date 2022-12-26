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
	if a > b {
		return b
	}
	return a
}

func solve(N, A, B int) {
	diff := B - A
	if diff%2 == 0 {
		out(diff / 2)
		return
	}

	diff = min(A-1, N-B)
	diff = diff + 1 + (B-A-1)/2
	out(diff)
}

func main() {
	sc.Split(bufio.ScanWords)

	N, A, B := getInt(), getInt(), getInt()

	solve(N, A, B)

}
