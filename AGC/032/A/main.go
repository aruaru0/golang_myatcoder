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

func del(a []int) int {
	ret := -1
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == i {
			ret = i
			break
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt() - 1
	}

	ans := make([]int, 0, N)
	for i := 0; i < N; i++ {
		remove := del(a)
		if remove == -1 {
			out(-1)
			return
		}
		a = append(a[:remove], a[remove+1:]...)
		ans = append(ans, remove)
	}
	for i := len(ans) - 1; i >= 0; i-- {
		out(ans[i] + 1)
	}
}
