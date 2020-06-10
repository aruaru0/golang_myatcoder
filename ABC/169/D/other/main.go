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

// PfsMap :
func PfsMap(n int) map[int]int {
	pfs := make(map[int]int)
	for n%2 == 0 {
		pfs[2] = pfs[2] + 1
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i] = pfs[i] + 1
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = pfs[n] + 1
	}

	return pfs
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := PfsMap(N)

	ans := 0
	for _, v := range a {
		c := 1
		for v >= c {
			ans++
			v -= c
			c++
		}
	}
	out(ans)
}
