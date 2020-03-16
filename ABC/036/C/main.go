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

// Data :
type Data struct {
	x, pos int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i].x < p[j].x
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make(Datas, N)
	for i := 0; i < N; i++ {
		a[i] = Data{getInt(), i}
	}
	sort.Sort(a)
	b := make([]int, N)
	val := a[0].x
	cnt := 0
	for i := 0; i < N; i++ {
		if a[i].x == val {
			b[a[i].pos] = cnt
		} else {
			cnt++
			val = a[i].x
			b[a[i].pos] = cnt
		}
	}

	for i := 0; i < N; i++ {
		out(b[i])
	}

}
