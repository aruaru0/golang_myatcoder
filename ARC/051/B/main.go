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

var counter = 0

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	counter++
	return gcd(b, a%b)
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func f(n int) int {
	a0 := 1
	a1 := 1
	if n == 1 || n == 2 {
		return 1
	}
	for i := 3; i <= n; i++ {
		a0, a1 = a1, a0+a1
	}
	return a1
}

func main() {
	sc.Split(bufio.ScanWords)
	K := getInt()

	f0 := f(K)
	f1 := f(K + 1)

	//	gcd(f0, f1)
	//	out("-->", counter)
	out(f0, f1)

}
