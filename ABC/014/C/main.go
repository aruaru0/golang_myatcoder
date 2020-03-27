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

	const N = 1001001
	n := getInt()
	x := make([]int, N)

	for i := 0; i < n; i++ {
		a, b := getInt(), getInt()
		x[a]++
		x[b+1]--
		//out(x[:10])
	}
	m := x[0]
	for i := 1; i < N; i++ {
		x[i] += x[i-1]
		m = max(m, x[i])
	}
	//out(x[:10])
	out(m)
}
