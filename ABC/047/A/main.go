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

	a := getInt()
	b := getInt()
	c := getInt()

	ma := max(a, max(b, c))
	su := 0
	if ma == a {
		su = b + c
	} else if ma == b {
		su = a + c
	} else {
		su = a + b
	}

	if ma == su {
		out("Yes")
	} else {
		out("No")
	}

}
