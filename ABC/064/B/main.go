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

	N := getInt()
	//a := make([]int, N)
	min := 1001001001
	max := 0
	for i := 0; i < N; i++ {
		a := getInt()
		if a > max {
			max = a
		}
		if a < min {
			min = a
		}
	}

	out(max - min)
}
