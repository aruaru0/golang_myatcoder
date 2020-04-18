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

type item struct {
	v, l int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, K := getInt(), getInt()
	s := getString()
	a := make([]item, 0)
	cnt := 1
	c := s[0]
	for i := 1; i < N; i++ {
		if s[i] != c {
			a = append(a, item{int(c - '0'), cnt})
			cnt = 0
			c = s[i]
		}
		cnt++
	}
	a = append(a, item{int(c - '0'), cnt})
	// out(a)
	ans := 0
	cnt = 0
	idx := 0
	flg := false
	pidx := 0
	for cnt != K {
		ans += a[idx].l
		if a[idx].v == 0 {
			cnt++
			if !flg {
				pidx = idx
				flg = true
			}
		}
		idx++
		if idx == len(a) {
			out(ans)
			return
		}
	}

	if idx != len(a)-1 {
		ans += a[idx].l
		idx++
	}

	prev := ans
	for idx < len(a) {
		x := prev - a[pidx].l
		if pidx != 0 {
			x -= a[pidx-1].l
		}
		// out(x)
		x += a[idx].l
		if idx+1 != len(a) {
			x += a[idx+1].l
		}
		// out(x)
		// out("idx", idx, "pidx", pidx, "x", x)
		prev = x
		ans = max(ans, x)
		pidx += 2
		idx += 2
	}

	out(ans)
}
