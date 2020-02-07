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

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()

	ans := 1
	for i := 1; i <= N; i++ {
		ans = (ans * i) % mod
	}
	out(ans)
}
