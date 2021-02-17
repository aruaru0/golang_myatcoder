package main

import (
	"bufio"
	"fmt"
	"math"
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	L := getI()
	a := 1
	x := []int{}
	y := []int{}
	cnt := 1
	for a < int(1e6) {
		x = append(x, a)
		y = append(y, cnt)
		cnt++
		a *= 2
	}

	b := make([][2]int, 0)
	for i := len(x) - 1; i >= 0; i-- {
		if L >= x[i] {
			L -= x[i]
			b = append(b, [2]int{x[i], y[i]})
		}
	}
	idx := 1
	ans := make([][3]int, 0)
	cnt = 0
	// out(b)
	e := b[0]
	for i := 0; i < e[1]-1; i++ {
		ans = append(ans, [3]int{idx, idx + 1, 0})
		ans = append(ans, [3]int{idx, idx + 1, 1 << i})
		idx++
	}
	// out(ans)
	tot := e[0]
	for i := 1; i < len(b); i++ {
		ans = append(ans, [3]int{b[i][1], idx, tot})
		// out(b[i], tot)
		tot += b[i][0]
	}
	out(idx, len(ans))
	for _, e := range ans {
		out(e[0], e[1], e[2])
	}
}
