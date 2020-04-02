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

	N := getInt()
	a := make(map[int]int)
	for i := 0; i < N; i++ {
		x := getInt()
		a[x]++
	}

	sum := 0
	for _, v := range a {
		sum += v - 1
	}
	if sum%2 == 1 {
		sum++
	}
	out(N - sum)
}
