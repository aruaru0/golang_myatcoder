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

func group(n int) int {
	if n == 2 {
		return 2
	}
	if n == 4 || n == 6 || n == 9 || n == 11 {
		return 1
	}
	return 0
}

func main() {
	sc.Split(bufio.ScanWords)

	x, y := getInt(), getInt()

	if group(x) == group(y) {
		out("Yes")
	} else {
		out("No")
	}

}
