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

	N := getInt()
	l := list.New()

	dir := true
	for i := 0; i < N; i++ {
		a := getInt()
		if dir == true {
			l.PushBack(a)
		} else {
			l.PushFront(a)
		}
		dir = !dir
	}

	if dir == true {
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", e.Value)
		}
	} else {
		for e := l.Back(); e != nil; e = e.Prev() {
			fmt.Printf("%v ", e.Value)
		}
	}
	out()
}
