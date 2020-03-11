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

func getString() string {
	sc.Scan()
	return sc.Text()
}

func asub(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func solve(a []int, N int) int {
	l := list.New()
	l.PushBack(a[0])
	a = a[1:]
	ans := 0
	fpos := 0
	bpos := len(a) - 1
	for i := 0; i < len(a); i++ {
		f := l.Front().Value.(int)
		b := l.Back().Value.(int)
		diff := [4]int{
			asub(f, a[fpos]),
			asub(f, a[bpos]),
			asub(b, a[fpos]),
			asub(b, a[bpos]),
		}
		//out(f, b, a[fpos], a[bpos], diff)
		sel := 0
		for j := 0; j < 4; j++ {
			if diff[sel] < diff[j] {
				sel = j
			}
		}
		switch sel {
		case 0:
			l.PushFront(a[fpos])
			fpos++
		case 1:
			l.PushFront(a[bpos])
			bpos--
		case 2:
			l.PushBack(a[fpos])
			fpos++
		case 3:
			l.PushBack(a[bpos])
			bpos--
		}
		ans += diff[sel]
	}

	return (ans)
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	sort.Ints(a)

	ans := solve(a, N)
	out(ans)
}
