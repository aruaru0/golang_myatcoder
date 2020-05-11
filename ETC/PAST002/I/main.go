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

type person struct {
	idx, a int
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	n := 1 << uint(N)
	a := make([]person, n)
	for i := 0; i < n; i++ {
		x := getInt()
		a[i] = person{i, x}
	}

	po := make([]int, n)
	for i := 0; i < N; i++ {
		for j := 0; j < n; j += 2 {
			if a[j].a > a[j+1].a {
				po[a[j+1].idx] = i + 1
				a[j/2] = a[j]
			} else {
				po[a[j].idx] = i + 1
				a[j/2] = a[j+1]
			}
		}
		n /= 2
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, v := range po {
		if v == 0 {
			fmt.Fprintln(w, N)
		} else {
			fmt.Fprintln(w, v)
		}
	}
}
