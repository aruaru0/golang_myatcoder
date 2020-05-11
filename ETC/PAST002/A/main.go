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

func main() {
	sc.Split(bufio.ScanWords)
	s, t := getString(), getString()

	a := []string{"B9", "B8", "B7", "B6", "B5", "B4", "B3", "B2", "B1",
		"1F", "2F", "3F", "4F", "5F", "6F", "7F", "8F", "9F"}

	x := 0
	for a[x] != s {
		x++
	}
	y := 0
	for a[y] != t {
		y++
	}

	out(abs(x - y))
}
