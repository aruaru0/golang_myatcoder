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

	s := getString()

	a := 0
	b := 0
	for _, v := range s {
		if v == 'A' {
			a++
		} else {
			b++
		}
	}
	if a != 0 && b != 0 {
		out("Yes")
	} else {
		out("No")
	}
}
