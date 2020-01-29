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
	x := getInt()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
		if a[i] > x {
			b[i] = a[i] - x
		}
	}

	s := 0
	cnt := 0
	for i := 0; i < N; i++ {
		s = s + a[i]
		d := 0
		if s > x {
			d = s - x
			s = a[i] - d
			a[i] -= d
		} else {
			s = a[i]
		}
		cnt += d
	}
	out(cnt)
}
