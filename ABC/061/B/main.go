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

// Node :
type Node struct {
	to []int
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	n := make([]Node, N)
	for i := 0; i < M; i++ {
		from, to := getInt()-1, getInt()-1
		n[from].to = append(n[from].to, to)
		n[to].to = append(n[to].to, from)
	}

	for i := 0; i < N; i++ {
		out(len(n[i].to))
	}

}
