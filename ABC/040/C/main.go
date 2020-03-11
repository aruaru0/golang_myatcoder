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

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	n := make([]int, N)

	n[0] = 0
	for i := 1; i < N; i++ {
		if i == 1 {
			n[i] = asub(a[i], a[i-1])
		} else {
			n[i] = min(
				n[i-1]+asub(a[i], a[i-1]),
				n[i-2]+asub(a[i], a[i-2]))
		}
	}
	out(n[N-1])
}
