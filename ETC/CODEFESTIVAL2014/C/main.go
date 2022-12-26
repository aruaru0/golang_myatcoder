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

func check(v, n int) bool {
	if n == 1 {
		return false
	}
	x := make([]int, 0)
	for v > 0 {
		a := v % n
		if a >= 10 {
			return false
		}
		x = append(x, a)
		v /= n
	}

	ans := true
	m := 0
	for i := len(x) - 1; i >= 0; i-- {
		m = m*10 + x[i]
	}
	if m != n {
		ans = false
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)

	A := getInt()
	ans := -1
	for i := 1; i <= 10000; i++ {
		if check(A, i) == true {
			ans = i
			break
		}
	}
	out(ans)
}
