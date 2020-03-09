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
	a := make([]int, N+1)
	sum := 0
	for i := 1; i <= N; i++ {
		v := getInt()
		sum += v
		a[i] = sum
	}

	ans := 0
	for i := 1; i <= N-K+1; i++ {
		x := a[i+K-1] - a[i-1]
		ans += x
	}

	out(ans)

}
