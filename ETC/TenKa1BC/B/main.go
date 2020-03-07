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
	a := make([]int, N)
	b := make([]int, N)

	mi := 1001001001
	mipos := -1
	for i := 0; i < N; i++ {
		a[i], b[i] = getInt(), getInt()
		if b[i] < mi {
			mi = b[i]
			mipos = a[i]
		}
	}

	ans := mipos + mi
	out(ans)
}
