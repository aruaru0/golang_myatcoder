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

var mod = 1000000007

func f(n int, m map[int]int) int {
	v, ok := m[n]
	if ok {
		return v
	}
	m[n] = (f(n/2, m) + f((n-1)/2, m) + f(n/2-1, m)) % mod
	return m[n]
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()

	m := make(map[int]int)
	m[0] = 1
	m[1] = 2
	out(f(N, m))
}
