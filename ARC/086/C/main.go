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
	N, K := getInt(), getInt()
	a := make(map[int]int)
	for i := 0; i < N; i++ {
		x := getInt()
		a[x]++
	}
	b := make([]int, len(a))
	idx := 0
	for _, v := range a {
		b[idx] = v
		idx++
	}
	sort.Ints(b)
	// out(b)
	sum := 0
	for i := len(b) - 1; i >= 0; i-- {
		// out(i, sum, K)
		if sum >= K {
			break
		}
		sum++
	}
	ans := 0
	for i := 0; i < len(b)-sum; i++ {
		ans += b[i]
	}
	out(ans)
}
