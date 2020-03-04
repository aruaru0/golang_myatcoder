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
	b := make([]int, 100001)
	for i := 0; i < N; i++ {
		a := getInt()
		if a-1 >= 0 {
			b[a-1]++
		}
		b[a]++
		if a+1 != N {
			b[a+1]++
		}
	}
	max := 0
	for i := 0; i < 100000; i++ {
		if max < b[i] {
			max = b[i]
		}
	}
	out(max)
}
