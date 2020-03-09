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

const mod = 10007

func main() {
	sc.Split(bufio.ScanWords)

	n := getInt()

	switch n {
	case 1:
		out(0)
		return
	case 2:
		out(0)
		return
	case 3:
		out(1)
		return
	}
	a0 := 0
	a1 := 0
	a2 := 1
	ans := 0
	for i := 3; i < n; i++ {
		ans = (a2 + a1 + a0) % mod
		a0, a1, a2 = a1, a2, ans
	}
	out(ans)
}
