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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 200200)
	_ = getInt()
	s := getString()

	wr := 0
	br := 0
	for _, v := range s {
		if v == '#' {
			br++
		} else {
			wr++
		}
	}
	wl := 0
	bl := 0
	m := wr
	for _, v := range s {
		if v == '#' {
			bl++
			br--
		} else {
			wl++
			wr--
		}
		if bl+wr < m {
			m = bl + wr
		}
	}

	out(m)
}
