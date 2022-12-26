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

type col struct {
	n, c int
}

type cols []col

func (p cols) Len() int {
	return len(p)
}

func (p cols) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p cols) Less(i, j int) bool {
	return p[i].n > p[j].n
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	n, c := getInt(), getInt()

	p := make([]int, n)
	x := make(cols, 11)
	y := make(cols, 11)
	for i := 0; i < 11; i++ {
		x[i].c = i
		y[i].c = i
	}
	for i := 0; i < n; i++ {
		a := getInt()
		p[i] = a
		if i%2 == 0 {
			x[a].n++
		} else {
			y[a].n++
		}
	}

	sort.Sort(x)
	sort.Sort(y)

	col0 := x[0].c
	col1 := y[0].c
	if col0 == col1 {
		if x[0].n > y[0].n {
			col1 = y[1].c
		} else {
			col0 = x[1].c
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if p[i] != col0 {
				ans += c
			}
		} else {
			if p[i] != col1 {
				ans += c
			}
		}
	}

	out(ans)
}
