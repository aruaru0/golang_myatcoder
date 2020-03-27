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

	L := getInt()
	B := make([]int, L)
	x := 0
	for i := 0; i < L; i++ {
		B[i] = getInt()
		x ^= B[i]
	}

	if x != 0 {
		out(-1)
	} else {
		for i := 0; i < L; i++ {
			out(x)
			x ^= B[i]
		}
	}
}
