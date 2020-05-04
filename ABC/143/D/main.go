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

func lowerBound(a []int, x int) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m] >= x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	l := make([]int, N)
	for i := 0; i < N; i++ {
		l[i] = getInt()
	}

	sort.Ints(l)

	ans := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			a := l[i]
			b := l[j]
			n := lowerBound(l, a+b)
			ans += n - (j + 1)
		}
	}
	out(ans)
}
