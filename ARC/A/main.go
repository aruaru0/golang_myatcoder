package main

import (
	"bufio"
	"container/list"
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
	N, M := getInt(), getInt()

	l := list.New()
	for i := 0; i < M; i++ {
		a := getInt()
		l.PushFront(a)
	}

	n := make([]int, N+1)
	for e := l.Front(); e != nil; e = e.Next() {
		if n[e.Value.(int)] == 0 {
			out(e.Value.(int))
			n[e.Value.(int)] = 1
		}
	}
	for i := 1; i <= N; i++ {
		if n[i] == 0 {
			out(i)
		}
	}
}
