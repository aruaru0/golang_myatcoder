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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	N, T := getInt(), getInt()
	t := make([]int, N)
	for i := 0; i < N; i++ {
		t[i] = getInt()
	}

	total := 0
	for i := 1; i < N; i++ {
		total += min(T, t[i]-t[i-1])
	}
	total += T

	out(total)
}
