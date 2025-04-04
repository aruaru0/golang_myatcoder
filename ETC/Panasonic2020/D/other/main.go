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

func solve(s string, mx byte, N int) {
	if len(s) == N {
		out(s)
		return
	}
	for c := byte('a'); c <= mx; c++ {
		next := mx
		if c == mx {
			next++
		}
		solve(s+string(c), next, N)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	solve("", 'a', N)
}
