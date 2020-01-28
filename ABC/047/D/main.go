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
	T := getInt()
	T++ // エラーにならないためのダミー
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	min := a[0]
	cnt := 0
	maxf := 0
	for i := 1; i < N; i++ {
		diff := a[i] - min

		if diff > maxf {
			maxf = diff
			cnt = 1
		} else if diff == maxf {
			cnt++
		}

		if min > a[i] {
			min = a[i]
		}
	}

	ans := cnt
	out(ans)
}
