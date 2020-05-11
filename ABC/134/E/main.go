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

func lowerBound(a []int, x int) int {
	l := 0
	r := len(a) - 1
	for l <= r {
		m := (l + r) / 2
		if a[m] >= x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(a []int, x int) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m] <= x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = -1
	}

	for i := 0; i < N; i++ {
		p := lowerBound(dp, a[i])
		dp[p-1] = a[i]
		// out(p, dp)
	}
	ans := 0
	for i := 0; i < N; i++ {
		if dp[i] != -1 {
			ans++
		}
	}
	out(ans)
}
