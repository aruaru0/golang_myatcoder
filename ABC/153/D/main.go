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

func attack(H int) int {
	if H == 1 {
		return 1
	}
	return attack(H/2)*2 + 1
}

func main() {
	sc.Split(bufio.ScanWords)

	H := getInt()

	ans := attack(H)

	out(ans)
}
