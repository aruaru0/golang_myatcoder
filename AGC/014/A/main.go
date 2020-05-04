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

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func solve(a, b, c int) {
	if a%2 == 1 || b%2 == 1 || c%2 == 1 {
		out(0)
		return
	}
	flg := true
	cnt := 0
	for i := 0; i < 10000; i++ {
		a, b, c = b+c, a+c, a+b
		if a%2 == 1 || b%2 == 1 || c%2 == 1 {
			flg = false
			break
		}
		cnt++
		a /= 2
		b /= 2
		c /= 2
	}
	if flg == false {
		out(cnt)
	} else {
		out(-1)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	A, B, C := getInt(), getInt(), getInt()

	// a := PfsMap(A)
	// b := PfsMap(A)
	// c := PfsMap(A)

	// out(min(a[2], min(b[2], c[2])))
	solve(A, B, C)
}
