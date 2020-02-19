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

func main() {
	sc.Split(bufio.ScanWords)

	n := []int{getInt(), getInt(), getInt()}
	K := getInt()

	idx := 0
	if n[idx] < n[1] {
		idx = 1
	}

	if n[idx] < n[2] {
		idx = 2
	}

	for i := 0; i < K; i++ {
		n[idx] *= 2
	}

	out(n[0] + n[1] + n[2])
}
