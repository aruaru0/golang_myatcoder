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

	A, B := getInt(), getInt()
	ans := -1
	for i := 0; i < 1500; i++ {
		a := int(float64(i) * 0.08)
		b := int(float64(i) * 0.10)
		if a == A && b == B {
			ans = i
			break
		}
	}
	out(ans)
}
