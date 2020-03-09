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
	sc.Buffer([]byte{}, 100100)

	s := getString()
	Q := getInt()

	l := list.New()
	for i := 0; i < len(s); i++ {
		l.PushBack(s[i])
	}

	rev := false
	for i := 0; i < Q; i++ {
		t := getInt()
		if t == 1 {
			rev = !rev
		} else {
			f, c := getInt(), getString()[0]
			if f == 1 {
				if rev == true {
					l.PushBack(c)
				} else {
					l.PushFront(c)
				}
			} else {
				if rev == true {
					l.PushFront(c)
				} else {
					l.PushBack(c)
				}
			}
		}
	}

	ans := make([]byte, l.Len())

	if rev == true {
		idx := 0
		for e := l.Back(); e != nil; e = e.Prev() {
			ans[idx] = e.Value.(byte)
			idx++
		}
	} else {
		idx := 0
		for e := l.Front(); e != nil; e = e.Next() {
			ans[idx] = e.Value.(byte)
			idx++
		}
	}

	out(string(ans))
}
