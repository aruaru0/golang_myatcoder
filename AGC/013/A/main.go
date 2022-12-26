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
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	if N == 1 {
		out(1)
		return
	}

	dir := 0
	ans := 1
	pre := a[0]
	for i := 1; i < N; i++ {
		if pre == a[i] {
			continue
		}
		if pre > a[i] && dir == -1 {
			pre = a[i]
			continue
		}
		if pre < a[i] && dir == 1 {
			pre = a[i]
			continue
		}
		if dir == 0 {
			if pre > a[i] {
				dir = -1
			} else {
				dir = 1
			}
			pre = a[i]
			continue
		}
		ans++
		dir = 0
		pre = a[i]
	}
	out(ans)
}
