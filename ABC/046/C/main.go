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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func b2i(b bool) int {
	if b == true {
		return 1
	}
	return 0
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	T := make([]int, N)
	A := make([]int, N)

	for i := 0; i < N; i++ {
		T[i] = getInt()
		A[i] = getInt()
	}

	t := T[0]
	a := A[0]
	for i := 1; i < N; i++ {
		tt := t/T[i] + b2i(t%T[i] != 0)
		aa := a/A[i] + b2i(a%A[i] != 0)
		max := tt
		if tt < aa {
			max = aa
		}
		t = max * T[i]
		a = max * A[i]
	}
	fmt.Println(a + t)
}
