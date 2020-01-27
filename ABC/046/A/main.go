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

	a := getInt()
	b := getInt()
	c := getInt()

	ans := 3
	if (a == b) && (b == c) {
		ans = 1
	} else if a == b {
		ans = 2
	} else if b == c {
		ans = 2
	} else if a == c {
		ans = 2
	}
	out(ans)
}
