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

// 解説を見て書き直したプログラム
func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make(map[int]int)
	for i := 0; i < N; i++ {
		n := getInt()
		a[n]++
	}

	k := len(a)

	//out(a, k)
	if k%2 == 0 {
		k--
	}

	out(k)
}
