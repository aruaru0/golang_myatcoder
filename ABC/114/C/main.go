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

func check753(v int) bool {
	x3 := 0
	x5 := 0
	x7 := 0
	for v > 0 {
		x := v % 10
		if x == 3 {
			x3 = 1
		}
		if x == 5 {
			x5 = 1
		}
		if x == 7 {
			x7 = 1
		}
		v /= 10
	}
	if x3+x5+x7 == 3 {
		return true
	}
	return false
}

func rec(v, N int) int {
	if v > N {
		return 0
	}
	ret := 0
	if check753(v) {
		ret = 1
	}
	ret += rec(v*10+3, N)
	ret += rec(v*10+5, N)
	ret += rec(v*10+7, N)

	return ret
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()

	out(rec(0, N))
}
