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

func solve(s string) bool {
	f := 0
	t := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[f] != s[t] {
			return false
		}
		f++
		t--
		if f > t {
			break
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	s := getString()

	N := len(s)
	n := (N - 1) / 2

	s0 := s[:n]
	ans0 := solve(s0)

	n = (N+3)/2 - 1
	s1 := s[n:]
	ans1 := solve(s1)

	ans2 := solve(s)
	ans := ans0 && ans1 && ans2

	if ans == true {
		out("Yes")
	} else {
		out("No")
	}
}
