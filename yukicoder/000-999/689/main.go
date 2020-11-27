package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

type pair struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	K := getI()

	if K < 4 {
		out(K + 1)
		fmt.Fprint(wr, 6, " ")
		for i := 0; i < K; i++ {
			fmt.Fprint(wr, 7, " ")
		}
		out()
		return
	}

	a := make([][]pair, 10001)
	for i := 2; i <= 10000; i++ {
		for j := 2; j*j <= i; j++ {
			if i%j == 0 && j+i/j <= 240 {
				a[i] = append(a[i], pair{j, i / j})
			}
		}
	}
	b := make([]int, 0, 10001)
	for i := 0; i <= 10000; i++ {
		if len(a[i]) != 0 {
			b = append(b, i)
		}
	}

	l := 0
	r := len(b) - 1
	for l+1 != r {
		m := (l + r) / 2
		if b[m] > K {
			r = m
		} else {
			l = m
		}
	}

	x := a[b[l]]
	e := x[len(x)-1]
	diff := K - e.x*e.y
	out(e.x + e.y + diff + 1)
	for i := 0; i < e.x; i++ {
		fmt.Fprint(wr, 6, " ")
	}
	for i := 0; i < e.y; i++ {
		fmt.Fprint(wr, 7, " ")
	}
	fmt.Fprint(wr, 2, " ")
	for i := 0; i < diff; i++ {
		fmt.Fprint(wr, 3, " ")
	}
	out("")
}
