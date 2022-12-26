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

type food struct {
	a, b int
}
type foods []food

func (p foods) Len() int {
	return len(p)
}

func (p foods) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p foods) Less(i, j int) bool {
	return p[i].a+p[i].b > p[j].a+p[j].b
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	A := make(foods, N)
	for i := 0; i < N; i++ {
		a, b := getInt(), getInt()
		A[i] = food{a, b}
	}
	sort.Sort(A)
	ans := 0

	for i := 0; i < N; i++ {
		if i%2 == 0 {
			ans += A[i].a
		} else {
			ans -= A[i].b
		}
	}

	fmt.Println(ans)
}
