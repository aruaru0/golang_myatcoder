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

func check(n int, k int, a []int) int {
	for i := 0; i < k; i++ {
		n = n - n%a[i]
	}
	return n
}

func checkmax(k, ma int, a []int) int {
	l := 0
	r := 2 + k*ma
	for l+1 != r {
		m := (l + r) / 2
		ret := check(m, k, a)
		// out(m, ret)
		if ret > 2 {
			r = m
		} else {
			l = m
		}
	}
	// out(l, r)
	if check(r, k, a) == 2 {
		return r
	}
	if check(l, k, a) == 2 {
		return l
	}
	return -1
}

func checkmin(k, ma int, a []int) int {
	l := 0
	r := 2 + k*ma
	for l+1 != r {
		m := (l + r) / 2
		ret := check(m, k, a)
		// out(m, ret)
		if ret >= 2 {
			r = m
		} else {
			l = m
		}
	}
	// out(l, r)
	if check(l, k, a) == 2 {
		return l
	}
	if check(r, k, a) == 2 {
		return r
	}
	return -1
}

func main() {
	sc.Split(bufio.ScanWords)

	K := getInt()
	a := make([]int, K)
	mx := 0
	for i := 0; i < K; i++ {
		a[i] = getInt()
		mx = max(mx, a[i])
	}

	mi := checkmin(K, mx, a)
	ma := checkmax(K, mx, a)
	if mi != -1 && ma != -1 {
		out(mi, ma)
	} else {
		out(-1)
	}
}
