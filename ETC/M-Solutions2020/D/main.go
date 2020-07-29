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

var N int
var a []int
var tot = 0

func rec(n, x, y int) {
	if n == N-1 {
		tot = max(tot, x*a[n]+y)
		return
	}
	// buy
	rec(n+1, x+y/a[n], y%a[n])
	rec(n+1, 0, y+x*a[n])
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N = getInt()
	a = getInts(N)

	i := 0
	x := 0
	y := 1000
	for i < N-1 {
		pre := a[i]
		for i < N-1 && pre > a[i+1] {
			i++
			pre = a[i]
		}
		x = y / a[i]
		y = y % a[i]
		for i < N-1 && pre <= a[i+1] {
			i++
			pre = a[i]
		}
		y += x * a[i]
		x = 0
	}
	out(y)
	//rec(0, 0, 1000)
	//out(tot)
}
