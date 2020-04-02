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

func srarch(v int, a []int) (int, int) {
	c := v / 2
	l := 0
	r := len(a) - 1
	for l+1 != r {
		m := (l + r) / 2
		if a[m] < c {
			l = m
		} else {
			r = m
		}
	}
	return a[l], a[r]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	n := getInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = getInt()
	}
	sort.Ints(a)

	m := a[len(a)-1]
	//out(a)
	x, y := srarch(m, a)

	if m == y {
		out(m, x)
		return
	}

	if asub(x*2, m) < asub(y*2, m) {
		out(m, x)
	} else {
		out(m, y)
	}
}
