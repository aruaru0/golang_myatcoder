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
	x := make([]int, N)
	sum := 0
	for i := 0; i < N; i++ {
		x[i] = getInt()
		sum += x[i]
	}
	ave := (sum + N>>1) / N
	ans := 0
	for i := 0; i < N; i++ {
		diff := x[i] - ave
		ans += diff * diff
	}
	out(ans)
}
