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

	ans := true

	if len(s)%2 == 1 {
		out("No")
		return
	}

	for i := 0; i < len(s); i += 2 {
		if s[i] != 'h' || s[i+1] != 'i' {
			ans = false
			break
		}
	}
	if ans {
		out("Yes")
	} else {
		out("No")
	}

}
