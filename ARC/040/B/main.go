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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	_, R := getInt(), getInt()
	s := []byte(getString())

	cnt := 0
	pos := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			i = max(0, i-(R-1))
			if pos == -1 {
				pos = i
			}
			cnt++
		}
	}

	if pos == -1 {
		out(cnt)
	} else {
		out(pos + cnt)
	}
}
