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

func calc(a, s []int, x int) (int, int) {
	tot := 0
	num := 0
	n := len(a)
	for i := 0; i < n; i++ {
		j := lowerBound(a, x-a[i])
		num += n - j
		tot += s[n] - s[j]
		tot += a[i] * (n - j)
	}
	return tot, num
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	sort.Ints(a)
	b := make([]int, N+1)
	for i := 0; i < N; i++ {
		b[i+1] = b[i] + a[i]
	}

	l := 0
	r := 200005
	for l+1 < r {
		c := (l + r) >> 1
		_, num := calc(a, b, c)
		if num >= M {
			l = c
		} else {
			r = c
		}
	}
	ans, num := calc(a, b, l)
	ans -= (num - M) * l
	out(ans)
}
