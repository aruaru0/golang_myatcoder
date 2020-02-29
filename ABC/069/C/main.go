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
	canDiv2 := 0
	canDiv4 := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		if a[i]%4 == 0 {
			canDiv4++
		} else if a[i]%2 == 0 {
			canDiv2++
		}
	}

	l := canDiv4*2 + canDiv2
	if canDiv2 == 0 {
		l++
	}

	if l >= N {
		out("Yes")
	} else {
		out("No")
	}
}
