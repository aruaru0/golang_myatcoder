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
	sc.Buffer([]byte{}, 1000000)
	N, K, X, Y := getInt(), getInt(), getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = 1 - getInt()
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	n := Y / X
	tot := 0
	x := 0
	y := 0
	for len(a) != 0 {
		if len(a) > n {
			r := (-a[0] - tot + K - 1) / K
			y += r
			tot += r * K
		} else {
			x += (-a[0] - tot + K - 1) / K
		}
		a = a[1:]
	}
	out(x*X + y*Y)
}
