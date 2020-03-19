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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calc(month, day int) int {
	days := []int{
		0,
		31, 29, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31}

	cnt := 0
	for i := 0; i < month; i++ {
		cnt += days[i]
	}
	cnt += day
	cnt--
	return (cnt)
}

func monthday(s string) (int, int) {
	m := 0
	d := 0
	flg := false
	for _, v := range s {
		if v == '/' {
			flg = true
		} else if flg == false {
			m = m*10 + int(v-'0')
		} else {
			d = d*10 + int(v-'0')
		}
	}

	return m, d
}

func main() {
	sc.Split(bufio.ScanWords)

	n := 366
	holiday := make([]bool, n)
	for i := 0; i < n; i++ {
		if i%7 == 0 || i%7 == 6 {
			holiday[i] = true
		}
	}

	N := getInt()
	for i := 0; i < N; i++ {
		s := getString()
		m, d := monthday(s)
		c := calc(m, d)
		for c < n && holiday[c] == true {
			c++
		}
		if c < n {
			holiday[c] = true
		}
	}

	ma := 0
	cnt := 0
	for i := 0; i < n; i++ {
		if holiday[i] == true {
			cnt++
			ma = max(ma, cnt)
		} else {
			cnt = 0
		}
	}
	out(ma)
}
