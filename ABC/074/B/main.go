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

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	x := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = getInt()
	}
	ans := 0
	for i := 0; i < N; i++ {
		d0 := asub(x[i], K)
		d1 := x[i]
		if d0 > d1 {
			ans += d1 * 2
		} else {
			ans += d0 * 2
		}
	}
	out(ans)
}
