package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	l := make([]int, N)
	for i := 0; i < N; i++ {
		l[i] = getInt()
	}
	sort.Ints(l)
	ans := 0
	for i := 1; i <= K; i++ {
		ans += l[N-i]
	}

	out(ans)
}
