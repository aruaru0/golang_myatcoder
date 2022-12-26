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
	sc.Buffer([]byte{}, 100100)

	s := getString()
	m := make(map[int]int)
	for _, v := range s {
		m[int(v-'a')]++
	}
	//out(m)

	m2 := 0
	m1 := 0
	for _, v := range m {
		m2 += v / 2
		m1 += v % 2
	}
	//out(m2, m1)

	if m1 != 0 {
		out((m2/m1)*2 + 1)
	} else {
		out(m2 * 2)
	}

}
