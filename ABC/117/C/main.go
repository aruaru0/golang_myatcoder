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

	N, M := getInt(), getInt()
	a := make([]int, M)
	for i := 0; i < M; i++ {
		a[i] = getInt()
	}
	sort.Ints(a)
	d := make([]int, 0)
	for i := 1; i < len(a); i++ {
		d = append(d, a[i]-a[i-1])
	}
	sort.Ints(d)
	// out(d, N)
	ans := 0
	for i := 0; i < len(d)-(N-1); i++ {
		ans += d[i]
	}
	out(ans)
}
