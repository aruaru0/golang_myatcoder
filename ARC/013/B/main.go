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
	C := getInt()
	n := make([]int, C)
	m := make([]int, C)
	l := make([]int, C)

	for i := 0; i < C; i++ {
		x := []int{getInt(), getInt(), getInt()}
		sort.Ints(x)
		n[i], m[i], l[i] = x[0], x[1], x[2]
	}

	sort.Ints(n)
	sort.Ints(m)
	sort.Ints(l)

	out(n[C-1] * m[C-1] * l[C-1])

}
