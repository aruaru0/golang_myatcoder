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

func f(x int, a []int) (int, int) {
	if len(a) == 1 {
		return a[0], a[0]
	}

	l := 0
	r := len(a) - 1
	for l+1 != r {
		m := (l + r) / 2
		if a[m] > x {
			r = m
		} else {
			l = m
		}
	}

	if x == a[l] {
		pos := max(0, l-1)
		d0 := a[l] - a[pos]
		d1 := a[r] - a[l]
		out(d0, d1)
		if d0 < d1 {
			l, r = pos, l
		}
	}
	return a[l], a[r]
}

const inf = 1001001001001

func main() {
	sc.Split(bufio.ScanWords)
	A, B, Q := getInt(), getInt(), getInt()
	s := make([]int, A)
	t := make([]int, B)
	for i := 0; i < A; i++ {
		s[i] = getInt()
	}
	for i := 0; i < B; i++ {
		t[i] = getInt()
	}

	for i := 0; i < Q; i++ {
		x := getInt()
		s0, s1 := f(x, s)
		t0, t1 := f(x, t)

		if s0 > x {
			s0, s1 = inf, s0
		} else if s1 < x {
			s0, s1 = s1, inf
		}

		if t0 > x {
			t0, t1 = inf, t0
		} else if t1 < x {
			t0, t1 = t1, inf
		}

		// out(x, "----")
		// out(s0, s1, t0, t1)
		s0 = abs(x - s0)
		s1 = abs(s1 - x)
		t0 = abs(x - t0)
		t1 = abs(t1 - x)
		// out(s0, s1, t0, t1)
		a0 := max(s0, t0)
		a1 := max(s1, t1)
		a2 := min(s0, t1)*2 + max(s0, t1)
		a3 := min(s1, t0)*2 + max(s1, t0)
		ans := min(min(a0, a1), min(a2, a3))
		out(ans)
	}
}
