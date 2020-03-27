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

	N, Q := getInt(), getInt()
	a := make([]int, N+1)
	for i := 0; i < Q; i++ {
		l, r := getInt()-1, getInt()
		a[l]++
		a[r]--
	}
	cnt := 0
	for i := 0; i < N; i++ {
		cnt += a[i]
		fmt.Print(cnt % 2)
	}
	out()
}
