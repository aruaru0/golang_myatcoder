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

// 整数用
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
	sc.Buffer([]byte{}, 1000000)

	s := getString()
	t := getString()

	N := len(s)
	sm := make([][]int, 26)
	for i, v := range s {
		p := int(v - 'a')
		sm[p] = append(sm[p], i)
	}

	ans := 0
	prev := -1
	for _, v := range t {
		p := int(v - 'a')
		list := sm[p]
		if len(list) == 0 {
			out(-1)
			return
		}
		pos := upperBound(list, prev)
		// out(list, prev, pos)
		if prev == -1 {
			pos = list[0]
			ans++
		} else if pos == len(list) || prev == -1 {
			ans++
			pos = list[0]
		} else {
			pos = list[pos]
		}
		// out(pos, "ans", ans)
		prev = pos
	}

	prev++
	// out(ans, prev)
	out(ans*N - (N - prev))
}
