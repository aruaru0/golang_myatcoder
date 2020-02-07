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

	S, C := getInt(), getInt()

	ans := 0
	if S*2 == C {
		ans = S
	} else if S*2 < C {
		ans = (2*S + C) / 4
	} else {
		ans = C / 2
	}
	out(ans)
}
