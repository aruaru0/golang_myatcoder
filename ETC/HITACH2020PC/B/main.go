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

	A, B, M := getInt(), getInt(), getInt()
	a := make([]int, A)
	b := make([]int, B)

	minA := 1001001001
	minB := 1001001001
	for i := 0; i < A; i++ {
		a[i] = getInt()
		if a[i] < minA {
			minA = a[i]
		}
	}
	for i := 0; i < B; i++ {
		b[i] = getInt()
		if b[i] < minB {
			minB = b[i]
		}
	}
	minAB := 1001001001
	for i := 0; i < M; i++ {
		x, y, c := getInt()-1, getInt()-1, getInt()
		sum := a[x] + b[y] - c
		if minAB > sum {
			minAB = sum
		}
	}
	ans := minA + minB
	if ans > minAB {
		ans = minAB
	}
	out(ans)
}
