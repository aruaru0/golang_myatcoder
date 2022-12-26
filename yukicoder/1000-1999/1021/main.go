package main

import (
	"bufio"
	"container/list"
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, _ := getInt(), getInt()
	a := getInts(N)
	S := getString()

	l := list.New()
	for i := 0; i < N; i++ {
		l.PushBack(a[i])
	}
	for _, e := range S {
		if e == 'L' {
			e := l.Front()
			x := e.Value.(int)
			l.Remove(e)
			e = l.Front()
			x += e.Value.(int)
			l.Remove(e)
			l.PushFront(x)
			l.PushBack(0)
		} else {
			e := l.Back()
			x := e.Value.(int)
			l.Remove(e)
			e = l.Back()
			x += e.Value.(int)
			l.Remove(e)
			l.PushBack(x)
			l.PushFront(0)
		}
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	out()
}
