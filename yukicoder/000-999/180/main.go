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

const inf = int(1e18)

var N int
var a, b []int

func calc(x int) int {
	mi, ma := inf, 0
	for i := 0; i < N; i++ {
		c := a[i] + b[i]*x
		mi = min(mi, c)
		ma = max(ma, c)
	}
	d := ma - mi
	return d
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	a = make([]int, N)
	b = make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getInt(), getInt()
	}

	l := 1
	r := int(1e10)
	for l+2 < r {
		diff := abs(r-l) / 3
		m0 := l + diff
		m1 := r - diff
		d0 := calc(m0)
		d1 := calc(m1)
		// out(l, r, d0, d1, m0, m1)
		if d0 <= d1 {
			r = m1
		} else {
			l = m0
		}
	}

	// out(l, r)

	mi := calc(l)
	pos := l
	for i := l + 1; i <= r; i++ {
		d := calc(i)
		if d < mi {
			mi = d
			pos = i
		}
	}

	out(pos)
}
