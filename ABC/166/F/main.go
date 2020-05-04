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

func check(x, y int, a, next [3]int) ([3]int, int, bool) {
	if a[x] == 0 && a[y] == 0 {
		return a, 0, false
	}
	if a[x] == 0 {
		a[x]++
		a[y]--
		return a, x, true
	}
	if a[y] == 0 {
		a[x]--
		a[y]++
		return a, y, true
	}
	if next[x] == 1 {
		a[x]++
		a[y]--
		return a, x, true
	}
	a[x]--
	a[y]++
	return a, y, true
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := [3]int{}
	for i := 0; i < 3; i++ {
		a[i] = getInt()
	}
	p := make([][3]int, N+1)
	for i := 0; i < N; i++ {
		s := getString()
		switch s {
		case "AB":
			p[i][0] = 1
			p[i][1] = 1
		case "BC":
			p[i][1] = 1
			p[i][2] = 1
		case "AC":
			p[i][0] = 1
			p[i][2] = 1
		}
	}

	// out(a, p)

	ans := make([]int, N)
	for i := 0; i < N; i++ {
		var flg bool
		var num int
		if p[i][0] == 1 && p[i][1] == 1 { // AB
			a, num, flg = check(0, 1, a, p[i+1])
		} else if p[i][1] == 1 && p[i][2] == 1 { // BC
			a, num, flg = check(1, 2, a, p[i+1])
		} else { // AC
			a, num, flg = check(0, 2, a, p[i+1])
		}
		if flg == false {
			out("No")
			return
		}
		ans[i] = num
		// out(a)
	}

	out("Yes")

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := 0; i < N; i++ {
		fmt.Println(string('A' + ans[i]))
	}
}
