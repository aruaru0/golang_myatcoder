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
	sc.Buffer([]byte{}, 1000000)

	s := getString()

	nest := 0
	ans := "set"
	for _, v := range s {
		switch v {
		case '{':
			nest++
		case '}':
			nest--
		case ':':
			if nest == 1 {
				ans = "dict"
			}
		}
	}

	if s == "{}" {
		ans = "dict"
	}
	out(ans)
}
