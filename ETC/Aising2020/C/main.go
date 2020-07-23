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

func f(n int) []int {
	ret := make([]int, n+1)
	for x := 1; x*x < n; x++ {
		xx := x * x
		for y := 1; y*y < n; y++ {
			yy := y * y
			xy := x * y
			for z := 1; z*z < n; z++ {
				zz := z * z
				xz := z * x
				yz := z * y
				tot := xx + yy + zz + xy + xz + yz
				if tot <= n {
					ret[tot]++
				}
			}
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	ans := f(N)
	for i := 1; i <= N; i++ {
		out(ans[i])
	}
}
