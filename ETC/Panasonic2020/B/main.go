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

func solve(H, W int) int {
	l0 := W / 2
	l1 := W / 2
	if W%2 == 1 {
		l0++
	}
	h := H / 2
	ans := (l0 + l1) * h
	if H%2 == 1 {
		ans += l0
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	H, W := getInt(), getInt()

	n := H * W

	if H == 1 || W == 1 {
		out(1)
		return
	}
	if n%2 == 1 {
		out(n/2 + 1)
	} else {
		out(n / 2)
	}
}
