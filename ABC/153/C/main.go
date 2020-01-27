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

	N := getInt()
	K := getInt()
	H := make([]int, N)
	for i := 0; i < N; i++ {
		H[i] = getInt()
	}
	sort.Ints(H)

	ans := 0
	for i := 0; i < N-K; i++ {
		ans += H[i]
	}

	out(ans)
}
