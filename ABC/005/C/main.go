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

	T, N := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	M := getInt()
	b := make([]int, M)
	for i := 0; i < M; i++ {
		b[i] = getInt()
	}

	if M > N {
		out("no")
		return
	}

	for m, n := 0, 0; m < M; m++ {
		if n == N {
			out("no")
			return
		}
		if b[m] < a[n] {
			out("no")
			return
		}
		for b[m] > a[n]+T {
			n++
			if n == N {
				out("no")
				return
			}
		}
		n++
	}
	out("yes")
}
