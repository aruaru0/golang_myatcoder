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

	N, K := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	cnt := 1
	ans := 0
	if K == 1 {
		ans++
	}
	for i := 1; i < N; i++ {
		if a[i] > a[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		//		out(a[i], a[i-1], cnt)
		if cnt >= K {
			ans++
		}
	}

	out(ans)
}
