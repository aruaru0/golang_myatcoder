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
	N := getInt()
	A, B := 0, 0
	a := make([]int, N-1)
	b := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		a[i], b[i] = getInt(), getInt()
		A += a[i]
		B += b[i]
	}
	// if N == 2 {
	// 	out(1)
	// 	return
	// }
	alim := 0
	blim := 0
	for i := 0; i < N-1; i++ {
		if b[i] > A-a[i] {
			alim = b[i] - (A - a[i])
		}
		if a[i] > B-b[i] {
			blim = a[i] - (B - b[i])
		}
	}
	// out(A, B, alim, blim)
	// out(A, B)
	if A > B {
		out(B + 1 - alim)
		return
	}
	out(A + 1 - blim)
}
