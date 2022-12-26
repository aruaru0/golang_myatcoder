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

const r = 1000000000

func main() {
	sc.Split(bufio.ScanWords)

	S := getInt()
	a := r
	b := 1

	var c, d int
	if S == int(1e18) {
		c = 0
		d = r
	} else if S/r == 0 {
		c = S
		d = 0
	} else {
		c = r - S%r
		d = S/r + 1
	}

	out(0, 0, a, b, c, d)
}
