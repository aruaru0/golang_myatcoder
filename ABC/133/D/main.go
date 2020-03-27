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
	a := make([]int, N)
	n := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		if i%2 == 0 {
			n += a[i]
		} else {
			n -= a[i]
		}
	}

	x := n / 2
	fmt.Print(2*x, " ")
	for i := 0; i < N-1; i++ {
		x = a[i] - x
		fmt.Print(2*x, " ")
	}
	out()
}
