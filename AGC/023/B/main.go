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

func check(A, B, N int, s []string) bool {
	ok := true
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			// out(i, j, s[i+A][j+B], s[j+A][i+B])
			if s[i+A][j+B] != s[j+A][i+B] {
				ok = false
				break
			}
		}
	}
	return ok
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	s := make([]string, 2*N)
	for i := 0; i < N; i++ {
		a := getString()
		s[i] = a + a
	}
	for i := 0; i < N; i++ {
		s[N+i] = s[i]
	}

	ans := 0
	for i := 0; i < N; i++ {
		ok := check(i, 0, N, s)
		if ok {
			ans++
		}
	}

	out(ans * N)
}
