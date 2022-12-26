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

func nextDay(y, m, d int) (int, int, int) {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	day := days[m-1]
	if m == 2 {
		if y%400 == 0 {
			day++
		} else if y%100 == 0 {

		} else if y%4 == 0 {
			day++
		}
	}
	d++
	if d > day {
		d = 1
		m++
		if m == 13 {
			y++
			m = 1
		}
	}
	return y, m, d
}

func main() {
	sc.Split(bufio.ScanWords)
	var y, m, d int
	fmt.Scanf("%d/%d/%d", &y, &m, &d)
	for {
		if y%(m*d) == 0 {
			fmt.Printf("%4.4d/%2.2d/%2.2d\n", y, m, d)
			return
		}
		y, m, d = nextDay(y, m, d)
	}
}
