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

func asub(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	sc.Split(bufio.ScanWords)

	X, Y := getInt(), getInt()

	if asub(X, Y) <= 1 {
		out("Brown")
	} else {
		out("Alice")
	}

}
