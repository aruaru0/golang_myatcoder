package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func isUpper(a byte) bool {
	if a >= 'A' && a <= 'Z' {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	s := getString()

	m := make([]string, 0)

	st := 0
	flg := false
	for i := 0; i < len(s); i++ {
		if isUpper(s[i]) {
			if flg == true {
				flg = false
				m = append(m, strings.ToLower(s[st:i+1]))
			} else {
				st = i
				flg = true
			}
		}
	}

	sort.Strings(m)
	for i := 0; i < len(m); i++ {
		x := make([]byte, len(m[i]))
		copy(x, m[i])
		x[0] = x[0] - 'a' + 'A'
		x[len(x)-1] = x[len(x)-1] - 'a' + 'A'
		fmt.Print(string(x))
	}
	out()
}
