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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	a := make([]int, N)
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		a[i] = getInt()
		m[a[i]]++
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	all := 0
	for _, v := range m {
		if v >= 2 {
			all += v * (v - 1) / 2
		}
	}

	for i := 0; i < N; i++ {
		x := m[a[i]]
		y := all
		if x >= 2 {
			y = all - x*(x-1)/2 + (x-1)*(x-2)/2
		}
		fmt.Fprintln(w, y)
	}
}
