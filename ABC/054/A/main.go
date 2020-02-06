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
	A, B := getInt(), getInt()
	if A == 1 {
		A = 14
	}
	if B == 1 {
		B = 14
	}

	if A > B {
		out("Alice")
	} else if A < B {
		out("Bob")
	} else {
		out("Draw")
	}
}
