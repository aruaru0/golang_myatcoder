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

	N, C, K := getInt(), getInt(), getInt()
	t := make([]int, N)
	for i := 0; i < N; i++ {
		t[i] = getInt()
	}

	sort.Ints(t)
	// out(t)
	ans := 1
	cnt := 1
	end := t[0] + K
	for i := 1; i < N; i++ {
		if cnt < C && t[i] <= end {
			cnt++
			// out(t[i])
		} else {
			// out("----")
			// out(t[i])
			cnt = 1
			end = t[i] + K
			ans++
		}
	}
	out(ans)
}
