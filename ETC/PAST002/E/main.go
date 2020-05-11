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

func check(a []int, p int) int {
	cnt := 1
	x := a[p]
	// out(x, p)
	for x != p {
		// out(x)
		x = a[x]
		cnt++
	}
	return cnt
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt() - 1
	}
	for i := 0; i < N; i++ {
		ret := check(a, i)
		fmt.Print(ret, " ")
	}
	out()
}
