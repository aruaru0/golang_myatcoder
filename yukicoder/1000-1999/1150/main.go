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

func main() {
	sc.Split(bufio.ScanWords)
	N, s, t := getInt(), getInt()-1, getInt()-1
	a := make([]int, N)
	tot := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		tot += a[i]
	}
	b := make([]int, 0, N*2)
	for i := t + 1; i < N; i++ {
		b = append(b, a[i])
	}
	for i := 0; i < t; i++ {
		b = append(b, a[i])
	}
	pos := N - 1 - t + s
	if s > t {
		pos = s - t - 1
	}
	n := (N - 2 + 1) / 2
	l := pos / 2
	r := (len(b)-pos)/2 + pos - n
	totS := 0
	for i := 0; i <= n; i++ {
		totS += b[i+l]
	}
	ans := totS - (tot - totS)
	for i := l; i < r; i++ {
		totS += -b[i] + b[i+n+1]
		x := totS - (tot - totS)
		ans = max(ans, x)
	}
	out(ans)
}
