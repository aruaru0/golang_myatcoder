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

func rec(a []int, p int) int {
	if p < 0 {
		return 0
	}
	A := make([]int, 0)
	B := make([]int, 0)
	for _, e := range a {
		if (e>>p)%2 == 0 {
			A = append(A, e)
		} else {
			B = append(B, e)
		}
	}
	ret := 0
	if len(A) == 0 || len(B) == 0 {
		ret = rec(a, p-1)
	} else {
		ret = rec(A, p-1)
		ret = min(ret, rec(B, p-1))
		ret += 1 << p
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := getInts(N)
	out(rec(a, 30))
}
