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

type data struct {
	a, i int
}

// Datas :
type Datas []data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	if p[i].a == p[j].a {
		return p[i].i > p[j].i
	}
	return p[i].a > p[j].a
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make(Datas, N)
	for i := 0; i < N; i++ {
		x := getInt()
		a[i] = data{x, i}
	}
	sort.Sort(a)
	a = append(a, data{0, 0})
	b := make([]int, N)
	b[0] = a[0].i
	for i := 1; i < N; i++ {
		b[i] = min(a[i].i, b[i-1])
	}
	// out(a)
	// out(b)
	s := make([]int, N)
	for i := 0; i < N; i++ {
		n := a[i].a - a[i+1].a
		m := b[i]
		// out(n, n*(i+1), m)
		s[m] += n * (i + 1)
		// out(s)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, v := range s {
		fmt.Fprintln(w, v)
	}
}
