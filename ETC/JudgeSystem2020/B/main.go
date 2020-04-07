package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	N := getInt()
	r := make([]int, 0)
	b := make([]int, 0)
	for i := 0; i < N; i++ {
		x, c := getInt(), getString()
		if c == "R" {
			r = append(r, x)
		} else {
			b = append(b, x)
		}
	}
	sort.Ints(r)
	sort.Ints(b)
	for _, v := range r {
		out(v)
	}
	for _, v := range b {
		out(v)
	}
}
