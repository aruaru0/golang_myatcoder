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
	N := getInt()
	c := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		v[i] = getInt()
	}
	for i := 0; i < N; i++ {
		c[i] = getInt()
	}
	n := 1 << uint(N)
	ans := -10001000100
	for i := 0; i < n; i++ {
		C := 0
		V := 0
		for j := 0; j < N; j++ {
			if (i>>uint(j))%2 == 1 {
				C += c[j]
				V += v[j]
			}
		}
		// fmt.Printf("%b %d %d\n", i, V, C)
		ans = max(ans, V-C)
	}
	out(ans)
}