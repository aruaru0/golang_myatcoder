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

	ans := "YES"
	for i := 0; i < N; i++ {
		idx := getInt()
		a[idx]++
		if a[idx] > 1 {
			ans = "NO"
		}
	}

	out(ans)
}
