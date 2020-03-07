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
	a := make(map[int]int)
	ans := 0
	for i := 0; i < N; i++ {
		v := getInt()
		if a[v] == 0 {
			ans++
			a[v]++
		} else {
			a[v]--
			ans--
		}
	}
	out(ans)
}
