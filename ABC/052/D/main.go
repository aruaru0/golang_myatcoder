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

func main() {
	sc.Split(bufio.ScanWords)

	N, A, B := getInt(), getInt(), getInt()
	x := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = getInt()
	}

	cost := 0
	for i := 1; i < N; i++ {
		diff := (x[i] - x[i-1]) * A
		if diff > B {
			cost += B
		} else {
			cost += diff
		}
	}

	out(cost)
}
