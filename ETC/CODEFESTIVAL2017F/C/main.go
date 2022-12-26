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

func calc(t []int) int {
	ret := 24
	cur := 0
	for i := 1; i < 24; i++ {
		if t[i] == 1 {
			x := i - cur
			ret = min(ret, x)
			cur = i
		}
	}
	ret = min(ret, 24-cur)
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	d := make([]int, 13)
	d[0] = 1
	for i := 0; i < N; i++ {
		x := getInt()
		d[x]++
	}

	m := 0
	if d[0] >= 2 || d[12] >= 2 {
		out(0)
		return
	}
	for _, v := range d {
		m = max(m, v)
	}
	if m >= 3 {
		out(0)
		return
	}

	len := 1 << 13
	ans := 0
	for i := 0; i < len; i++ {
		a := make([]int, 13)
		idx := 0
		for n := i; n > 0; n /= 2 {
			a[12-idx] = n % 2
			idx++
		}
		// out(a)
		t := make([]int, 24)
		for i, v := range a {
			if i == 0 {
				t[0] = 1
				continue
			}
			switch d[i] {
			case 1:
				if v == 0 {
					t[i] = 1
				} else {
					if i == 12 {
						t[i] = 1
					} else {
						t[24-i] = 1
					}
				}
			case 2:
				t[i] = 1
				t[24-i] = 1
			}
		}
		l := calc(t)
		ans = max(ans, l)
		// out(t, l)
	}
	out(ans % 24)
}
