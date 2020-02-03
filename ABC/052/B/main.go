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
	_ = getInt()
	S := getString()

	x := 0
	max := 0
	for _, v := range S {
		if v == 'I' {
			x++
		}
		if v == 'D' {
			x--
		}
		if x > max {
			max = x
		}
	}

	out(max)
}
