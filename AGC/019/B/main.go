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
	sc.Buffer([]byte{}, 1000000)

	S := []byte(getString())
	n := len(S)

	N := n*(n-1)/2 + 1
	a := make(map[byte]int)
	for _, v := range S {
		a[v]++
	}
	for _, v := range a {
		N -= v * (v - 1) / 2
	}
	out(N)
}
