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

type shop struct {
	s      string
	p, idx int
}

// shops :
type shops []shop

func (p shops) Len() int {
	return len(p)
}

func (p shops) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p shops) Less(i, j int) bool {
	if p[i].s == p[j].s {
		return p[i].p > p[j].p
	}
	return p[i].s < p[j].s
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	n := make(shops, N)
	for i := 0; i < N; i++ {
		s, p := getString(), getInt()
		n[i] = shop{s, p, i}
	}

	sort.Sort(n)

	for _, v := range n {
		out(v.idx + 1)
	}
}
