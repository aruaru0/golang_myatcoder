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
	sc.Buffer([]byte{}, 1000000)

	N, K := getInt(), getInt()
	x := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = getInt()
	}

	ans := 10010010010
	for i := 0; i < N-K+1; i++ {
		a, b := x[i], x[i+K-1]
		l := 0
		if a >= 0 {
			l = b
		} else if b < 0 {
			l = -a
		} else {
			l = min(-a*2+b, -a+b*2)
		}
		ans = min(ans, l)
		//out(a, b, l)
	}

	out(ans)
}
