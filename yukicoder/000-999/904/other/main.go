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
	sumA, sumB := 0, 0
	A := make([]int, N-1)
	B := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		A[i], B[i] = getInt(), getInt()
		sumA += A[i]
		sumB += B[i]
	}

	ans := 0
	for a := 0; a <= sumB; a++ {
		b := (sumA + a) - sumB
		ok := false
		if b >= 0 {
			ok = true
		}
		for i := 0; i < N-1; i++ {
			if A[i] > (b+sumB)-B[i] {
				ok = false
			}
		}
		if ok {
			ans++
		}
	}
	out(ans)
}
