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

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	s := getString()

	r := 0
	for i := 0; i < N; i++ {
		if s[i] == ')' {
			if r != 0 {
				r--
			}
		}
		if s[i] == '(' {
			r++
		}
	}
	l := 0
	for i := N - 1; i >= 0; i-- {
		if s[i] == '(' {
			if l != 0 {
				l--
			}
		}
		if s[i] == ')' {
			l++
		}
	}

	ans := ""
	for i := 0; i < l; i++ {
		ans += "("
	}
	ans += s
	for i := 0; i < r; i++ {
		ans += ")"
	}
	out(ans)
}
